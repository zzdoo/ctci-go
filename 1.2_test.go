package ctci
import "testing"
import "fmt"

func TestReverseCString(t *testing.T) {
    input := "hello\n"
    fmt.Println("Before: ", input)
    output := ReverseCString(input)
    fmt.Println("After: ", output)

    if output != "olleh\n" {
        t.Error("ReverseCString wrong!")
    }
}
