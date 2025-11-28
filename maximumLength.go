package main

import (
	"fmt"
)

func findLength(nums1 []int, nums2 []int) int {
    n, m := len(nums1), len(nums2)
    dynamicP := make([][]int, n+1)
    for i := range dynamicP {
        dynamicP[i] = make([]int, m+1)
    }
    res := 0
    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            if nums1[i] == nums2[j] {
                dynamicP[i][j] = dynamicP[i+1][j+1] + 1
                if dynamicP[i][j] > res {
                    res = dynamicP[i][j]
                }
            } else {
                dynamicP[i][j] = 0
            }
        }
    }

    return res
}

func main() {
	// nums1 := []int{1, 2, 3, 6, 1}
	nums1 := []int{1, 2, 3, 2, 1}
    nums2 := []int{3, 2, 1, 4, 7}
    res := findLength(nums1, nums2)
    fmt.Println(res)
}
