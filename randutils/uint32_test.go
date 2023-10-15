package randutils

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRandUint32Std(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.SetParallelism(1000)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Intn(100)
		}
	})
}

func BenchmarkRandUint32(b *testing.B) {
	b.SetParallelism(1000)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Uint32n(100)
		}
	})
}
