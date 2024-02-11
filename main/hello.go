package main

// Go에서 어플리케이션으로 실행되려면 main package 안에 있어야 한당
//

import (
	"fmt"
)

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}
