package parser

/*
Represents a function that can be called from within an expression.
This method must return an error if, for any reason, it is unable to produce exactly one unambiguous result.
An error returned will halt execution of the expression.
*/
type ExpressionFunction struct {
	Name       string
	Parameters []string
	ReturnType string
}
