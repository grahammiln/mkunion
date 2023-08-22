package workflow

import (
	"fmt"
	"github.com/widmogrod/mkunion/x/schema"
)

type (
	Context struct {
		Functions map[string]Function
		Variables map[string]schema.Schema // variables are immutable, once set they can't be changed

		Result schema.Schema
		Name   string

		Root *Context
		Prev *Context

		ExecutionPath      []string                  // holds information how execution flow goes, what variables were set, and result of execution
		ExecutionVariables map[string]string         // holds what variables were set
		ExecutionInfo      map[string]*ExecutionInfo // holds information about execution
	}

	Function func(args []schema.Schema) (schema.Schema, error)

	ExecutionInfo struct {
		SetVariables []string
		DidFail      bool
		Retried      int
	}
)

func (c *Context) Errorf(format string, a ...any) error {
	return fmt.Errorf(fmt.Sprintf("%s: %s", c.Name, format), a...)
}

func (c *Context) GetForFlow(x Worflow) *Context {
	c.Root.ExecutionPath = append(c.Root.ExecutionPath, c.Name)
	//c.Name = c.Name + ".Flow"

	return &Context{
		Root: c,
		Prev: c,
		Name: c.Name + ".Flow",
	}
}

func (c *Context) GetForReshaper(x Reshaper) *Context {
	name := MustMatchReshaper(
		x,
		func(x *GetValue) string {
			return ".GetValue"
		},
		func(x *SetValue) string {
			return ".SetValue"
		},
	)

	c.Root.ExecutionPath = append(c.Root.ExecutionPath, c.Name)

	return &Context{
		Root: c.Root,
		Prev: c,
		Name: c.Name + name,
	}
}

func (c *Context) GetForExpr(x Expr) *Context {
	name := MustMatchExpr(
		x,
		func(x *End) string {
			return ".End"
		},
		func(x *Assign) string {
			return ".Assign"
		},
		func(x *Apply) string {
			return ".Apply"
		},
		func(x *Choose) string {
			return ".Choose"
		},
	)

	c.Root.ExecutionPath = append(c.Root.ExecutionPath, c.Name)

	return &Context{
		Root: c.Root,
		Prev: c,
		Name: c.Name + name,
	}
}

func (c *Context) GetVariable(name string) (schema.Schema, bool) {
	ctx := c
	for {
		value, ok := ctx.Variables[name]
		if ok {
			return value, ok
		}

		if ctx.Prev == nil {
			return nil, false
		}

		ctx = ctx.Prev
	}
}

func (c *Context) SetVariable(name string, val schema.Schema) error {
	if _, ok := c.GetVariable(name); ok {
		return c.Errorf("variable %s already set", name)
	}

	c.Root.Variables[name] = val
	c.Root.ExecutionVariables[name] = c.Name

	return nil
}

//func (c *Context) RecordExecutionInfo(info *ExecutionInfo) {
//	c.ExecutionInfo[c.Name] = info
//}

type ExecutionStack struct {
	Stack []ASTNode
	//Executed []Result{}
}

func (s *ExecutionStack) Push(node ASTNode) {
	s.Stack = append(s.Stack, node)
}

func (s *ExecutionStack) Pop() ASTNode {
	if len(s.Stack) == 0 {
		return nil
	}

	node := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	return node
}

func (s *ExecutionStack) Peek() ASTNode {
	if len(s.Stack) == 0 {
		return nil
	}

	return s.Stack[len(s.Stack)-1]
}

func (s *ExecutionStack) Clear() {
	s.Stack = nil
}

func ExprID(expr Expr) string {
	return MustMatchExpr(
		expr,
		func(x *End) string {
			return x.ID
		},
		func(x *Assign) string {
			return x.ID
		},
		func(x *Apply) string {
			return x.ID
		},
		func(x *Choose) string {
			return x.ID
		})
}

func GetExprFromProgram(program Worflow, id string) Expr {
	return MustMatchWorflow(
		program,
		func(x *Flow) Expr {
			for _, expr := range x.Body {
				if ExprID(expr) == id {
					return expr
				}

				if expr := GetExprFromExpr(expr, id); expr != nil {
					return expr
				}
			}

			return nil
		},
	)
}

func GetExprFromExpr(expr Expr, id string) Expr {
	return MustMatchExpr(
		expr,
		func(x *End) Expr {
			if x.ID == id {
				return x
			}
			return nil
		},
		func(x *Assign) Expr {
			if x.ID == id {
				return x
			}

			return GetExprFromExpr(x.Val, id)
		},
		func(x *Apply) Expr {
			if x.ID == id {
				return x
			}

			return nil

		},
		func(x *Choose) Expr {
			if x.ID == id {
				return x
			}

			for _, expr := range x.Then {
				if expr := GetExprFromExpr(expr, id); expr != nil {
					return expr
				}
			}

			for _, expr := range x.Else {
				if expr := GetExprFromExpr(expr, id); expr != nil {
					return expr
				}
			}

			return nil
		},
	)
}

//func ExecuteStatus(expr Expr, context *Context) Status {
//	return MustMatchExpr(
//		expr,
//		func(x *End) Status {
//			if x.Fail != nil {
//				val, err := ExecuteReshaper(context, x.Result)
//				if err != nil {
//					return &Error{
//						Code:   "1",
//						Reason: fmt.Errorf("failed to execute node=%s: %w", x.ID, err).Error(),
//					}
//				}
//
//				return &Fail{
//					Result: val,
//				}
//			}
//
//			val, err := ExecuteReshaper(context, x.Result)
//			if err != nil {
//				return &Error{
//					Code:   "2",
//					Reason: fmt.Errorf("failed to execute node=%s: %w", x.ID, err).Error(),
//				}
//			}
//
//			return &Done{
//				Result: val,
//			}
//		},
//		func(x *Assign) Status {
//			val := ExecuteStatus(x.Val, context)
//
//			if err != nil {
//				return nil, context.Errorf("failed to execute flow: %w", err)
//			}
//
//			return context.SetVariable(x.Var, val)
//		},
//		func(x *Apply) Status {
//			args := make([]schema.Schema, len(x.Args))
//			for i, arg := range x.Args {
//				val, err := ExecuteReshaper(context, arg)
//				if err != nil {
//					return &Error{
//						Code:   "3",
//						Reason: fmt.Errorf("failed to execute node=%s: %w", x.ID, err).Error(),
//					}
//				}
//				args[i] = val
//			}
//
//			if fn, ok := context.Root.Functions[x.Name]; ok {
//				val, err := fn(args)
//				if err != nil {
//					return &Error{
//						Code:   "4",
//						Reason: fmt.Errorf("failed to execute node=%s: %w", x.ID, err).Error(),
//					}
//				}
//				return &Finished{
//					Result: val,
//				}
//			} else {
//				return &Error{
//					Code:   "5",
//					Reason: fmt.Errorf("failed to execute node=%s: no function", x.ID).Error(),
//				}
//			}
//		},
//		func(x *Choose) Status {
//			return nil, context.Errorf("not implemented")
//		},
//	)
//}

//func ExecuteNode(context *Context, stack *ExecutionStack, node ASTNode) ([]ASTNode, error) {
//	return MustMatchASTNodeR2(
//		node,
//		func(x *Flow) ([]ASTNode, error) {
//			for _, expr := range x.Body {
//				stack.Push(expr.(ASTNode))
//			}
//
//			return x.Body, nil
//		},
//		func(x *End) ([]ASTNode, error) {
//			if x.Fail != nil {
//				return nil, context.Errorf("execution failed: %s", x.Fail)
//			}
//
//			res, err := ExecuteReshaper(context, x.Result)
//			if err != nil {
//				return nil, context.Errorf("failed to execute reshaper: %w", err)
//			}
//
//			context.Result = res
//			stack.Clear()
//			return nil, nil
//		},
//		func(x *Assign) ([]ASTNode, error) {
//
//		},
//		func(x *Apply) ([]ASTNode, error) {
//
//		},
//		func(x *Choose) ([]ASTNode, error) {
//
//		},
//		func(x *GetValue) ([]ASTNode, error) {
//
//		},
//		func(x *SetValue) ([]ASTNode, error) {
//
//		},
//	)
//}
