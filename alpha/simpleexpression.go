package alpha

import (
	"reflect"
	"regexp"
)

const (
	EXPRESSION AlphaElement = iota - 1
	EFFECT
	CONDITIONAL
	SECTION
	ENTRY
)

type AlphaElement int

// type SimpleExpression interface {
// 	GetName() string       // Devuelve el nombre como una cadena.
// 	GetPattern() string    // Devuelve un patrón como una cadena.
// 	GetType() AlphaElement // Devuelve un tipo personalizado AlphaElement.
// 	ReturnType() reflect.Type
// }

type SimpleExpression struct {
	name        string
	pattern     string
	elementType AlphaElement
}

func (s *SimpleExpression) Base(elementType AlphaElement, pattern string, name string) {
	s.name = name
	s.pattern = pattern
	s.elementType = elementType
}

func (s *SimpleExpression) Init() *SimpleExpression {
	return s
}

func (s *SimpleExpression) PreCode(context Context, match regexp.Regexp) {

}

func (s *SimpleExpression) GetName() string {
	return s.name
}

func (s *SimpleExpression) GetPattern() string {
	return s.pattern
}

func (s *SimpleExpression) GetType() AlphaElement {
	return s.elementType
}

func (expr *SimpleExpression) ReturnType() reflect.Type {
	return nil
}
