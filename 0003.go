/*함수 연습1*/

package main

import "fmt"

func input(a, b int) (int, int) {
	return a * b
}

func main() {
	num := input()
  fmt.Println(num)
}
