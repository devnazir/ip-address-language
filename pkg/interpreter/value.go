package interpreter

type Interpreter struct {
	Environment *Environment
}

type Environment struct {
	Variables map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{Variables: make(map[string]interface{})}
}

func (e *Environment) SetVariable(name string, value interface{}) {
	e.Variables[name] = value
}

func (e *Environment) GetVariable(name string) interface{} {
	return e.Variables[name]
}

func (e *Environment) HasVariable(name string) bool {
	_, ok := e.Variables[name]
	return ok
}
