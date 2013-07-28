/*
原文：
Design an algorithm and write code to remove the duplicate characters in a string without using any additional buffer. NOTE: One or two additional variables are fine. An extra copy of the array is not.
FOLLOW UP
Write the test cases for this method.
译文：
设计算法并写出代码移除字符串中重复的字符，不能使用额外的缓存空间。注意： 可以使用额外的一个或两个变量，但不允许额外再开一个数组拷贝。
进一步地，为你的程序写测试用例。
*/
package ctci

func RemoveRepeatChars(input string) string {
    last := rune(0)
    output := ""
    for _,v := range(input) {
        if v == last {
            continue
        }
        output = output + string(v)
        last = v
    }

    return output
}
