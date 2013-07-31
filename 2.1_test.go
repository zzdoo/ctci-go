package ctci
import "testing"
import "container/list"
import "fmt"

func TestRemoveDuplicatedItems(t *testing.T) {
    l := list.New()
    l.PushBack(4)
    l.PushBack(5)
    l.PushBack(4)
    l.PushBack(4)
    l.PushBack(5)
    l.PushBack(7)
    l.PushBack(4)
    l.PushBack(4)
    l.PushBack(4)

    printList(l)
    RemoveDuplicateItems(l)
    printList(l)
}

func printList(l *list.List) {
    for e:= l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
    fmt.Println("================")
}
