/*
수식 연산자는 두 개의 피연산자를 요구하는 이항 연산자(binary operator)입니다.
다른 언어들과 마찬가지로 기본적인 사칙연산이 있고, 값을 나눈 나머지 값을 반환하는 연산자가 있습니다.
*/

package main

import "fmt"

func main() {
	num1, num2 := 17, 5
	str1, str2 := "Hello", "goorm!"

	fmt.Println("num1 + num2 =", num1+num2)
	fmt.Println("str1 + str2 =", str1+str2)
	fmt.Println("num1 - num2 =", num1-num2)
	fmt.Println("num1 * num2 =", num1*num2)
	fmt.Println("num1 / num2 =", num1/num2)
	fmt.Println("num1 % num2 =", num1%num2)
}
