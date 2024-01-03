//generated by mkunion
export type ChatCMD = {
	"$type"?: "main.UserMessage",
	"main.UserMessage": UserMessage
}

export type UserMessage = {
	Message?: string,
}

export type ChatResult = {
	"$type"?: "main.SystemResponse",
	"main.SystemResponse": SystemResponse
} | {
	"$type"?: "main.UserResponse",
	"main.UserResponse": UserResponse
} | {
	"$type"?: "main.ChatResponses",
	"main.ChatResponses": ChatResponses
}

export type SystemResponse = {
	Message?: string,
	ToolCalls?: openai.ToolCall[],
}

export type UserResponse = {
	Message?: string,
}

export type ChatResponses = {
	Responses?: ChatResult[],
}

export type ListWorkflowsFn = {
	Count?: number,
	Words?: string[],
	EnumTest?: string,
}

export type PageResult = schemaless.PageResult<schemaless.Record<any>>

export type State = workflow.State

export type Command = workflow.Command

export type RefreshStates = {}

export type Service<CMD, State> = {}

export type UpdateRecords = schemaless.UpdateRecords<schemaless.Record<any>>

export type Expr = workflow.Expr

export type RefreshFlows = {}

export type Workflow = workflow.Workflow

export type Predicate = workflow.Predicate

export type FunctionInput = workflow.FunctionInput

export type GenerateImage = {
	Width?: number,
	Height?: number,
}

export type Reshaper = workflow.Reshaper

export type FindRecords = schemaless.FindingRecords<schemaless.Record<any>>

export type Schema = schema.Schema

export type FunctionOutput = workflow.FunctionOutput


//eslint-disable-next-line
import * as openai from './github_com_sashabaranov_go-openai'
//eslint-disable-next-line
import * as schemaless from './github_com_widmogrod_mkunion_x_storage_schemaless'
//eslint-disable-next-line
import * as workflow from './github_com_widmogrod_mkunion_x_workflow'
//eslint-disable-next-line
import * as schema from './github_com_widmogrod_mkunion_x_schema'
