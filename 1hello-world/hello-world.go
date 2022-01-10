/*
我们的第一个程序将打印传说中的“hello world”， 右边是完整的程序代码。
 */
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}



// 运行方法：
// go run 1hello-world/hello-world.go

// 编译方法：(生成到当前目录下，名为hello-world)
// go build 1hello-world/hello-world.go
// 编译后可直接运行此二进制
// ./hello-world