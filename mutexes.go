package main 

import (
	"fmt"
	"math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
	var state = make(map[int]int)
	
	var mutex = &sync.Mutex{}

	var reads uint64
	var writes uint64

	for r := 0; r < 100; r++ {
		go func () {
			total := 0
			for {
				key := rand.Intn(5)
				// read before locking
				mutex.Lock()
				total = total + state[key]

				// unlock after reading done
				mutex.Unlock()
				atomic.AddUint64(&reads, 1)

				time.Sleep(time.Millisecond)

			}
		}()
	}

	for w := 0; w < 30; w++ {
		go func () {
			// taking a key randomly and assigning a random value
			key := rand.Intn(5)
			val  := rand.Intn(100)

			// lock the mutex
			mutex.Lock()

			state[key] = val	
			
			// unlock after writing work donw
			mutex.Unlock()

			atomic.AddUint64(&writes, 1)
			time.Sleep(time.Millisecond)
			
		}()
	}

	time.Sleep(time.Second)

}