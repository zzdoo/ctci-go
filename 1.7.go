/*
原文：
Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column is set to 0.
译文：
写一个函数处理一个MxN的矩阵，如果矩阵中某个元素为0，那么把它所在的行和列都置为0.
*/
package ctci

// 简化成一维矩阵，如何做到？
func SetZero(mat [][]int) {
    l := len(mat[0])
    i0 := []int{}
    j0 := []int{}
    for j:=0; j<l; j++ {
        for i:=0; i<l; i++ {
            if mat[j][i] == 0 {
                i0 = append(i0, i)
                j0 = append(j0, j)
            }
        }
    }

    for _,i := range(i0) {
        for j:=0; j<l; j++ { mat[j][i] = 0 }
    }
    for _,j := range(j0) {
        for i:=0; i<l; i++ { mat[j][i] = 0 }
    }
}
