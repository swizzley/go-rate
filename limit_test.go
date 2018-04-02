package rate

import (
	"testing"
	"time"
)

func TestLimits(t *testing.T) {
	rate, err := Limits(1, time.Second)
	if err != nil {
		t.Fail()
	}

	rate.Add()

	if len(rate.count) != 1 {
		t.Fail()
	}

	if rate.rate != time.Second {
		t.Fail()
	}
}
