/*
증감 연산자는 값을 1만큼 증가시키거나 감소시키는 연산자입니다. Go언어에서 증감 연산자를 사용할 때
주의할 점이 두 가지가 있습니다. 첫 번째는 증감 연산자를 사용하고 동시에 대입할 수 없습니다.
두 번째는 전위 연산을 할 수 없습니다. 아래 예시가 있습니다.
*/

package main

import "fmt"

func main() {
	count1, count2 := 1, 10.4

	count1++
	count2--

	fmt.Println("count1++ :", count1)
	fmt.Println("count2-- :", count2)
}
