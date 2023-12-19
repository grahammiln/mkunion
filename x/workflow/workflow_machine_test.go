package workflow

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/widmogrod/mkunion/x/machine"
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/storage/schemaless"
	"github.com/widmogrod/mkunion/x/storage/schemaless/typedful"
	"testing"
	"time"
)

var functions = map[string]Function{
	"concat": func(body *FunctionInput) (*FunctionOutput, error) {
		args := body.Args
		a, ok := schema.As[string](args[0])
		if !ok {
			return nil, fmt.Errorf("expected string, got %T", args[0])
		}
		b, ok := schema.As[string](args[1])
		if !ok {
			return nil, fmt.Errorf("expected string, got %T", args[1])
		}

		return &FunctionOutput{
			Result: schema.MkString(a + b),
		}, nil
	},
}

func TestExecution(t *testing.T) {
	program := &Flow{
		Name: "hello_world_flow",
		Arg:  "input",
		Body: []Expr{
			&Assign{
				ID:    "assign1",
				VarOk: "res",
				Val: &Apply{ID: "apply1", Name: "concat", Args: []Reshaper{
					&SetValue{Value: schema.MkString("hello ")},
					&GetValue{Path: "input"},
				}},
			},
			&End{
				ID:     "end1",
				Result: &GetValue{Path: "res"},
			},
		},
	}

	runId := "1"

	di := &DI{
		FindWorkflowF: func(flowID string) (*Flow, error) {
			return program, nil
		},
		FindFunctionF: func(funcID string) (Function, error) {
			if fn, ok := functions[funcID]; ok {
				return fn, nil
			}

			return nil, fmt.Errorf("function %s not found", funcID)
		},
		GenerateRunIDF: func() string {
			return runId
		},
	}

	store := schemaless.NewInMemoryRepository[schema.Schema]()
	repo := typedful.NewTypedRepository[State](store)
	state, err := repo.Get("1", "workflow")
	assert.ErrorIs(t, err, schemaless.ErrNotFound)

	work := NewMachine(di, state.Data)
	err = work.Handle(&Run{
		Flow:  &FlowRef{FlowID: "hello_world_flow"},
		Input: schema.MkString("world"),
	})
	assert.NoError(t, err)

	newState := work.State()
	err = repo.UpdateRecords(schemaless.Save(schemaless.Record[State]{
		ID:   "1",
		Type: "workflow",
		Data: newState,
	}))
	assert.NoError(t, err)

	assert.Equal(t, &Done{
		Result: schema.MkString("hello world"),
		BaseState: BaseState{
			RunID:  runId,
			StepID: "end1",
			Flow:   &FlowRef{FlowID: "hello_world_flow"},
			Variables: map[string]schema.Schema{
				"input": schema.MkString("world"),
				"res":   schema.MkString("hello world"),
			},
			ExprResult:        make(map[string]schema.Schema),
			DefaultMaxRetries: 3,
		},
	}, newState)

	state, err = repo.Get("1", "workflow")
	assert.NoError(t, err)

	work = NewMachine(di, state.Data)
	err = work.Handle(&Run{
		Flow:  &FlowRef{FlowID: "hello_world_flow"},
		Input: schema.MkString("world"),
	})
	assert.ErrorAs(t, err, &ErrStateReachEnd)
}

func TestMachine(t *testing.T) {
	program := &Flow{
		Name: "hello_world_flow",
		Arg:  "input",
		Body: []Expr{
			&Assign{
				VarOk: "res",
				Val: &Apply{Name: "concat", Args: []Reshaper{
					&SetValue{Value: schema.MkString("hello ")},
					&GetValue{Path: "input"},
				}},
			},
			&End{
				Result: &GetValue{Path: "res"},
			},
		},
	}

	callbackID := "callback1"
	runID := "123"

	program_await := &Flow{
		Name: "hello_world_flow_await",
		Arg:  "input",
		Body: []Expr{
			&Assign{
				VarOk: "res",
				Val: &Apply{
					Name: "concat",
					Args: []Reshaper{
						&SetValue{Value: schema.MkString("hello ")},
						&GetValue{Path: "input"},
					},
					Await: &ApplyAwaitOptions{
						Timeout: int64(time.Second * 10),
					},
				},
			},
			&End{
				Result: &GetValue{Path: "res"},
			},
		},
	}

	program_if := &Flow{
		Name: "hello_world_flow_if",
		Arg:  "input",
		Body: []Expr{
			&Assign{
				VarOk: "res",
				Val: &Apply{ID: "apply1", Name: "concat", Args: []Reshaper{
					&SetValue{Value: schema.MkString("hello ")},
					&GetValue{Path: "input"},
				}},
			},
			&Choose{
				If: &Compare{
					Operation: "=",
					Left:      &GetValue{Path: "res"},
					Right:     &SetValue{Value: schema.MkString("hello world")},
				},
				Then: []Expr{
					&End{
						Result: &GetValue{Path: "res"},
					},
				},
				Else: []Expr{
					&End{
						Result: &SetValue{Value: schema.MkString("only Spanish will work!")},
					},
				},
			},
		},
	}

	timeNow := time.Now()
	di := &DI{
		FindWorkflowF: func(flowID string) (*Flow, error) {
			switch flowID {
			case "hello_world_flow":
				return program, nil
			case "hello_world_flow_await":
				return program_await, nil
			case "hello_world_flow_if":
				return program_if, nil
			}
			return nil, fmt.Errorf("flow %s not found", flowID)
		},
		FindFunctionF: func(funcID string) (Function, error) {
			if fn, ok := functions[funcID]; ok {
				return fn, nil
			}

			return nil, fmt.Errorf("function %s not found", funcID)
		},

		GenerateCallbackIDF: func() string {
			return callbackID
		},

		GenerateRunIDF: func() string {
			return runID
		},
		MockTimeNow: &timeNow,
	}

	suite := machine.NewTestSuite(func() *machine.Machine[Command, State] {
		return NewMachine(di, nil)
	})

	suite.Case("start execution", func(c *machine.Case[Command, State]) {
		c.
			GivenCommand(&Run{
				Flow:  &FlowRef{FlowID: "hello_world_flow"},
				Input: schema.MkString("world"),
			}).
			ThenState(&Done{
				Result: schema.MkString("hello world"),
				BaseState: BaseState{
					RunID:  runID,
					StepID: "end",
					Flow:   &FlowRef{FlowID: "hello_world_flow"},
					Variables: map[string]schema.Schema{
						"input": schema.MkString("world"),
						"res":   schema.MkString("hello world"),
					},
					ExprResult:        make(map[string]schema.Schema),
					DefaultMaxRetries: 3,
				},
			})
	})
	suite.Case("start scheduled execution delay 10s", func(c *machine.Case[Command, State]) {
		c.
			GivenCommand(&Run{
				Flow:  &FlowRef{FlowID: "hello_world_flow"},
				Input: schema.MkString("world"),
				RunOption: &DelayRun{
					DelayBySeconds: 10,
				},
			}).
			ThenState(&Scheduled{
				ParentRunID:          runID,
				ExpectedRunTimestamp: di.TimeNow().Add(time.Duration(10) * time.Second).Unix(),
				BaseState: BaseState{
					RunID:  runID,
					StepID: "",
					Flow:   &FlowRef{FlowID: "hello_world_flow"},
					Variables: map[string]schema.Schema{
						"input": schema.MkString("world"),
					},
					ExprResult:        make(map[string]schema.Schema),
					DefaultMaxRetries: 3,
					RunOption: &DelayRun{
						DelayBySeconds: 10,
					},
				},
			}).
			ForkCase("resume execution", func(c *machine.Case[Command, State]) {
				c.
					GivenCommand(&Run{}).
					ThenState(&Done{
						Result: schema.MkString("hello world"),
						BaseState: BaseState{
							RunID:  runID,
							StepID: "end",
							Flow:   &FlowRef{FlowID: "hello_world_flow"},
							Variables: map[string]schema.Schema{
								"input": schema.MkString("world"),
								"res":   schema.MkString("hello world"),
							},
							ExprResult:        make(map[string]schema.Schema),
							DefaultMaxRetries: 3,
							RunOption: &DelayRun{
								DelayBySeconds: 10,
							},
						},
					})
			}).
			ForkCase("stop execution", func(c *machine.Case[Command, State]) {
				c.
					GivenCommand(&StopSchedule{
						ParentRunID: runID,
					}).
					ThenState(&ScheduleStopped{
						ParentRunID: runID,
						BaseState: BaseState{
							RunID:  runID,
							StepID: "",
							Flow:   &FlowRef{FlowID: "hello_world_flow"},
							Variables: map[string]schema.Schema{
								"input": schema.MkString("world"),
							},
							ExprResult:        make(map[string]schema.Schema),
							DefaultMaxRetries: 3,
							RunOption: &DelayRun{
								DelayBySeconds: 10,
							},
						},
					}).
					GivenCommand(&ResumeSchedule{
						ParentRunID: runID,
					}).
					ThenState(&Scheduled{
						ParentRunID:          runID,
						ExpectedRunTimestamp: di.TimeNow().Add(time.Duration(10) * time.Second).Unix(),
						BaseState: BaseState{
							RunID:  runID,
							StepID: "",
							Flow:   &FlowRef{FlowID: "hello_world_flow"},
							Variables: map[string]schema.Schema{
								"input": schema.MkString("world"),
							},
							ExprResult:        make(map[string]schema.Schema),
							DefaultMaxRetries: 3,
							RunOption: &DelayRun{
								DelayBySeconds: 10,
							},
						},
					})
			})
	})
	suite.Case("start execution that awaits for callback", func(c *machine.Case[Command, State]) {
		c.
			GivenCommand(&Run{
				Flow:  &FlowRef{FlowID: "hello_world_flow_await"},
				Input: schema.MkString("world"),
			}).
			ThenState(&Await{
				Timeout:    int64(10 * time.Second),
				CallbackID: callbackID,
				BaseState: BaseState{
					RunID:  runID,
					StepID: "apply-concat",
					Flow:   &FlowRef{FlowID: "hello_world_flow_await"},
					Variables: map[string]schema.Schema{
						"input": schema.MkString("world"),
					},
					ExprResult:        make(map[string]schema.Schema),
					DefaultMaxRetries: 3,
				},
			}).
			ForkCase("callback received", func(c *machine.Case[Command, State]) {
				// Assuming that callback is received before timeout.
				c.
					GivenCommand(&Callback{
						CallbackID: callbackID,
						Result:     schema.MkString("hello + world"),
					}).
					ThenState(&Done{
						Result: schema.MkString("hello + world"),
						BaseState: BaseState{
							RunID:  runID,
							StepID: "end",
							Flow:   &FlowRef{FlowID: "hello_world_flow_await"},
							Variables: map[string]schema.Schema{
								"input": schema.MkString("world"),
								"res":   schema.MkString("hello + world"),
							},
							ExprResult: map[string]schema.Schema{
								"apply-concat": schema.MkString("hello + world"),
							},
							DefaultMaxRetries: 3,
						},
					})
			}).
			ForkCase("received invalid callbackID", func(c *machine.Case[Command, State]) {
				c.
					GivenCommand(&Callback{
						CallbackID: "invalid_callback_id",
						Result:     schema.MkString("hello + world"),
					}).
					ThenStateAndError(&Await{
						Timeout:    int64(10 * time.Second),
						CallbackID: callbackID,
						BaseState: BaseState{
							RunID:  runID,
							StepID: "apply-concat",
							Flow:   &FlowRef{FlowID: "hello_world_flow_await"},
							Variables: map[string]schema.Schema{
								"input": schema.MkString("world"),
							},
							ExprResult:        make(map[string]schema.Schema),
							DefaultMaxRetries: 3,
						},
					}, ErrCallbackNotMatch)
			})
	})
	suite.Case("start execution no input variable", func(c *machine.Case[Command, State]) {
		c.
			GivenCommand(&Run{
				Flow: &FlowRef{FlowID: "hello_world_flow"},
			}).
			ThenState(&Error{
				Code:   "function-execution",
				Reason: "function concat() returned error: expected string, got <nil>",
				BaseState: BaseState{
					RunID:  runID,
					StepID: "apply-concat",
					Flow:   &FlowRef{FlowID: "hello_world_flow"},
					Variables: map[string]schema.Schema{
						"input": nil,
					},
					ExprResult:        map[string]schema.Schema{},
					DefaultMaxRetries: 3,
				},
			})
	})
	suite.Case("start execution fails on non existing flowID", func(c *machine.Case[Command, State]) {
		c.
			GivenCommand(&Run{
				Flow:  &FlowRef{FlowID: "hello_world_flow_non_existing"},
				Input: schema.MkString("world"),
			}).
			ThenStateAndError(nil, fmt.Errorf("flow hello_world_flow_non_existing not found"))
	})
	suite.Case("start execution fails on function retrival", func(c *machine.Case[Command, State]) {
		c.
			GivenCommand(&Run{
				Flow:  &FlowRef{FlowID: "hello_world_flow"},
				Input: schema.MkString("world"),
			}, machine.WithBefore(func() {
				di.FindFunctionF = func(funcID string) (Function, error) {
					return nil, fmt.Errorf("function funcID='%s' not found", funcID)
				}
			}), machine.WithAfter(func() {
				di.FindFunctionF = func(funcID string) (Function, error) {
					if fn, ok := functions[funcID]; ok {
						return fn, nil
					}

					return nil, fmt.Errorf("function %s not found", funcID)
				}
			})).
			ThenState(&Error{
				Code:   "function-missing",
				Reason: "function concat() not found, details: function funcID='concat' not found",
				BaseState: BaseState{
					RunID:  runID,
					StepID: "apply-concat",
					Flow:   &FlowRef{FlowID: "hello_world_flow"},
					Variables: map[string]schema.Schema{
						"input": schema.MkString("world"),
					},
					ExprResult:        map[string]schema.Schema{},
					DefaultMaxRetries: 3,
				},
			}).
			ForkCase("retry execution", func(c *machine.Case[Command, State]) {
				c.
					GivenCommand(&TryRecover{},
						machine.WithBefore(func() {
							di.FindFunctionF = func(funcID string) (Function, error) {
								return nil, fmt.Errorf("function funcID='%s' not found", funcID)
							}
						}), machine.WithAfter(func() {
							di.FindFunctionF = func(funcID string) (Function, error) {
								if fn, ok := functions[funcID]; ok {
									return fn, nil
								}

								return nil, fmt.Errorf("function %s not found", funcID)
							}
						}),
					).
					ThenState(&Error{
						Code:    "function-missing",
						Reason:  "function concat() not found, details: function funcID='concat' not found",
						Retried: 1,
						BaseState: BaseState{
							RunID:  runID,
							StepID: "apply-concat",
							Flow:   &FlowRef{FlowID: "hello_world_flow"},
							Variables: map[string]schema.Schema{
								"input": schema.MkString("world"),
							},
							ExprResult:        map[string]schema.Schema{},
							DefaultMaxRetries: 3,
						},
					})
			})
	})
	suite.Case("execute function with if statement", func(c *machine.Case[Command, State]) {
		c.
			GivenCommand(&Run{
				Flow:  &FlowRef{FlowID: "hello_world_flow_if"},
				Input: schema.MkString("El Mundo"),
			}).
			ThenState(&Done{
				Result: schema.MkString("only Spanish will work!"),
				BaseState: BaseState{
					RunID:  runID,
					StepID: "end-1",
					Flow:   &FlowRef{FlowID: "hello_world_flow_if"},
					Variables: map[string]schema.Schema{
						"input": schema.MkString("El Mundo"),
						"res":   schema.MkString("hello El Mundo"),
					},
					ExprResult:        map[string]schema.Schema{},
					DefaultMaxRetries: 3,
				},
			})
	})

	suite.Run(t)
	suite.Fuzzy(t)

	if true || suite.AssertSelfDocumentStateDiagram(t, "machine") {
		suite.SelfDocumentStateDiagram(t, "machine")
	}
}
