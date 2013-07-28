/*
原文：
Design an algorithm and write code to remove the duplicate characters in a string without using any additional buffer. NOTE: One or two additional variables are fine. An extra copy of the array is not.
FOLLOW UP
Write the test cases for this method.
译文：
设计算法并写出代码移除字符串中重复的字符，不能使用额外的缓存空间。注意： 可以使用额外的一个或两个变量，但不允许额外再开一个数组拷贝。
进一步地，为你的程序写测试用例。

问题简化为，如何去除数据里所有为0的元素？
有个思路陷阱，遍历应该以src作为主线。
这个版本还有改进，就是以int32作为变量置位来判断是否字符出现过
*/
package ctci

func RemoveRepeatChars(input string) string {
    data := []rune(input)

    for i:=0; i<len(data)-1; i++ {
        if data[i] == 0 {
            continue
        }

        for j:=i+1; j<len(data); j++ {
            if data[i] == data[j] {
                data[j] = 0
            }
        }
    }

    // 改进后的版本
    dst := 0
    for src := 0; src<len(data); src++ {
        if data[src] == 0 {
            continue
        }

        if dst < src {
            data[dst] = data[src]
        }
        dst++
    }

    /*
    dst := 0; src := 0
    for src < len(data) {
        if data[src] == 0 {
            src++
            continue
        }

        if dst < src {
            data[dst] = data[src]
        }
        dst++; src++
    }
    */

    return string(data[0:dst])
}
