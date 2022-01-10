package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 错误检查太多，直接拎出来写个方法
func check(e error) {
	if e != nil{
		panic(e)
	}
}

func main(){
	// 1. (基本) 文件读取任务或许就是将文件内容读取到内存中。
	dat, err := ioutil.ReadFile("./read/string.txt")
	check(err)
	fmt.Println(string(dat))

	// =====================================================================================================================
	// 2. (高级) 可以对文件的读取方式和内容进行更多控制

	// 2.1 先用open打开文件，获取一个os.File对象
	f, err := os.Open("./read/string.txt")
	check(err)

	// 2.2 f.Read方法需要传入一个byte切片-带长度，用来存   (切片不是可以无限扩容的吗？为什么还要长度？ 答切片是可变长度，用var就不用指定长度。只是make可以预先分配初始的而已)
	// 发现只会取5长度
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)  // 读取后，读到的内容存b1里，读到的真正长度存n1里(因为有可能字符不够5个字)。
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))



}
