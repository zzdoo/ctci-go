package ctci
import "testing"

func TestCheckAnagrams(t *testing.T) {
    str1 := "abbcd"
    str2 := "abcdb"
    if !CheckAnagrams(str1, str2) ||
        !CheckAnagramsV2(str1, str2) {
        t.Error()
    }
}

