package src

import (
    "custom-lang/src/definitions"
    "custom-lang/src/execution"
    "custom-lang/src/instructions"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Check if a variable exist in a substack tree
func checkSubstackForVariableName(substack *definitions.SubStack, variableName string) bool {
    if substack == nil {
        return false
    }

    var found bool
    _, found = substack.Variables[variableName]

    return found || checkSubstackForVariableName(substack.Parent, variableName)
}

// Get a variable in a substack tree if it exists
func getVariableInSubstack(substack *definitions.SubStack, variableName string) (*definitions.Variable, error) {
    if substack == nil {
        return nil, fmt.Errorf("variable %s does not exist", variableName)
    }

    var found bool
    var variable *definitions.Variable
    variable, found = substack.Variables[variableName]

    if found {
        return variable, nil
    }

    return getVariableInSubstack(substack, variableName)
}

// Parse a line for a variable declaration
func declareVariable(line string, lineNumber int, substack *definitions.SubStack, types map[string]definitions.Typing) (*definitions.Variable, error) {
    var split []string = strings.Split(line, " ")
    if len(split) != 3 {
        // Should have "VariableDecl_keyword type name" structure
        return nil, fmt.Errorf("line %d is not a valid variable", lineNumber)
    }

    if checkSubstackForVariableName(substack, split[1]) {
        // This variable should not exist in parent substack
        return nil, fmt.Errorf("line %d variable name already exists", lineNumber)
    }
    var typing definitions.Typing
    var found bool
    typing, found = types[split[1]]
    if !found {
        return nil, fmt.Errorf("line %d variable does not have a valid type", lineNumber)
    }

    var variable *definitions.Variable = new(definitions.Variable)
    variable.Typing = typing
    variable.Name = split[2]
    return variable, nil
}

// Parse a line for a variable assignation
func assignVariable(variableName string, variableLiteral string, lineNumber int, substack *definitions.SubStack, types map[string]definitions.Typing) (definitions.Instruction, error) {
    var variable *definitions.Variable
    var err error
    variable, _ = getVariableInSubstack(substack, variableName)

    var assign definitions.AssignationInstruction = variable.Typing.AssignLiteralInstruction()
    if assign == nil {
        return nil, fmt.Errorf("the variable at line %d does not support literal assignation", lineNumber)
    }
    assign.SetLineNumber(lineNumber)
    assign.SetVariableName(variableName)
    err = assign.SetLiteralWithError(variableLiteral)
    if err != nil {
        return nil, err
    }
    return assign, nil
}

// Parse text to create a substack and a list of instruction
func LoadSubstack(lines []string, lineOffset int, parentSubstack *definitions.SubStack, types map[string]definitions.Typing) (*definitions.SubStack, []definitions.Instruction, error) {
    var substack *definitions.SubStack = new(definitions.SubStack)
    substack.Parent = parentSubstack
    substack.Variables = make(map[string]*definitions.Variable)
    var instructionList []definitions.Instruction = make([]definitions.Instruction, 0)

    // One instruction per line
    for lineNumber, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" {
            // No parsing to be done
            continue
        }

        if strings.HasPrefix(line, definitions.VariableDecl) {
            // Variable creation
            var variable *definitions.Variable
            var err error
            variable, err = declareVariable(line, lineNumber, substack, types)
            if err != nil {
                return nil, nil, err
            }
            substack.Variables[variable.Name] = variable
        } else {
            var l []string = strings.Split(line, " ")
            if len(l) >= 2 && checkSubstackForVariableName(substack, l[0]) {
                // Variable assignation
                var variableName string = l[0]
                var varValue string = l[1]
                var err error
                var assign definitions.Instruction

                assign, err = assignVariable(variableName, varValue, lineNumber, substack, types)
                if err != nil {
                    return nil, nil, err
                }

                instructionList = append(instructionList, assign)
            }
        }
    }

    return substack, instructionList, nil
}

// Parse a file
func LoadModule(filename string) (*execution.Program, error) {
    if !exist(filename) {
        return nil, fmt.Errorf("%s does not exist", filename)
    }

    var err error
    var content []byte
    content, err = os.ReadFile(filename)

    if err != nil {
        return nil, err
    }

    var contentStr = string(content)
    var lines = strings.Split(contentStr, "\n")

    var types map[string]definitions.Typing = make(map[string]definitions.Typing, 0)

    // Define default types
    types["int"] = definitions.Typing{
        TypeName: "int",
        ParseLiteralFunc: func(literal string, literalSettable definitions.LiteralSettable) error {
            var e error

            _, e = strconv.Atoi(literal)
            if e != nil {
                return errors.New("")
            }
            literalSettable.SetLiteral(literal)
            return nil
        },
        AssignLiteralInstruction: func() definitions.AssignationInstruction {
            return new(instructions.IntAssignation)
        },
    }
    types["bool"] = definitions.Typing{
        TypeName: "bool",
        ParseLiteralFunc: func(literal string, literalSettable definitions.LiteralSettable) error {
            if literal == "true" || literal == "false" {
                literalSettable.SetLiteral(literal)
                return nil
            }

            return errors.New("")
        },
        AssignLiteralInstruction: func() definitions.AssignationInstruction {
            return new(instructions.BoolAssignation)
        },
    }

    var substack *definitions.SubStack
    var instr []definitions.Instruction
    substack, instr, err = LoadSubstack(lines, 0, nil, types)
    if err != nil {
        return nil, err
    }

    var program *execution.Program = new(execution.Program)
    program.Stacks = make([]*execution.Stack, 0)
    program.Stacks = append(program.Stacks, &execution.Stack{
        Variables:    make(map[string]*definitions.Variable),
        Instructions: make([]definitions.Instruction, 0),
    })

    for varName, variable := range substack.Variables {
        program.Stacks[0].Variables[varName] = variable
        program.Stacks[0].Instructions = append(program.Stacks[0].Instructions, instr...)
    }

    return program, nil
}
