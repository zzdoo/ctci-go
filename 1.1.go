/*
原文：
Implement an algorithm to determine if a string has all unique characters. What if you can not use additional data structures?
译文：
实现一个算法来判断一个字符串中的字符是否唯一(即没有重复).不能使用额外的数据结构。 (即只使用基本的数据结构)
*/
package ctci

// 需要确定字符串的范围
func IsUniqChar(str string) bool {
    bArr := make([]bool, 256)
    for _,v := range(str) {
        if bArr[v] {
            return false
        }
        bArr[v] = true
    }
    return true
}
