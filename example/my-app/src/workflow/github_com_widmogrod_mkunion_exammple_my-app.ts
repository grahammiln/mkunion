//generated by mkunion
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

export type ChatCMD = {
	"$type"?: "main.UserMessage",
	"main.UserMessage": UserMessage
}

export type UserMessage = {
	Message?: string,
}

export type GenerateImage = {
	Width?: number,
	Height?: number,
}
export type RefreshFlows = {}
export type Service<CMD, State> = {}
export type ListWorkflowsFn = {
	Count?: number,
	Words?: string[],
	EnumTest?: string,
}
export type RefreshStates = {}

//eslint-disable-next-line
import * as openai from './'
