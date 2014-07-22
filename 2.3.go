/*
实现一个算法来删除单链表中间的一个结点，只给出指向那个结点的指针。
输入：指向链表a->b->c->d->e中结点c的指针
结果：不需要返回什么，得到一个新链表：a->b->d->e
*/

package ctci
import "container/list"

// 用不重复的int代替指针
func DeleteNodeInList(l *list.List, e int) {
    iter := l.Front()
    for ; iter.Value.(int) != e; iter = iter.Next() {
    }

    for ; iter.Next() != nil; {
        iter.Value = iter.Next().Value
        iter = iter.Next()
    }

    l.Remove(iter)
}
