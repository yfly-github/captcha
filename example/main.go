package main

import (
	"fmt"
	mathRand "math/rand"
	"time"
)

func main() {
	mathRand.Seed(time.Now().UnixNano())
	b := make([]byte, 0)
	for i := 0; i < 10; i++ {
		t := mathRand.Intn(9)
		b = append(b, byte(t))
	}

	fmt.Println("b", b)
}
