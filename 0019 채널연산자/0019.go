/*
Go언어는 채널이라는 개념이 있습니다.
채널이랑 고루틴(goroutine)끼리 데이터를 주고 받고 실행 흐름을 제어하는 기능.
채널에 관한 상세한 내용은 채널 챕터에서 다루겠습니다.
채널 연산자는 채널을 사용할때 쓰는 연산자입니다.
*/

package main

import "fmt"

func main() {
	ch := make(chan int) //정수형 채널 생성

	go func() {
		ch <- 10
	}() //채널에 10을 보냄

	result := <-ch //채널로부터 10을 전달받음
	fmt.Println(result)
}
