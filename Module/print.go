package main

import "fmt"

func main() {
	// 向右填充空格，不满15个位置，则使用空格填充
	fmt.Printf("%-15v $%4v\n", "向右填充空格", 15)
	// 向左填充
	fmt.Printf("%15v $%4v\n", "向左填充空格", 15)
}
