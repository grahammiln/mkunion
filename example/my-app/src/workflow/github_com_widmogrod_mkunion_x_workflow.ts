//generated by mkunion
export type Predicate = {
	"$type"?: "workflow.And",
	"workflow.And": And
} | {
	"$type"?: "workflow.Or",
	"workflow.Or": Or
} | {
	"$type"?: "workflow.Not",
	"workflow.Not": Not
} | {
	"$type"?: "workflow.Compare",
	"workflow.Compare": Compare
}

export type And = {
	L?: Predicate[],
}

export type Or = {
	L?: Predicate[],
}

export type Not = {
	P?: Predicate,
}

export type Compare = {
	Operation?: string,
	Left?: Reshaper,
	Right?: Reshaper,
}

export type RunOption = {
	"$type"?: "workflow.ScheduleRun",
	"workflow.ScheduleRun": ScheduleRun
} | {
	"$type"?: "workflow.DelayRun",
	"workflow.DelayRun": DelayRun
}

export type ScheduleRun = {
	Interval?: string,
	ParentRunID?: string,
}

export type DelayRun = {
	DelayBySeconds?: number,
}

export type Command = {
	"$type"?: "workflow.Run",
	"workflow.Run": Run
} | {
	"$type"?: "workflow.Callback",
	"workflow.Callback": Callback
} | {
	"$type"?: "workflow.TryRecover",
	"workflow.TryRecover": TryRecover
} | {
	"$type"?: "workflow.StopSchedule",
	"workflow.StopSchedule": StopSchedule
} | {
	"$type"?: "workflow.ResumeSchedule",
	"workflow.ResumeSchedule": ResumeSchedule
}

export type Run = {
	Flow?: Worflow,
	Input?: schema.Schema,
	RunOption?: RunOption,
}

export type Callback = {
	CallbackID?: string,
	Result?: schema.Schema,
}

export type TryRecover = {
	RunID?: string,
}

export type StopSchedule = {
	ParentRunID?: string,
}

export type ResumeSchedule = {
	ParentRunID?: string,
}

export type State = {
	"$type"?: "workflow.NextOperation",
	"workflow.NextOperation": NextOperation
} | {
	"$type"?: "workflow.Done",
	"workflow.Done": Done
} | {
	"$type"?: "workflow.Error",
	"workflow.Error": Error
} | {
	"$type"?: "workflow.Await",
	"workflow.Await": Await
} | {
	"$type"?: "workflow.Scheduled",
	"workflow.Scheduled": Scheduled
} | {
	"$type"?: "workflow.ScheduleStopped",
	"workflow.ScheduleStopped": ScheduleStopped
}

export type NextOperation = {
	Result?: schema.Schema,
	BaseState?: BaseState,
}

export type Done = {
	Result?: schema.Schema,
	BaseState?: BaseState,
}

export type Error = {
	Code?: string,
	Reason?: string,
	Retried?: number,
	BaseState?: BaseState,
}

export type Await = {
	CallbackID?: string,
	Timeout?: number,
	BaseState?: BaseState,
}

export type Scheduled = {
	ExpectedRunTimestamp?: number,
	BaseState?: BaseState,
}

export type ScheduleStopped = {
	BaseState?: BaseState,
}

export type Worflow = {
	"$type"?: "workflow.Flow",
	"workflow.Flow": Flow
} | {
	"$type"?: "workflow.FlowRef",
	"workflow.FlowRef": FlowRef
}

export type Flow = {
	Name?: string,
	Arg?: string,
	Body?: Expr[],
}

export type FlowRef = {
	FlowID?: string,
}

export type Expr = {
	"$type"?: "workflow.End",
	"workflow.End": End
} | {
	"$type"?: "workflow.Assign",
	"workflow.Assign": Assign
} | {
	"$type"?: "workflow.Apply",
	"workflow.Apply": Apply
} | {
	"$type"?: "workflow.Choose",
	"workflow.Choose": Choose
}

export type End = {
	ID?: string,
	Result?: Reshaper,
}

export type Assign = {
	ID?: string,
	VarOk?: string,
	VarErr?: string,
	Val?: Expr,
}

export type Apply = {
	ID?: string,
	Name?: string,
	Args?: Reshaper[],
	Await?: ApplyAwaitOptions,
}

export type Choose = {
	ID?: string,
	If?: Predicate,
	Then?: Expr[],
	Else?: Expr[],
}

export type Reshaper = {
	"$type"?: "workflow.GetValue",
	"workflow.GetValue": GetValue
} | {
	"$type"?: "workflow.SetValue",
	"workflow.SetValue": SetValue
}

export type GetValue = {
	Path?: string,
}

export type SetValue = {
	Value?: schema.Schema,
}

export type BaseState = {
	Flow?: Worflow,
	RunID?: string,
	StepID?: string,
	Variables?: {[key: string]: schema.Schema},
	ExprResult?: {[key: string]: schema.Schema},
	DefaultMaxRetries?: number,
	RunOption?: RunOption,
}
export type ApplyAwaitOptions = {
	Timeout?: number,
}
export type ResumeOptions = {
	Timeout?: number,
}
export type Execution = {
	FlowID?: string,
	Status?: State,
	Location?: string,
	StartTime?: number,
	EndTime?: number,
	Variables?: {[key: string]: schema.Schema},
}
export type FunctionInput = {
	Name?: string,
	CallbackID?: string,
	Args?: schema.Schema[],
}
export type FunctionOutput = {
	Result?: schema.Schema,
}

//eslint-disable-next-line
import * as schema from './github_com_widmogrod_mkunion_x_schema'
