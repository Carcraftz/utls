package tls

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zhangyunhao116/fastrand"
)

func RandomizeSeed() int64 {
	return time.Now().UnixNano()*int64(fastrand.Int()) - int64(fastrand.Int()) + int64(fastrand.Int())/2
}

func RandomInt(min int, max int) int {
	go func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	seed := rand.NewSource(RandomizeSeed())

	if max < 0 || max < min {
		return min
	}

	randInt := rand.New(seed)

	return randInt.Intn((max+1)-min) + min
}
