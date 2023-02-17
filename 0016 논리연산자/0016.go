/* 논리 연산자
논리 연산자는 익숙하게 알고있는 다른 언어들과 똑같습니다. AND(논리곱), OR(논리합), NOT(논리부정)을 연산합니다.
예시는 변수 A와 B로 들겠습니다. 주의할 점은 Go언어에서 논리부정 연산시 bool 형의 선언 및 사용만이 가능합니다.
즉, false와 true값만 사용할 수 있습니다. 아래 예시가 있습니다.*/

package main

import "fmt"

func main() {
	var a bool = true
	b := false

	fmt.Println("0 && 0 : ", b && b)
	fmt.Println("0 && 1 : ", b && a)
	fmt.Println("1 && 1 : ", a && a)
	fmt.Println("0 || 0 : ", b || b)
	fmt.Println("0 || 1 : ", b || a)
	fmt.Println("1 || 1 : ", a || a)
	fmt.Println("!1 ", !true)
	fmt.Println("!0 ", !false)
}
