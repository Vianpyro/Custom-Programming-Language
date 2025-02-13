package tests

import (
    "testing"
)
import "custom-lang/src"

func TestModuleImport(t *testing.T) {
    var err error

    _, err = src.LoadModule("./tests/data/not-a-file")
    if err == nil {
        t.Error("Expected error, got none")
    }

    _, err = src.LoadModule("./tests/data/test-module-1")
    if err != nil {
        t.Errorf("Expected no error, got an error %s", err)
    }
}
