/*
你有两个由单链表表示的数。每个结点代表其中的一位数字。数字的存储是逆序的， 也就是说个位位于链表的表头。写一函数使这两个数相加并返回结果，结果也由链表表示。
例子：(3 -> 1 -> 5), (5 -> 9 -> 2)
输入：8 -> 0 -> 8
*/

package ctci
import "container/list"

func PlusTwoLinkLists(l1 *list.List, l2 *list.List) *list.List {
    iter1 := l1.Front()
    iter2 := l2.Front()
    l := list.New()

    carry := 0
    for ; iter1 != nil || iter2 != nil; {
        plus := iter1.Value.(int) + iter2.Value.(int)    
        total := plus + carry
        l.PushBack(total % 10)
        carry = total / 10

        iter1 = iter1.Next()
        iter2 = iter2.Next()
    }

    for ; iter1 != nil; {
        total := iter1.Value.(int) + carry
        l.PushBack(total % 10)
        carry = total / 10

        iter1 = iter1.Next()
    }

    for ; iter2 != nil; {
        total := iter2.Value.(int) + carry
        l.PushBack(total % 10)
        carry = total / 10

        iter2 = iter2.Next()
    }

    if carry > 0 {
        l.PushBack(carry)
    }

    return l
}
