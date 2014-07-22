package ctci
import "testing"
import "container/list"

func TestDeleteNodeInList(t *testing.T) {
    l := list.New()
    l.PushBack(0)
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    l.PushBack(4)
    l.PushBack(5)

    DeleteNodeInList(l, 3)

    if l.Len() != 5 || l.Back().Value.(int) != 5 {
        t.Error("TestDeleteNodeInList failed, Len:", l.Len(), " Back:", l.Back().Value.(int)) 
    }
}
