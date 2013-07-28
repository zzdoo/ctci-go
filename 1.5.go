/*
原文：
Write a method to replace all spaces in a string with ‘%20’.
译文：
写一个函数，把字符串中所有的空格替换为%20 。
*/
package ctci

// O(2n)
func ReplaceWhiteSpace(input string) string {
    r := []rune(input)
    num := 0
    for i:= 0; i<len(r); i++ {
        if r[i] == ' ' {
            num++
        }
    }
    o := make([]rune, len(r) + num*2)
    for i, j:=0, 0; i<len(r); i++ {
       if r[i] == ' ' {
            o[j] = '%'
            o[j+1] = '2'
            o[j+2] = '0'
            j += 3
       } else {
           o[j] = r[i]
           j++
       }
    }

    return string(o)
}
