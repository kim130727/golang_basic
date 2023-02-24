/*
int형으로 표현되어 있는 데이터 표현방식을 float32형으로 바꾸거나, int형으로 표현되어있는 데이터의 표현방식을
uint형으로 바꾸는 것이 바로 '자료형의 변환'입니다.
즉, 자료형의 변환이라는 것은 데이터의 표현방식을 바꾸는 것입니다.
이러한 자료형의 변환은 다른 언어들을 비추어 보았을 때 크게 두 종류가 있습니다.
- 자동 형 변환(묵시적 형 변환)
- 강제 형 변환(명시적 형 변환)
이름에서 볼 수 있듯이 '자동 형 변환'은 자동으로 발생하는 형 변환이고,
'강제 형 변환'은 프로그래머가 형 변환을 명시해서 강제로 변환이 일어나게 하는 것입니다.
그런데 Go언어에서는 형 변환을 할 때 변환을 명시적으로 지정해주어야합니다.
예를 들어 float32에서 uint로 변환할 때, 암묵적 변환은 일어나지 않으므로
"uint(변수이름)"과 같이 반드시 변환을 지정해줘야합니다. 만약 명시적인 지정이 없다면 런타임 에러가 발생합니다.
*/

package main

import "fmt"

func main() {
	var num int = 10
	var changef float32 = float32(num) //int형을 float32형으로 변환
	changei := int8(num)               //int형을 int8형으로 변환

	var str string = "goorm"
	changestr := []byte(str)  //바이트 배열
	str2 := string(changestr) //바이트 배열을 다시 문자열로 변환

	fmt.Println(num)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)
}
