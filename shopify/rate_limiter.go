package shopify

import (
	"time"
)

type rateLimiter struct {
	nextChan chan bool
}

var (
	apiLimit         = time.Second / 2
	shopNameLimitMap = make(map[string]*rateLimiter)
)

func rateLimitFor(shopName string) *rateLimiter {
	if _, ok := shopNameLimitMap[shopName]; !ok {
		shopNameLimitMap[shopName] = &rateLimiter{nextChan: make(chan bool)}
		shopNameLimitMap[shopName].next()
	}
	return shopNameLimitMap[shopName]
}

func (limiter *rateLimiter) next() {
	go func() {
		select {
		case <-time.Tick(apiLimit):
			limiter.nextChan <- true
		}
	}()
}

func (limiter *rateLimiter) Wait() {
	<-limiter.nextChan
	limiter.next()
}
