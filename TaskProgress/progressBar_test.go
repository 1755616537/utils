package TaskProgress

import (
	"testing"
	"time"
)

func Test_progressBar(t *testing.T) {
	b := NewBar(0, 1000)
	for i := 0; i < 1000; i++ {
		b.Add(1)
		time.Sleep(time.Microsecond * 100)
	}
}
