package ctci
import "testing"

func TestReplaceWhiteSpace(t *testing.T) {
    input := "hello world "
    output := ReplaceWhiteSpace(input)
    if "hello%20world%20" != output {
        t.Error("Error result: " + output)
    }
}

