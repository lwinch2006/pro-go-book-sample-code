package benchmarks

import (
	"chapter31/services"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Benchmark_UnitTesting1_1(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 1000
	testValues := make([]int, size)

	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			testValues[j] = rand.Int()
		}

		services.UnitTesting1_1(testValues)
	}
}

func Benchmark_UnitTesting1_1_ResettingTimer(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 1000
	testValues := make([]int, size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < size; j++ {
			testValues[j] = rand.Int()
		}

		b.StartTimer()
		services.UnitTesting1_1(testValues)
	}
}

func Benchmark_UnitTesting1_1_DifferentSliceSizes(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{250, 500, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Running benchmark for size %v", size), func(subB *testing.B) {
			testValues := make([]int, size)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				subB.StopTimer()
				for j := 0; j < size; j++ {
					testValues[j] = rand.Int()
				}

				subB.StartTimer()
				services.UnitTesting1_1(testValues)
			}
		})
	}
}
