/*
原文：
Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring ( i.e., “waterbottle” is a rotation of “erbottlewat”).
译文：
假设你有一个isSubstring函数，可以检测一个字符串是否是另一个字符串的子串。 给出字符串s1和s2，只使用一次isSubstring就能判断s2是否是s1的旋转字符串， 请写出代码。旋转字符串："waterbottle"是"erbottlewat"的旋转字符串。
waterbottlewaterbottle
*/
package ctci
import "strings"

// 问题转化为是否包含互换的字符串。分界线是对称的
// 思路错了！
// 方法是复制一个str1
func IsRotateString(str1 string, str2 string) bool {
    r1 := []rune(str1)
    r2 := []rune(str2)

    if len(r1) != len(r2) {
        return false
    }
    str11 := str1 + str1
    return strings.Contains(str11, str2)
}
