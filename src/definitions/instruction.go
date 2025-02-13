package definitions

type Instruction interface {
    Execute(stackVariable map[string]*Variable) error
    LineNumber() int
    SetLineNumber(line int)
}

type AssignationInstruction interface {
    Instruction
    LiteralSettable
    SetVariableName(name string)
}
