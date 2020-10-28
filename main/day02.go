package main

import "fmt"

func main() {
	str := "你好北京"
	// 字符串遍历方式: 传统for循环
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	fmt.Println("---------------")
	for i := range str {
		fmt.Printf("index = %d, val = %c\n", i, str[i])
	}
	fmt.Println("---------------")
	// 字符串遍历方式2: for-range
	for index, val := range str {
		fmt.Printf("index = %d, val = %c\n", index, val)
	}
}
