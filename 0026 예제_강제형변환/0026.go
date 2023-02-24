/*
실습 내용
- int형 변수 num1, num2, num3을 선언합니다.
- 사용자로부터 입력받을 때 두 번째 수는 무조건 '음수'여야합니다.
- num1은 float32 형으로 num2는 uint 형으로 num3은 int64 형으로 변환 후 새로운 변수에 초기화합니다.
- Printf 함수를 사용해 형 변환 된 변수 세 개가 출력됩니다.
- 타입을 출력하는 서식문자는 '%T', 실수형을 출력하는 서식문자는 '%f' 를 입력합니다.
*/

package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Scanln(&num1, &num2, &num3)

	var change1 float32 = float32(num1)
	var change2 uint = uint(num2)
	var change3 int64 = int64(num3)

	fmt.Printf("%T, %f\n", change1, change1)
	fmt.Printf("%T, %v\n", change2, change2)
	fmt.Printf("%T, %d\n", change3, change3)
}
