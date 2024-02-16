package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(getOne())
}

func getOne() string {
	people := []string{
		"주광",
		"병룡",
		"찬욱",
		"원상",
	}

	return people[rand.Intn(len(people))]
}
