package dice_test

import (
	"math/rand"
	"testing"

	. "github.com/clearcodecn/v2ray/common/dice"
)

func BenchmarkRoll1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Roll(1)
	}
}

func BenchmarkRoll20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Roll(20)
	}
}

func BenchmarkIntn1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1)
	}
}

func BenchmarkIntn20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(20)
	}
}
