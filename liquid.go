/*
Package liquid is a pure Go implementation of Shopify Liquid templates, developed for use in https://github.com/osteele/gojekyll.

See the project README https://github.com/osteele/liquid for additional information and implementation status.


The liquid package itself is versioned in gopkg.in. Subpackages have no compatibility guarantees. Except where specifically documented, the “public” entities of subpackages are intended only for use by the liquid package and its subpackages.
*/
package liquid

import (
	"github.com/osteele/liquid/evaluator"
	"github.com/osteele/liquid/expressions"
	"github.com/osteele/liquid/parser"
	"github.com/osteele/liquid/render"
)

// Bindings is a map of variable names to values.
//
// Clients need not use this type. It is used solely for documentation. Callers can use instances
// of map[string]interface{} itself as argument values to functions declared with this parameter type.
type Bindings map[string]interface{}

// A Renderer returns the rendered string for a block. This is the type of a tag definition.
//
// See the examples at Engine.RegisterTag and Engine.RegisterBlock.
type Renderer func(render.Context) (string, error)

// SourceError records an error with a source location and optional cause.
type SourceError interface {
	error
	Cause() error
	Path() string
	LineNumber() int
}

// IsTemplateError returns true iff the error represents an error in the template
// syntax or execution. It is used to distinguish errors in input values from errors in the Liquid implementation, or the implementation of tags and filters, themselves.
//
// Use this function to avoid coding the specific types of subpackage errors, which
// are likely to change.
func IsTemplateError(err error) bool {
	switch err.(type) {
	// TODO some of these clauses, or maybe the entire function, is unnecessary
	// now that interface calls return SourceError
	case evaluator.TypeError:
		return true
	case expressions.InterpreterError:
		return true
	case expressions.ParseError:
		return true
	case parser.Error:
		return true
	case render.Error:
		return true
	case SourceError:
		return true
	default:
		return false
	}
}
