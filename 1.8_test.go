package ctci
import "testing"

func TestIsRotateString(t *testing.T) {
    str1 := "waterbottle"
    str2 := "erbottlewat"

    if !IsRotateString(str1, str2) ||
        !IsRotateString(str2, str1) {
        t.Error()
    }
}
