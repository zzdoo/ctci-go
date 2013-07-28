package ctci
import "testing"
import "fmt"

func TestRemoveRepeatChars(t *testing.T) {
    input := "hellllllllooob"
    fmt.Println("Before: ", input)
    output := RemoveRepeatChars(input)
    fmt.Println("After: ", output)

    if output != "helob" {
        t.Error()
    }
}
