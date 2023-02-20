/*
1 매개변수는 없고 반환값도 없는 함수 : `func guide()`
- 두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.
2 매개변수는 없고 반환값 두개 있는 함수 : `func input()`
- 두 정수를 입력
3 매개변수 있고 반환값도 있는 함수 : `func multi()`
- 두 정수를 곱하기
4 매개변수 있고 반환값 없는 함수 : `func printResult()`
- 결과 도출

매개변수 (parameter) : 서로 종속인 변수들을 묶어주는 변수
- 컴퓨터 프로그래밍 : 보통은 함수에 투입되는 변수
- 매개변수는 변수이고 전달인자는 값이다. 때문에 주소를 함수에 매개변수로서
  넣으면 전달 인자는 그 주소가 가리키는 값이 된다. 엄밀히 말하면
  매개변수(패러미터)는 함수 정의 시 초기 값이 프로그래머에 의해 정해지는 것이며,
  반면 런타임 동안에 사용자로부터 나중에 값을 입력 받는 변수는
  아규먼트(argument)라고 한다.
*/

/*함수*/

package main

import "fmt"

/*기능들만 모아놓은 함수들*/
func guide() { //매개변수 X 반환 값 X
	fmt.Println("두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.\n두 정수를 차례로 띄어 써주세요.")
	fmt.Print("두 정수를 입력해주세요 :")
}

func input() (int, int) { //매개변수 X 반환 값 O(두 개)
	var a, b int
	fmt.Scanln(&a, &b)
	return a, b
}

func multi(a, b int) int { //매개변수 O, 반환 값 O
	return a * b
}

func printResult(num int) { //매개변수 O, 반환 값 X
	fmt.Printf("결과값은 %d입니다. 프로그램을 종료합니다.\n", num)
}

func main() { //간결해진 main문
	guide()
	num1, num2 := input()
	result := multi(num1, num2)
	printResult(result)
}
