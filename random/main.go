package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func mathRand() int {
	rand.Seed(123456789)
	return rand.Int()
}

func cryptRand() int {
	var s int64
	// mac はリトルエンディアンだったので
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
	return rand.Int()
}

func main() {
	// first try
	fmt.Println("math/rand 1st: ", mathRand())
	fmt.Println("crypto/rand 1st: ", cryptRand())

	// second try
	fmt.Println("math/rand 2nd: ", mathRand())
	fmt.Println("crypto/rand 2nd: ", cryptRand())
}
