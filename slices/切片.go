package main

import "fmt"

func main() {
	// 与数组不同，slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）。
	// 要创建一个长度不为 0 的空 slice，需要使用内建函数 make。
	// 这里我们创建了一个长度为 3 的 string 类型的 slice（初始值为零值）
	s := make([]string, 3)
	fmt.Println("emp:", s) // emp: [  ] 初始值是零值

	// 我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)  // set: [a b c]
	fmt.Println("get:", s[2]) // get: c

	fmt.Println("len:", len(s))  // len: 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}