package definitions

// A SubStack is the equivalent of a code block
type SubStack struct {
    Variables map[string]*Variable
    Parent    *SubStack
}
