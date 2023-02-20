/*
포인터 연산자는 C++언어와 같이 &와 *연산자를 이용해
메모리에 접근할 수 있도록 함.
주의할 점은 Go언어는 포인터 연산자를 제공하지만 포인터 산술 즉,
포인터에 더하고 빼는 기능은 제공하지 않음
*/

package main

import "fmt"

func main() {
	var num int = 5
	var pnum = &num

	fmt.Println("num : ", num)   //num 값
	fmt.Println("pnum :", pnum)  //num의 메모리 주소
	fmt.Println("pnum :", *pnum) //num의 주소로 메모리에 할당돼있는 값 접근

	*pnum++
	fmt.Println("num : ", num)
	fmt.Println("pnum :", *pnum)
	//포인터 연산자를 이용한 값 변경
}
