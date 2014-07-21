package ctci
import "testing"
import "container/list"

/*
    Last 3
    0 1 2 3 4 5 null
    | | | * | |
*/
func TestFindLastK(t *testing.T) {
    l := list.New()
    l.PushBack(0)
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    l.PushBack(4)
    l.PushBack(5)

    n, err := FindLastKNodes(l, 3)
    if n != 3 || err != nil {
        t.Error("TestFindLastK failed.")
    }
}
