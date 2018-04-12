package rate

import (
	"time"
	"errors"
)

// A Limit is the number of cycles allowed during a given duration along with the current count of cycles in
// current duration. The duration begins when the first cycle is added to the cycle count.
type Limit struct {
	rate  time.Duration
	limit int
	count chan int
}

// Returns a Rate Limit variable after creating a count channel, and creates a go routine to de-increment the count
// based on the rate limit specified by limit (int) per rate (time.Duration)
func Limits(limit int, rate time.Duration) (r Limit, err error) {

	if limit <= 0 {
		err = errors.New("rate has no limit")
	}
	if rate < time.Millisecond {
		err = errors.New("rate limit is less than 1 millisecond")
	}

	r.limit = limit
	r.rate = rate

	r.count = make(chan int, limit)

	go limiter(r)

	return
}

// Add to the rate limit count
func (r Limit) Add() {
	r.count <- 1
}

// limiter de-increments the Count at the specified Rate duration
func limiter(r Limit) {
	for {
		if len(r.count) > 0 {
			time.Sleep(r.rate)
		}
		for i := 0; i < len(r.count); i++ {
			<-r.count
		}
	}
}
