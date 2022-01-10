/*
在 Go 中，变量 需要显式声明，并且在函数调用等情况下， 编译器会检查其类型的正确性。
 */
package main

import "fmt"

func main() {
	// var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)  //initial

	// 你可以一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c) //1 2

	// Go 会自动推断已经有初始值的变量的类型
	var d = true
	fmt.Println(d) // true

	// 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。 例如，int 的零值是 0
	var e int
	fmt.Println(e) // 0

	// := 语法是声明并初始化变量的简写
	// 例如 var f string = "short" 可以简写为右边这样
	f := "short"
	fmt.Println(f) // short
}