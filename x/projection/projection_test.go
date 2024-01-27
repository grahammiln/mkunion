package projection

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/storage/schemaless"
	"github.com/widmogrod/mkunion/x/stream"
	"math"
	"testing"
	"time"
)

func TestProjection_HappyPath(t *testing.T) {
	dataStream := stream.NewInMemoryStream[schema.Schema](stream.WithSystemTime)
	stream1 := stream.NewTypedStreamTopic[int](dataStream, "topic-out-1")
	stream2 := stream.NewTypedStreamTopic[float64](dataStream, "topic-out-2")
	stream3 := stream.NewTypedStreamTopic[float64](dataStream, "topic-out-3")
	stream4 := stream.NewTypedStreamTopic[string](dataStream, "topic-out-4")

	state1 := &PullPushContextState{
		Offset:    nil,
		PullTopic: "",
		PushTopic: stream1.TopicName(),
	}

	ctx1 := NewPushOnlyInMemoryContext[int](state1, stream1)

	err := DoLoad(ctx1, func(push func(*Record[int]) error) error {
		for i := 0; i < 10; i++ {
			err := push(&Record[int]{
				Key:  fmt.Sprintf("key-%d", i%2),
				Data: i,
			})
			if err != nil {
				return fmt.Errorf("projection.Range: push: %w", err)
			}
		}
		return nil
	})
	assert.NoError(t, err)

	state2 := &PullPushContextState{
		Offset:    nil,
		PullTopic: stream1.TopicName(),
		PushTopic: stream2.TopicName(),
	}
	ctx2 := NewPushAndPullInMemoryContext[int, float64](state2, stream1, stream2)
	err = DoMap[int, float64](ctx2, func(x *Record[int]) *Record[float64] {
		return &Record[float64]{
			Key:       x.Key,
			Data:      float64(x.Data) * 2,
			EventTime: x.EventTime,
		}
	})
	assert.NoError(t, err)

	orderOfEvents := []string{}
	ctx4 := NewJoinInMemoryContext[int, float64, float64](
		&JoinContextState{
			Offset1:    nil,
			PullTopic1: stream1.TopicName(),
			Offset2:    nil,
			PullTopic2: stream2.TopicName(),
			PushTopic:  stream3.TopicName(),
		},
		stream1,
		stream2,
		stream3,
	)

	err = DoSink[Either[int, float64]](ctx4, func(x *Record[Either[int, float64]]) error {
		return MatchEitherR1(
			x.Data,
			func(x *Left[int, float64]) error {
				orderOfEvents = append(orderOfEvents, fmt.Sprintf("left-%d", x.Left))
				return nil
			},
			func(x *Right[int, float64]) error {
				orderOfEvents = append(orderOfEvents, fmt.Sprintf("right-%.2f", x.Right))
				return nil
			},
		)
	})
	assert.NoError(t, err)

	expectedOrder := []string{
		"left-0",
		"right-0.00",
		"left-1",
		"right-2.00",
		"left-2",
		"right-4.00",
		"left-3",
		"right-6.00",
		"left-4",
		"right-8.00",
		"left-5",
		"right-10.00",
		"left-6",
		"right-12.00",
		"left-7",
		"right-14.00",
		"left-8",
		"right-16.00",
		"left-9",
		"right-18.00",
	}

	if diff := cmp.Diff(expectedOrder, orderOfEvents); diff != "" {
		t.Fatalf("NewJoinPushAndPullContext: diff: (-want +got)\n%s", diff)
	}

	ctx5 := NewJoinInMemoryContext[int, float64, string](
		&JoinContextState{
			Offset1:    nil,
			PullTopic1: stream1.TopicName(),
			Offset2:    nil,
			PullTopic2: stream2.TopicName(),
			PushTopic:  stream4.TopicName(),
		},
		stream1,
		stream2,
		stream4,
	)

	wd := &FixedWindow{Width: math.MaxInt64}
	fm := &Discard{}
	td := &AtWatermark{}
	err = DoWindow(ctx5, wd, fm, td, "", func(x Either[int, float64], agg string) (string, error) {
		var concat string
		if agg == "" {
			concat = fmt.Sprintf("%v", x)
		} else {
			concat = fmt.Sprintf("%s,%v", agg, x)
		}

		return concat, nil
	})
	assert.NoError(t, err)

	orderOfEvents = []string{}
	state3 := &PullPushContextState{
		Offset:    nil,
		PullTopic: "topic-out-4",
		PushTopic: "topic-out-3",
	}
	ctx6 := NewPullOnlyInMemoryContext[string](state3, stream4)
	err = DoSink[string](ctx6, func(x *Record[string]) error {
		orderOfEvents = append(orderOfEvents, fmt.Sprintf("record-%s:%s", x.Key, x.Data))
		return nil
	})
	assert.NoError(t, err)

	expectedOrder = []string{
		"record-key-0:&{0},&{0},&{2},&{4},&{4},&{8},&{6},&{12},&{8},&{16}",
		"record-key-1:&{1},&{2},&{3},&{6},&{5},&{10},&{7},&{14},&{9},&{18}",
	}

	if diff := cmp.Diff(expectedOrder, orderOfEvents); diff != "" {
		t.Fatalf("NewJoinPushAndPullContext: diff: (-want +got)\n%s", diff)
	}
}

func TestProjection_Recovery(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	probabilityOfFailure := 0.0
	recoveryAttempts := uint8(50)

	dataStream := stream.NewInMemoryStream[schema.Schema](stream.WithSystemTimeFixed(0))
	stream1 := stream.NewTypedStreamTopic[int](dataStream, "topic-out-1")
	stream2 := stream.NewTypedStreamTopic[float64](dataStream, "topic-out-2")
	stream3 := stream.NewTypedStreamTopic[float64](dataStream, "topic-out-3")

	dataStream.SimulateRuntimeProblem(&stream.SimulateProblem{
		ErrorOnPullProbability: probabilityOfFailure,
		ErrorOnPush:            fmt.Errorf("simulated push error"),

		ErrorOnPushProbability: probabilityOfFailure,
		ErrorOnPull:            fmt.Errorf("simulated pull error"),
	})

	var store schemaless.Repository[SnapshotState] = schemaless.NewInMemoryRepository[SnapshotState]()

	recovery :=
		NewRecoveryOptions(
			"recovery-load",
			&PullPushContextState{
				Offset:    nil,
				PullTopic: "",
				PushTopic: stream1.TopicName(),
			},
			store,
		).
			WithMaxRecoveryAttempts(recoveryAttempts).
			WithAutoSnapshot(false)

	err := Recovery[*PullPushContextState](
		recovery,
		func(state *PullPushContextState) (*PushAndPullInMemoryContext[any, int], error) {
			ctx1 := NewPushAndPullInMemoryContext[any, int](state, nil, stream1)
			ctx1.SimulateRuntimeProblem(&SimulateProblem{
				ErrorOnPushOutProbability: probabilityOfFailure,
				ErrorOnPushOut:            fmt.Errorf("simulated push error"),
				ErrorOnPullInProbability:  probabilityOfFailure,
				ErrorOnPullIn:             fmt.Errorf("simulated pull error"),
			})
			return ctx1, nil
		},
		func(ctx1 *PushAndPullInMemoryContext[any, int]) error {
			state := ctx1.CurrentState()
			startValue := 0
			if x, ok := state.(*PullPushContextState); ok {
				if x.Offset.IsSet() {
					if off, err := stream.ParseOffsetAsInt(x.Offset); err == nil {
						startValue = off
					}
				}
			}
			for i := startValue; i < 10; i++ {
				time.Sleep(50 * time.Millisecond)
				et := 50 * time.Millisecond
				et *= time.Duration(i)
				err := ctx1.PushOut(&Record[int]{
					Key:       fmt.Sprintf("key-%d", i%2),
					Data:      i,
					EventTime: int64(et),
				})
				if err != nil {
					return fmt.Errorf("projection.Range: push: %w", err)
				}

				err = recovery.Snapshot(&PullPushContextState{
					Offset:    stream.MkOffsetFromInt(i),
					PullTopic: "",
					PushTopic: stream1.TopicName(),
				})
				if err != nil {
					return fmt.Errorf("projection.Range: snapshot: %w", err)
				}
			}
			return nil
		})
	assert.NoError(t, err)

	t.Log("WINDOW")

	recovery =
		NewRecoveryOptions(
			"recovery-window",
			&PullPushContextState{
				Offset:    nil,
				PullTopic: stream1.TopicName(),
				PushTopic: stream2.TopicName(),
			},
			store,
		).WithMaxRecoveryAttempts(recoveryAttempts)

	err = Recovery(
		recovery,
		func(state *PullPushContextState) (*PushAndPullInMemoryContext[int, float64], error) {
			ctx2 := NewPushAndPullInMemoryContext[int, float64](state, stream1, stream2)
			ctx2.SimulateRuntimeProblem(&SimulateProblem{
				ErrorOnPushOutProbability: probabilityOfFailure,
				ErrorOnPushOut:            fmt.Errorf("simulated push error"),
				ErrorOnPullInProbability:  probabilityOfFailure,
				ErrorOnPullIn:             fmt.Errorf("simulated pull error"),
			})
			return ctx2, nil
		},
		func(ctx *PushAndPullInMemoryContext[int, float64]) error {
			wd := &FixedWindow{Width: 200 * time.Millisecond}
			fm := &Discard{}
			td := &AtWatermark{}
			return DoWindow[int, float64](ctx, wd, fm, td, 0, func(x int, agg float64) (float64, error) {
				return float64(x) + agg, nil
			})
		})
	assert.NoError(t, err)

	t.Log("SINK")

	var orderOfEvents []string
	var orderOfUniquer = make(map[string]struct{})

	recovery =
		NewRecoveryOptions(
			"recovery-sink",
			&PullPushContextState{
				Offset:    nil,
				PullTopic: stream2.TopicName(),
				PushTopic: stream3.TopicName(),
			},
			store,
		).WithMaxRecoveryAttempts(recoveryAttempts)

	err = Recovery(
		recovery,
		func(state *PullPushContextState) (*PushAndPullInMemoryContext[float64, float64], error) {
			ctx3 := NewPushAndPullInMemoryContext[float64, float64](state, stream2, stream3)
			ctx3.SimulateRuntimeProblem(&SimulateProblem{
				ErrorOnPushOutProbability: probabilityOfFailure,
				ErrorOnPushOut:            fmt.Errorf("simulated push error"),
				ErrorOnPullInProbability:  probabilityOfFailure,
				ErrorOnPullIn:             fmt.Errorf("simulated pull error"),
			})
			return ctx3, nil
		},
		func(ctx *PushAndPullInMemoryContext[float64, float64]) error {
			for {
				val, err := ctx.PullIn()
				if err != nil {
					if errors.Is(err, stream.ErrEndOfStream) {
						return nil
					}
					return fmt.Errorf("projection.DoSink: pull: %w", err)
				}

				err = MatchDataR1(
					val,
					func(x *Record[float64]) error {
						time.Sleep(50 * time.Millisecond)

						entry := fmt.Sprintf("record-%s:%2f", x.Key, x.Data)
						if _, ok := orderOfUniquer[entry]; ok {
							return nil
						}
						orderOfUniquer[entry] = struct{}{}
						orderOfEvents = append(orderOfEvents, entry)
						return nil
					},
					func(x *Watermark[float64]) error {
						// TODO do snapshot
						return recovery.SnapshotFrom(ctx)
					},
				)
				if err != nil {
					return fmt.Errorf("projection.DoSink: sink: %w", err)
				}
			}
		})
	assert.NoError(t, err)

	expectedOrder := []string{
		"record-key-0:2.000000",
		"record-key-1:4.000000",
		"record-key-0:10.000000",
		"record-key-1:12.000000",
		"record-key-0:8.000000",
		"record-key-1:9.000000",
	}

	if diff := cmp.Diff(expectedOrder, orderOfEvents); diff != "" {
		t.Fatalf("diff: (-want +got)\n%s", diff)
	}

	results, err := store.FindingRecords(schemaless.FindingRecords[schemaless.Record[SnapshotState]]{
		RecordType: RecoveryRecordType,
	})
	assert.NoError(t, err)
	assert.Len(t, results.Items, 2)

	for _, result := range results.Items {
		t.Logf("recovery state: %v", result)
	}
}
