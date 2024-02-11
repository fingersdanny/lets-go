package main

import (
	"awesomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	// 정해진 로거의 설정을 변경한다 ->
	// 로그 메시지의 접두사를 정하고
	// 플래그를 설정해서 시간, 소스 코드, 라인 넘버를 출력을 막는다.
	// 플래그에 따라서 로그에 어떤 종류의 로깅을 남길 지 정할 수 있다.
	//
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// 환영 메시지를 요청한다.
	message, err := greetings.Hello("")
	// 예외가 발생하면, 콘솔에 출력하고 프로그램을 종료한다.
	if err != nil {
		log.Fatal(err)
	}

	// 아무 예외도 반환되지 않았다면, 반환된 메시지를 콘솔에 출력한다.
	fmt.Println(message)
}
