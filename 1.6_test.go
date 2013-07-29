package ctci
import "testing"
import "fmt"

func TestRotateMatrix(t *testing.T) {
    mat := [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}, {13,14,15,16}}
    matRotate := [][]int{{4,8,12,16}, {3,7,11,15},{2,6,10,14},{1,5,9,13}}
    RotateMatrix(mat)

    l := len(mat)
    err := false
    for j:=0; j<l; j++ {
        for i:=0; i<l; i++ {
            fmt.Printf("%v ", mat[j][i])
            if mat[j][i] != matRotate[j][i] {
                err = true
            }
        }
        fmt.Printf("\n")
    }
    if err == true {
        t.Error()
    }
}
