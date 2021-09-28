package glogger

import (
	"fmt"
	"testing"
)

func BenchmarkInfo(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Info("Test" + fmt.Sprint(n))
	}
}