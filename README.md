# rate
Rate Limits for your API

## Overview

This package is just to add generic rate limits on whatever you need. It uses channels as the built-in limiter.

## Install

`go get github.com/swizzley/rate`


## Usage

  * `rate.Limits` takes two parameters and returns an error or a Rate Limit incremented with the Add() method
  * `rate.Add()` just increments the Rate Limit count

## Example

In the example below I want to limit printing the time to twice per second, so `rate.Limits(2, time.Second)`

When using the limiter in my print loop, I add 1 cycle to the count with `rate.Add()`

```
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
```

