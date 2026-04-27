package main

import "fmt"

const (
	appName  = "Go Const Demo"
	version  = "1.0.0"
	maxScore = 10
)

func main() {
	var studentName string
	var score float64

	fmt.Print("Nhập tên học sinh: ")
	fmt.Scan(&studentName)

	fmt.Print("Nhập điểm: ")
	fmt.Scan(&score)

	fmt.Println("App:", appName)
	fmt.Println("Version:", version)
	fmt.Println("Tên học sinh:", studentName)
	fmt.Println("Điểm:", score)
	fmt.Println("Điểm tối đa:", maxScore)
}