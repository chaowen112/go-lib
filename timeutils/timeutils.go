package timeutils

import (
	"time"

	"github.com/chaowen112/go-lib/randutils"
)

func UnixMilli() float64 {
	return float64(time.Now().UnixNano()/1000) / 1000000
}

// RandomDuration returns a random duration between [a, b) for distributing API pressures
//
//	for {
//	    time.Sleep(timeutils.RandomDuration(5*time.Minute, 10*time.Minute)
//	    client.Do(arg1, arg2, arg3)
//	}
//
// above mechanism can avoid DDoS to the API server.
func RandomDuration(a, b time.Duration) time.Duration {
	return a + time.Duration(randutils.Uint32n(uint32((b-a)/time.Millisecond)))*time.Millisecond
}
