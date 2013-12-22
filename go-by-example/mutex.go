package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    // create an empty map
    var state = make(map[int]int)
    // create a new mutual exclusion
    var mutex = &sync.Mutex{}

    var ops int64 = 0

    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                /**
                 * Intn returns, as an int, a non-negative pseudo-random number
                 * in [0,n) from the default Source. It panics if n <= 0.
                 */
                key := rand.Intn(5)
                /**
                 * a locked mutex is not associated to a specific goroutine. It
                 * is allowed for one goroutine to lock a mutex and then arrange
                 * for another to unlock it
                 */
                mutex.Lock()
                total += state[key]
                /**
                 * it is a runtime error if mutex is not locked when trying to
                 * unlock.
                 */
                mutex.Unlock()

                /**
                 * AddInt64 atomically adds delta to *addr and returns the new
                 * value.
                 */
                atomic.AddInt64(&ops, 1)

                /**
                 * Gosched yields the processor, allowing other goroutines to
                 * run. It does not suspend the current goroutine, so execution
                 * resumes automatically.
                 */
                runtime.Gosched()
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    time.Sleep(time.Second)

    // atomically loads *addr
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()

}
