package ctci
import "testing"
import "container/list"

func TestPlusTwoLinkLists(t *testing.T) {
    l1 := list.New()
    l1.PushBack(3)
    l1.PushBack(1)
    l1.PushBack(5)

    l2 := list.New()
    l2.PushBack(5)
    l2.PushBack(9)
    l2.PushBack(4)

    l := PlusTwoLinkLists(l1, l2)
    if l.Len() != 4 {
        t.Error("TestPlusTwoLinkLists failed, Len:", l.Len)
    }

    iter := l.Front()
    expect_seri := []int{8, 0, 0, 1}
    for _, value := range expect_seri {
        if iter.Value.(int) != value {
            t.Error("TestPlusTwoLinkLists failed, value:", iter.Value.(int))
        }

        iter = iter.Next()
    }
}
