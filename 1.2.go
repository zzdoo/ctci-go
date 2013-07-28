/*
原文：
Write code to reverse a C-Style String. (C-String means that “abcd” is represented as five characters, including the null character.)
译文：
写代码翻转一个C风格的字符串。(C风格的意思是"abcd"需要用5个字符来表示，包含末尾的结束字符)
*/
package ctci

func ReverseCString(input string) string {
    runeStr := []rune(input)
    l := len(runeStr) - 1 // Ignore \n
    for i:=l/2; i>=0; i-- {
        runeStr[i], runeStr[l-i-1] = runeStr[l-i-1], runeStr[i]
    }

    return string(runeStr)
}
