package projection

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/widmogrod/mkunion/x/stream"
	"testing"
)

func TestProjection(t *testing.T) {
	out1 := stream.NewInMemoryStream[int](stream.WithSystemTime)
	ctx1 := NewPushOnlyInMemoryContext[int](out1)

	err := DoLoad(ctx1, func(push func(Data[int]) error) error {
		for i := 0; i < 10; i++ {
			err := push(&Record[int]{
				Key:  fmt.Sprintf("key-%d", i),
				Data: i,
			})
			if err != nil {
				return fmt.Errorf("projection.Range: push: %w", err)
			}
		}
		return nil
	})
	assert.NoError(t, err)

	out2 := stream.NewInMemoryStream[float64](stream.WithSystemTime)
	ctx2 := NewPushAndPullInMemoryContext[int, float64](out1, out2)
	err = DoMap[int, float64](ctx2, func(x Data[int]) Data[float64] {
		return MatchDataR1(
			x,
			func(x *Record[int]) Data[float64] {
				return &Record[float64]{
					Key:  x.Key,
					Data: float64(x.Data) * 2,
				}
			},
			func(x *Watermark[int]) Data[float64] {
				return &Watermark[float64]{
					EventTime: x.EventTime,
				}
			},
		)
	})
	assert.NoError(t, err)

	orderOfEvents := []string{}
	ctx4 := DoJoin[int, float64](out1, out2)
	err = DoSink(ctx4, func(x Data[Either[int, float64]]) error {
		return MatchDataR1(
			x,
			func(x *Record[Either[int, float64]]) error {
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
			},
			func(x *Watermark[Either[int, float64]]) error {
				orderOfEvents = append(orderOfEvents, fmt.Sprintf("watermark-%d", x.EventTime))
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
		t.Fatalf("DoJoin: diff: (-want +got)\n%s", diff)
	}
}
