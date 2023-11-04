//generated by mkunion

/**
* This function is used to remove $type field from Expr, so it can be understood by mkunion, 
* that is assuming unions have one field in object, that is used to discriminate between variants.
*/
export function dediscriminateExpr(x: Expr): Expr {
    if (x["$type"] !== undefined) {
        delete x["$type"]
        return x
    }
    return x
}

/**
* This function is used to populate $type field in Expr, so it can be used as discriminative switch statement
* @example https://www.typescriptlang.org/play#example/discriminate-types
*/
export function discriminateExpr(x: Expr): Expr {
    if (x["$type"] === undefined) {
        let keyx = Object.keys(x)
        if (keyx.length === 1) {
            x["$type"] = keyx[0] as any
        }
    }
    return x
}

export type Expr = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.End",
	"workflow.End": End
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Assign",
	"workflow.Assign": Assign
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Apply",
	"workflow.Apply": Apply
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
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

//generated by mkunion

/**
* This function is used to remove $type field from Reshaper, so it can be understood by mkunion, 
* that is assuming unions have one field in object, that is used to discriminate between variants.
*/
export function dediscriminateReshaper(x: Reshaper): Reshaper {
    if (x["$type"] !== undefined) {
        delete x["$type"]
        return x
    }
    return x
}

/**
* This function is used to populate $type field in Reshaper, so it can be used as discriminative switch statement
* @example https://www.typescriptlang.org/play#example/discriminate-types
*/
export function discriminateReshaper(x: Reshaper): Reshaper {
    if (x["$type"] === undefined) {
        let keyx = Object.keys(x)
        if (keyx.length === 1) {
            x["$type"] = keyx[0] as any
        }
    }
    return x
}

export type Reshaper = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.GetValue",
	"workflow.GetValue": GetValue
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.SetValue",
	"workflow.SetValue": SetValue
}

export type GetValue = {
	Path?: string,
}

export type SetValue = {
	Value?: schema.Schema,
}

//generated by mkunion

/**
* This function is used to remove $type field from Predicate, so it can be understood by mkunion, 
* that is assuming unions have one field in object, that is used to discriminate between variants.
*/
export function dediscriminatePredicate(x: Predicate): Predicate {
    if (x["$type"] !== undefined) {
        delete x["$type"]
        return x
    }
    return x
}

/**
* This function is used to populate $type field in Predicate, so it can be used as discriminative switch statement
* @example https://www.typescriptlang.org/play#example/discriminate-types
*/
export function discriminatePredicate(x: Predicate): Predicate {
    if (x["$type"] === undefined) {
        let keyx = Object.keys(x)
        if (keyx.length === 1) {
            x["$type"] = keyx[0] as any
        }
    }
    return x
}

export type Predicate = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.And",
	"workflow.And": And
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Or",
	"workflow.Or": Or
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Not",
	"workflow.Not": Not
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
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

//generated by mkunion

/**
* This function is used to remove $type field from RunOption, so it can be understood by mkunion, 
* that is assuming unions have one field in object, that is used to discriminate between variants.
*/
export function dediscriminateRunOption(x: RunOption): RunOption {
    if (x["$type"] !== undefined) {
        delete x["$type"]
        return x
    }
    return x
}

/**
* This function is used to populate $type field in RunOption, so it can be used as discriminative switch statement
* @example https://www.typescriptlang.org/play#example/discriminate-types
*/
export function discriminateRunOption(x: RunOption): RunOption {
    if (x["$type"] === undefined) {
        let keyx = Object.keys(x)
        if (keyx.length === 1) {
            x["$type"] = keyx[0] as any
        }
    }
    return x
}

export type RunOption = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.ScheduleRun",
	"workflow.ScheduleRun": ScheduleRun
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.DelayRun",
	"workflow.DelayRun": DelayRun
}

export type ScheduleRun = {
	Interval?: string,
}

export type DelayRun = {
	DelayBySeconds?: number,
}

//generated by mkunion

/**
* This function is used to remove $type field from Command, so it can be understood by mkunion, 
* that is assuming unions have one field in object, that is used to discriminate between variants.
*/
export function dediscriminateCommand(x: Command): Command {
    if (x["$type"] !== undefined) {
        delete x["$type"]
        return x
    }
    return x
}

/**
* This function is used to populate $type field in Command, so it can be used as discriminative switch statement
* @example https://www.typescriptlang.org/play#example/discriminate-types
*/
export function discriminateCommand(x: Command): Command {
    if (x["$type"] === undefined) {
        let keyx = Object.keys(x)
        if (keyx.length === 1) {
            x["$type"] = keyx[0] as any
        }
    }
    return x
}

export type Command = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Run",
	"workflow.Run": Run
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Callback",
	"workflow.Callback": Callback
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.TryRecover",
	"workflow.TryRecover": TryRecover
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.StopSchedule",
	"workflow.StopSchedule": StopSchedule
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
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
}

export type StopSchedule = {
	RunID?: string,
}

export type ResumeSchedule = {
	RunID?: string,
}

//generated by mkunion

/**
* This function is used to remove $type field from State, so it can be understood by mkunion, 
* that is assuming unions have one field in object, that is used to discriminate between variants.
*/
export function dediscriminateState(x: State): State {
    if (x["$type"] !== undefined) {
        delete x["$type"]
        return x
    }
    return x
}

/**
* This function is used to populate $type field in State, so it can be used as discriminative switch statement
* @example https://www.typescriptlang.org/play#example/discriminate-types
*/
export function discriminateState(x: State): State {
    if (x["$type"] === undefined) {
        let keyx = Object.keys(x)
        if (keyx.length === 1) {
            x["$type"] = keyx[0] as any
        }
    }
    return x
}

export type State = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.NextOperation",
	"workflow.NextOperation": NextOperation
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Done",
	"workflow.Done": Done
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Error",
	"workflow.Error": Error
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Await",
	"workflow.Await": Await
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Scheduled",
	"workflow.Scheduled": Scheduled
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
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

//generated by mkunion

/**
* This function is used to remove $type field from Worflow, so it can be understood by mkunion, 
* that is assuming unions have one field in object, that is used to discriminate between variants.
*/
export function dediscriminateWorflow(x: Worflow): Worflow {
    if (x["$type"] !== undefined) {
        delete x["$type"]
        return x
    }
    return x
}

/**
* This function is used to populate $type field in Worflow, so it can be used as discriminative switch statement
* @example https://www.typescriptlang.org/play#example/discriminate-types
*/
export function discriminateWorflow(x: Worflow): Worflow {
    if (x["$type"] === undefined) {
        let keyx = Object.keys(x)
        if (keyx.length === 1) {
            x["$type"] = keyx[0] as any
        }
    }
    return x
}

export type Worflow = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "workflow.Flow",
	"workflow.Flow": Flow
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
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

export type BaseState = {
	Flow?: Worflow,
	RunID?: string,
	StepID?: string,
	Variables?: {[key: string]: any},
	ExprResult?: {[key: string]: any},
	DefaultMaxRetries?: number,
	RunOption?: RunOption,
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
	Variables?: {[key: string]: any},
}
export type ApplyAwaitOptions = {
	Timeout?: number,
}

//eslint-disable-next-line
import * as schema from './github_com_widmogrod_mkunion_x_schema'
