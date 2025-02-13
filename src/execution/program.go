package execution

import (
    "custom-lang/src/definitions"
)

type Stack struct {
    Variables    map[string]*definitions.Variable
    Instructions []definitions.Instruction
}

type Program struct {
    Stacks []*Stack
}
