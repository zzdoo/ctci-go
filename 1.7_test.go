package ctci
import "testing"
import "fmt"

func TestSetZero(t *testing.T) {
    mat := [][]int{{1,2,3,4}, {5,0,7,8}, {9,10,11,12}, {13,14,15,0}}
    matZero := [][]int{{1,0,3,0}, {0,0,0,0}, {9,0,11,0}, {0,0,0,0}}
    SetZero(mat)

    l := len(mat)
    err := false
    for j:=0; j<l; j++ {
        for i:=0; i<l; i++ {
            fmt.Printf("%v ", mat[j][i])
            if mat[j][i] != matZero[j][i] {
                err = true
            }
        }
        fmt.Printf("\n")
    }
    if err == true {
        t.Error()
    }
}
