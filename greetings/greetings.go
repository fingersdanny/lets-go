package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello 메서드는 입력으로 받는 사람에 대한 환영 메시지를 반환합니다.
func Hello(name string) (string, error) {
	// 아무 이름도 안들어온다면, 메시지와 에러를 반환합니다.
	if name == "" {
		return name, errors.New("empty name")
	}

	// 이름이 들어온다면, 이름을 담고 있는 값을 환영 메시지에 담아서 반환합니다.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// 반복문을 통해, 각 이름을 Hello 함수에 넣어서 반환 값인 메시지를 가져온다.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// 이름을 key로 메시지를 value로 맵에 저장한다.
		messages[name] = message
	}
	return messages, nil
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
