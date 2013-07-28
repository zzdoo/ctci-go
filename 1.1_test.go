package ctci
import "testing"

func TestIsUniqChar(t *testing.T) {
    if !IsUniqChar("world") {
        t.Error("There are some chars not unique.")
    }
}
