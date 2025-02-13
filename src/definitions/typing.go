package definitions

import "fmt"

type LiteralSettable interface {
    SetLiteral(string)
    SetLiteralWithError(string) error
}

func DefaultParseLiteralFunc(literal string, l LiteralSettable) error {
    return fmt.Errorf("parse literals are not yet supported for this type")
}

type Typing struct {
    TypeName                 string
    ParseLiteralFunc         func(string, LiteralSettable) error
    AssignLiteralInstruction func() AssignationInstruction
}
