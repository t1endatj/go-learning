package main

import "fmt"

func main() {
	var name string = "Mark"
	age := 18
	score := 8.5
	isStudent := true

	fmt.Println("Tên:", name)
	fmt.Println("Tuổi:", age)
	fmt.Println("Điểm:", score)
	fmt.Println("Là học sinh:", isStudent)

	age = 19
	score = 9.0

	fmt.Println("Tuổi mới:", age)
	fmt.Println("Điểm mới:", score)
}