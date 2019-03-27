package lib

import (
	"math/rand"
	"time"
)

// both min and max are nano-second
func BackoffGen(min int32, max int32)func(){
	//rand.Seed(time.Now().Unix())
	var limit = min
	return func() {
		delay := rand.Int31n(limit)
		if 2*limit < max{
			limit = 2*limit
		}else{
			limit = max
		}
		time.Sleep(time.Duration(delay)*time.Nanosecond)
	}
}

