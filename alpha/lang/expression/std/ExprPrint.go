package std

import (
	"regexp"

	"github.com/angelgamer29ja/alpha/alpha"
)

type ExprPrint struct {
	alpha.SimpleExpression
}

func (expr *ExprPrint) Init() *ExprPrint {
	expr.Base(alpha.EXPRESSION, "print %string%", "print")
	return expr
}

func (expr *ExprPrint) PreCode(context alpha.Context, pattern regexp.Regexp) {

}
