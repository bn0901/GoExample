/*
Go 拥有多种值类型，包括字符串、整型、浮点型、布尔型等。 下面是一些基础的例子。
 */
package main

import "fmt"

func main() {

	// 字符串可以通过 + 连接
	fmt.Println("go" + "lang")  // golang

	// 整数和浮点数
	fmt.Println("1+1 =", 1+1)  // 1+1 = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0)  // 7.0/3.0 = 2.3333333333333335  >>只要除数或被除数有一个是浮点数，结果为浮点数
	fmt.Println("7.0/3.0 =", 7/3)  // 2 >>整数相除，结果为整数)

	// 布尔型，以及常见的布尔操作
	fmt.Println(true && false)  // false
	fmt.Println(true || false)  // true  >>只要一个为真则为真
	fmt.Println(!true)   // false
}