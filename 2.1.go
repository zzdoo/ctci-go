/*
原文：
Write code to remove duplicates from an unsorted linked list.
FOLLOW UP
How would you solve this problem if a temporary buffer is not allowed?
译文：
从一个未排序的链表中移除重复的项
进一步地，
如果不允许使用临时的缓存，你如何解决这个问题？
*/
package ctci
import "container/list"

func RemoveDuplicateItems(l *list.List) {
    pos := 0
    for e := l.Front(); e != nil; e = e.Next() {
        pos1 := 0
        for n := l.Front(); n != nil; n = n.Next() {
            if pos1 <= pos {
                pos1 += 1
                continue
            }
            if n.Value == e.Value {
                p := n.Prev()
                l.Remove(n)
                n = p
            }
        }
        pos += 1
    }
}
