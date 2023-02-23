/*
입력 받은 두 정수를 나누었을 때 얻게 되는 몫과
나머지를 출력하는 프로그램을 작성해봅니다.
예를 들어 7과 2가 입력되면 몫으로 3, 나머지로 1이 출력돼야 합니다.
*/

package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	fmt.Printf("몫 : %d, 나머지 : %d\n", num1/num2, num1%num2)
}
