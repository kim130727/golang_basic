/*
비트 단위의 연산을 진행하는 연산자.
비트 연산자는 사실 기계에 좀 더 친화적인 연산자지만
다른 영역에도 사용돼 효율성을 높이고 연산의 수를 줄이는 요인이 되기도 합니다.
활용적 측면을 이해하기엔 큰 부담이 따르기 때문에 연산자의 기능을 이해하는 데
초점을 맞추어 정리합니다.
*/

package main

import "fmt"

func main() {
	num1 := 15 //00001111
	num2 := 20 //00010100

	fmt.Printf("num1 & num2 : %08b, %d\n", num1&num2, num1&num2)
	fmt.Printf("num1 | num2 : %08b, %d\n", num1|num2, num1|num2)
	fmt.Printf("num1 ^ num2 : %08b, %d\n", num1^num2, num1^num2)
	fmt.Printf("num1 &^ num2 : %08b, %d\n", num1&^num2, num1&^num2)
	fmt.Printf("num1 << 4 : %08b, %d\n", num1<<4, num1<<4)
	fmt.Printf("num2 >> 2 : %08b, %d\n", num2>>2, num2>>2)
}
