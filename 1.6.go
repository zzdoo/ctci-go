/*
原文
Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
译文：
一张图像表示成NxN的矩阵，图像中每个像素是4个字节，写一个函数把图像旋转90度。 你能原地进行操作吗？(即不开辟额外的存储空间)
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16
===>
4 8 12 16
3 7 11 15
2 6 10 14
1 5 9 13
======分为两步====
1 5 9 13
2 6 10 14
3 7 11 15
4 8 12 16

4 8 12 16
3 7 11 15
2 6 10 14
1 5 9 13
*/
package ctci

// 假设每个像素是int
func RotateMatrix(mat [][]int) {
    l := len(mat[0])
    // First step: 
    for j:=0; j<l; j++ {
        for i:=0; i<j; i++ {
            mat[j][i], mat[i][j] = mat[i][j], mat[j][i]
        }
    }

    // Second step:
    for j:=0; j<l/2; j++ {
        for i:=0; i<l; i++ {
            mat[j][i], mat[l-j-1][i] = mat[l-j-1][i], mat[j][i]
        }
    }
}
