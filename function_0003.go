package main

import "fmt"

func input() int {
	var a, b int
	fmt.Scanln(&a, &b)
	return a * b
}

func main() {
	fmt.Println("두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.\n두 정수를 차례로 띄어 써주세요.")
	fmt.Print("두 정수를 입력해주세요 :")
	num := input()
	fmt.Printf("결과값은 %d입니다. 프로그램을 종료합니다.\n", num)
}
