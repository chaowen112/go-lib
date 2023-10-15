package randutils

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandShuffle(t *testing.T) {
	a := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	t.Logf("%s", a)
}

func BenchmarkRandShuffleStd(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.SetParallelism(1000)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		a := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
		for pb.Next() {
			rand.Shuffle(len(a), func(i, j int) {
				a[i], a[j] = a[j], a[i]
			})
		}
	})
}

func BenchmarkRandShuffle(b *testing.B) {
	b.SetParallelism(1000)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		a := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
		for pb.Next() {
			Shuffle(len(a), func(i, j int) {
				a[i], a[j] = a[j], a[i]
			})
		}
	})
}
