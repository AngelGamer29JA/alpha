package alpha

import (
	"fmt"
	"os"
	"reflect"
)

type Variable struct {
	Name     string
	Value    any
	context  Context
	Constant bool
}

func (v *Variable) SetValue(value any) {
	field := reflect.TypeOf(v.Value)
	argType := reflect.TypeOf(value)
	if field != argType {
		PrintError(fmt.Errorf("variable type %s cannot convert to %s", field.Name(), argType.Name()), "erro to assign variable value", 0, v.context.script.Name)
		os.Exit(-1)
	}

	v.Value = value
}
