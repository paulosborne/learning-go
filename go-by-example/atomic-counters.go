package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

    // unsigned integer (non-negative)
    var ops uint64 = 0

    for i := 0; i < 50; i++ {
        go func() {
            for {
                // AddUint64 atomically adds delta to *addr and returns the new value
                atomic.AddUint64(&ops, 1)

                // Gosched yields the processor, allowing other goroutines to run.
                // It does not suspend the current goroutine, so execution resumes automatically.
                runtime.Gosched()
            }
        }()
    }

    // wait a second to allow some ops to accumulate
    time.Sleep(time.Second * 2)

    // LoadInt64 atomically loads *addr.
    opsFinal := atomic.LoadUint64(&ops)

    fmt.Println("ops", opsFinal)

}
