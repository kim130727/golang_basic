/*
관계 연산자는 두 값의 대소와 동등의 관계를 따지는 연산자
관계 연산자는 조건을 만족하면 `true`를, 만족하지 않으면 `false`를 반환
*/

package main

import "fmt"

func main() {
	fmt.Println("13 == 13 : ", 13 == 13)
	fmt.Println("13 == 23 : ", 13 == 23)
	fmt.Println("13 != 13 : ", 13 != 13)
	fmt.Println("3 != 5 : ", 3 != 5)
	fmt.Println("0 < 1 : ", 0 < 1)
	fmt.Println("0 > 1 : ", 0 > 1)
	fmt.Println("0 >= 1 : ", 0 >= 1)
	fmt.Println("0 <= 1 : ", 0 <= 1)
}
