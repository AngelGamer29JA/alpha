package alpha

import "fmt"

type Context struct {
	script    *Script
	variables map[string]*Variable
}

func NewContext(script *Script) *Context {
	return &Context{
		script:    script,
		variables: make(map[string]*Variable),
	}
}

func (c *Context) RegisterVariable(name string, value any, constant bool) {
	c.variables[name] = &Variable{
		Name:     name,
		Value:    value,
		context:  *c,
		Constant: constant,
	}
}

func (c *Context) GetVariable(name string) *Variable {
	v, exist := c.variables[name]
	if !exist {
		panic(fmt.Sprintf("variable %s not found", v.Name))
	}
	return v
}
