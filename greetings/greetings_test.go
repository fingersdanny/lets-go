package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName은 name 변수를 넣어서 greetings.Hello를 호출하고 정상 반환값을 반환하는 지 테스트 합니다.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want, _ := regexp.Compile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty 빈 스트링 변수를 넣어서 greetings.Hello를 호출하고 에러를 확인한다.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
