package tests

import (
    "custom-lang/src"
    "custom-lang/src/execution"
    "testing"
)

func TestVariableAssignation(t *testing.T) {
    var err error
    var program *execution.Program

    program, err = src.LoadModule("./tests/data/test-variable-1")
    if err != nil {
        t.Errorf("Expected no error, got an error %s", err)
    }

    for _, instruction := range program.Stacks[0].Instructions {
        err = instruction.Execute(program.Stacks[0].Variables)
        if err != nil {
            t.Errorf("Expected no error, got an error %s", err)
        }
    }

    if program.Stacks[0].Variables["patate"].ValueInt != 10 {
        t.Errorf("Expected variable \"patate\" to be 10, got %d", program.Stacks[0].Variables["patate"].ValueInt)
    }

    if program.Stacks[0].Variables["test"].ValueBool != true {
        t.Errorf("Expected variable \"test\" to be true, got %v", program.Stacks[0].Variables["test"].ValueBool)
    }
}

func TestVariableIntBadLiteral(t *testing.T) {
    var err error
    _, err = src.LoadModule("./tests/data/test-variable-2")
    if err == nil {
        t.Errorf("Expected an error, got nothing")
    }
}

func TestVariableAssignationWithoutDeclaration(t *testing.T) {
    var err error
    _, err = src.LoadModule("./tests/data/test-variable-2")
    if err == nil {
        t.Errorf("Expected an error, got nothing")
    }
}
