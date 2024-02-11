package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// 아무 이름도 안들어온다면, 메시지와 에러를 반환한당~
	if name == "" {
		return "", errors.New("empty name")
	}

	// 이름이 들어온다면, 이름을 담고 있는 값을 환영 메시지에 담아서 반환한당.
	message := fmt.Sprintf("Hi, %v. Welcom!", name)
	return message, nil
}
