/*
原文：
Write a method to decide if two strings are anagrams or not.
译文：
写一个函数判断两个字符串是否是变位词。
例如abbcd和abcdb

*/
package ctci

// O(n)
func CheckAnagramsV2(str1 string, str2 string) bool {
    r1 := []rune(str1)
    r2 := []rune(str2)
    if len(r1) != len(r2) {
        return false
    }

    b := make([]int, 256)
    for _,v := range(r1) {
        b[v] += 1
    }
    for _,v := range(r2) {
        b[v] -= 1
    }

    for _,v := range(b) {
        if v != 0 {
            return false
        }
    }

    return true
}

// O(n2), my version
func CheckAnagrams(str1 string, str2 string) bool {
    r1 := []rune(str1)
    r2 := []rune(str2)

    if len(r1) != len(r2) {
        return false
    }

    l := len(r1)
    for i:=0; i<l; i++ {
        for j:=0; j<l; j++ {
            if r1[i] == r2[j] {
                r2[j] = 0
                continue
            }
        }
    }

    for j:=0; j<l; j++ {
        if r2[j] != 0 {
            return false
        }
    }

    return true
}
