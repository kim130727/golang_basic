/*
- 첫 번째는 Back Quote(`문자열`)을 이용한 방법
모양이 인용부호(' ')와 비슷해서 혼동할 수 있지만 다른 기호입니다. Back Quote로 둘러 싸인 문자열은 Raw String Literal이라고 부릅니다.
쉽게 말하자면, 이 안에 있는 문자열은 어느 기호든 문자열 자체로 인식되는 Raw String 값이라는 것입니다.
예를 들어, 개발자라면 익숙하게 알고 있는 이스케이프 시퀀스가 특별한 의미로 인식되지 않는다는 것입니다.
개행의 의미를 가지고 있는 \n이 Back Quote(역따옴표)에 있으면 개행 기능이 되지 않고 문자열 자체로 출력됩니다.

- 두 번째는 이중인용부호("문자열")를 이용한 방법
이중인용부호로 둘러싸인 문자열은 Interpreterd String literal이라고 부릅니다.
쉽게 말해, 안에 있는 문자열에 이스케이프 시퀀스같은 문자열들은 특별한 의미로 해석돼 그 기능을 수행한다는 것입니다.
그리고 이중인용부호를 사용할 시에 복수라인에 걸쳐 쓸 수 없습니다. 코딩을 할때 아무리 엔터를 치고 길게 쳐도 한 줄에 표현된다는 뜻입니다.
하지만 `\n`과 같은 기호가 있으면 개행의 의미로 해석돼 여러 라인에 걸쳐 쓸 수 있습니다.
그리고 '+' 연산자는 숫자 뿐만이 아니라 문자열도 합할 수 있다고 했습니다. 따라서 두 방법 모두 + 연산자를 이용해 결합해 표현할 수 있습니다.
*/

package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인. 역따옴표
	var rawLiteral string = `바로 실행해보면서 배우는 \n Golang`

	// Interpreted String Literal
	var interLiteral string = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름 " + "EDU\n" + "Golang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)
}
