/*
- Go언어에서는 꼭 `fmt` 패키지를 `import` 하지 않아도 기본적으로 콘솔 출력 함수인 `println`과 `print` 함수를 지원
  두 함수의 차이점은 단순히 호출 후 개행을 하느냐 안 하느냐 입니다.
- `println` 함수는 호출 후 개행을 하고 `print` 함수는 호출 후 개행을 하지 않습니다.
   따라서 개행을 하기 위해서 개행을 의미하는 이스케이프 시퀀스인 '`\n`'을 입력해야 합니다.
- 사용 형태는 `print(출력하고자 하는 데이터)`입니다. 여러 데이터를 출력할 때는 콤마(,)를 사용하면 됩니다.
   그리고 이 함수들은 함수 안에서의 연산 식을 결과 값으로 출력이 가능합니다.
- 예를 들어 `fmt.Println(3+5)`를 하면 `8`이 출력됩니다.
- Go언어에서 콘솔 입출력을 위해서는 fmt 패키지를 `import`해서 사용합니다.
  물론, Go언어에서 기본으로 제공하는 입출력 함수도 있지만, 그것을 사용하는 것보다는 조금 더 강력한 입출력 기능을 제공하는
fmt 패키지의 사용을 권장합니다. 앞으로 이 강좌의 모든 예제는 fmt 패키지를 이용하여 콘솔 입출력을 진행할 것입니다.
- `package main` 밑줄에 `import "fmt"`를 작성해둠으로써 이 패키지를 사용하겠다고 선언합니다.
- 사용 형태는 `fmt.Print(출력하고자 하는 데이터)`입니다. 콘솔 출력을 위해 가장 많이 쓰이는 함수인
   `Println`과 `Print` 함수의 사용법은 아래의 코드와 같습니다.
- 그리고 `Printf` 함수는 서식 문자를 활용하여 원하는 포맷으로 데이터를 채워서 출력하고자할 때 사용합니다.
  (C언어에도 동일한 이름의 함수가 있습니다.)
*/

package main

func main() {
	var num1 int = 1
	var num2 int = 2

	print("Hello goorm!")
	print(num1)
	print(num2)
	print(num1 + num2)
	print("Hello goorm!", num1+num2, "\n")

	println("Hello goorm!")
	println(num1)
	println(num2)
	println(num1 + num2)
	println("Hello goorm!", num1+num2)
}
