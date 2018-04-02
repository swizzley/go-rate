package main

import (
	"github.com/swizzley/rate"
	"time"
	"fmt"
)

func main() {

	rate, err := rate.Limits(2, time.Second)
	if err != nil {
		panic(err)
	}

	s := 1
	for {

		rate.Add()

		t := time.Now().Format(time.RFC3339Nano)
		fmt.Println(t, s)

		s++
	}
}
