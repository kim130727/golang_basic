/*
**선언 형태 : `const 상수이름 상수형`**
- 한번 선언 및 초기화되면 수정할 수 없기 때문에 꼭 선언과 동시에 초기화를 해야함. 선언만 한다면 에러가 발생함.
- 초기화 후에 사용하지 않아도 에러가 발생하지 않음. 변수와 다르게 상수는 명시하는 것 자체에 의미가 있기 때문.
- 상수는 `var` 키워드 대신에 `const` 키워드를 사용하고 생략할 수 없기 때문에 자연스럽게 `:=` 용법을 사용할 수 없음.
*/

package main

import "fmt"

const username = "kim"

func main() {
	const a int = 1
	const b, d = 10, 20 //여러 값을 초기화할 수 있음
	const c = "goorm"

	fmt.Println(username)
}
