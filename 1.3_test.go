package ctci
import "testing"
import "fmt"

func TestRemoveRepeatChars(t *testing.T) {
    input := "hello"
    fmt.Println("Before: ", input)
    output := RemoveRepeatChars(input)
    fmt.Println("After: ", output)

    if output != "helo" {
        t.Error()
    }
}
