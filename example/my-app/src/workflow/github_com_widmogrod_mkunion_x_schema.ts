//generated by mkunion
export type Schema = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "schema.None",
	"schema.None": None
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "schema.Bool",
	"schema.Bool": boolean
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "schema.Number",
	"schema.Number": number
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "schema.String",
	"schema.String": string
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "schema.Binary",
	"schema.Binary": Binary
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "schema.List",
	"schema.List": List
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "schema.Map",
	"schema.Map": Map
}

export type None = {
}

export type Binary = {
	B?: string,
}

export type List = {
	Items?: Schema[],
}

export type Map = {
	Field?: Field[],
}

export type Field = {
	Name?: string,
	Value?: Schema,
}
