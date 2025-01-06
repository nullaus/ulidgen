package main

import (
	"crypto/rand"
	"fmt"

	"github.com/oklog/ulid/v2"
)

var entropy = ulid.Monotonic(rand.Reader, 0)

func main() {
	id, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println(id.String())
}
