package main

import (
	"fmt"
	"math/rand"
	"time"
)

var word = []string{
	"testing",
	"no single quotes",
	"more words",
	"ugh",
	"cheese barn",
	"will it work",
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Hello World!")
	fmt.Println(genWord())
}

func genWord() string {
	return word[rand.Intn(len(word))+0]
}
