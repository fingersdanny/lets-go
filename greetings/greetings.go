package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// 아무 이름도 안들어온다면, 메시지와 에러를 반환한당~
	if name == "" {
		return "", errors.New("empty name")
	}

	// 이름이 들어온다면, 이름을 담고 있는 값을 환영 메시지에 담아서 반환한당.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat 함수는 환영 인사 메시지 세트 중 하나를 반환합니다.
// 반환하는 메시지는 랜덤하게 정해집니다.
func randomFormat() string {
	// 메시지 포맷 리스트의 슬라이스
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v!. Well met!",
	}

	// formats을 슬라이싱 하고 랜덤한 인덱스를 특정하여
	// 해당 랜덤 인덱스의 메시지 포맷을 반환합니다.
	return formats[rand.Intn(len(formats))]
}
