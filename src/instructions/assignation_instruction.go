package instructions

import (
    "custom-lang/src/definitions"
    "fmt"
    "strconv"
)

type IntAssignation struct {
    variable string
    value    int
    line     int
}

func (a *IntAssignation) Execute(stackVariable map[string]*definitions.Variable) error {
    var variable *definitions.Variable = stackVariable[a.variable]
    if variable.Typing.TypeName != "int" {
        return fmt.Errorf("assignation: Unsupported assignation type %s for variable %s on line %d", variable.Typing, a.variable, a.line)
    }
    variable.ValueInt = a.value
    return nil
}

func (a *IntAssignation) LineNumber() int {
    return a.line
}

func (a *IntAssignation) SetLineNumber(line int) {
    a.line = line
}

func (a *IntAssignation) SetVariableName(name string) {
    a.variable = name
}

func (a *IntAssignation) SetLiteral(literal string) {
    a.value, _ = strconv.Atoi(literal)
}

func (a *IntAssignation) SetLiteralWithError(literal string) error {
    var err error

    a.value, err = strconv.Atoi(literal)
    if err != nil {
        return fmt.Errorf("assignation is not an int at line %d", a.line)
    }
    return nil
}

type BoolAssignation struct {
    variable string
    value    bool
    line     int
}

func (a *BoolAssignation) Execute(stackVariable map[string]*definitions.Variable) error {
    var variable *definitions.Variable = stackVariable[a.variable]
    if variable.Typing.TypeName != "bool" {
        return fmt.Errorf("assignation: Unsupported assignation type %s for variable %s on line %d", variable.Typing, a.variable, a.line)
    }
    variable.ValueBool = a.value
    return nil
}

func (a *BoolAssignation) LineNumber() int {
    return a.line
}

func (a *BoolAssignation) SetLineNumber(line int) {
    a.line = line
}

func (a *BoolAssignation) SetVariableName(name string) {
    a.variable = name
}

func (a *BoolAssignation) SetLiteral(literal string) {
    a.value = literal == "true"
}

func (a *BoolAssignation) SetLiteralWithError(literal string) error {
    if literal == "true" {
        a.value = true
    } else if literal == "false" {
        a.value = false
    } else {
        return fmt.Errorf("assignation is not a bool at line %d", a.line)
    }
    return nil
}
