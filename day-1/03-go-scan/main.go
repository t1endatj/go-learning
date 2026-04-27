package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fullName string
	var age int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nhập họ tên: ")
	scanner.Scan()
	fullName = scanner.Text()

	fmt.Print("Nhập tuổi: ")
	fmt.Scan(&age)

	fmt.Println("Thông tin người dùng:")
	fmt.Println("Họ tên:", fullName)
	fmt.Println("Tuổi:", age)
}