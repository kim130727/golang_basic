/*
* Go에서의 변수 선언 방식은 `var 변수이름 변수형`
- 변수를 선언한 곳에서 바로 초기값을 설정할 수 있음
- 주의해야할 점은 이 용법은 함수(func) 내에서만 사용이 가능합니다.
  따라서 함수 밖에서(전역 변수)는 꼭 `var`키워드를 선언해줘야함
- a는 선언과 동시에 1로 초기화. c, d는 오류없이 int와 string이라는 자료형으로 자동 지정
- Go에서는 변수를 선언하고 초기값을 설정하지 않으면 'Zero value'로 설정
  : bool 타입 false, 숫자 타입 0, string 타입 ""(빈 문자열)
* 동일한 형의 변수를 한 번에 여러개 선언
- 변수의 개수와 초기화하는 값의 개수가 동일해야 함.
  초기화하지 않는다면 모든 값을 초기화 하지 않아야 함.
*/

package main

import "fmt"

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)

	i, j, k := 1, 2, 3
	fmt.Println(i, j, k)

	var str1, str2 string = "Hello", "goorm"
	fmt.Println(str1, str2)
}
