/*
实现一个算法，找出单向链表中倒数第k个结点。
*/

package ctci
import "container/list"
import "errors"

func FindLastKNodes(l *list.List, k int) (int, error) {
    e := l.Front()
    for i:=0; (e!= nil && i < k); e,i = e.Next(),i+1 {
    }

    if e == nil {
        return -1, errors.New("List is shorter than k.")
    }

    ee := l.Front()
    for ; e != nil; ee,e = ee.Next(),e.Next() {
    }

    return ee.Value.(int), nil
}
