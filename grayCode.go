package main

import (
	"errors"
	"fmt"
)

func grayCode(n int) ([]int, error) {
	if n < 1 || n > 16 {
		return nil, errors.New("1 <= n <= 16")
	}

	if n == 1 {
		return []int{0, 1}, nil
	}

	prev, _ := grayCode(n - 1);
	size := len(prev);
	out := make([]int, 0, size*2)

	out = append(out, prev...)

	p := 1 << (n - 1)
	for i := size - 1; i >= 0; i-- {
		out = append(out, prev[i]|p);
	}

	return out, nil
}

func main() {
	n := 6
	res, err := grayCode(n)
	if err != nil {
		fmt.Println("Error:", err);
		return
	}

	fmt.Println(res)
}
