
package cmd

import (
    "testing"
)


func TestVersion(t *testing.T) {
    if Version() != 1.1 {
        t.Error("Expected Version To Be 1.1")
    }
}