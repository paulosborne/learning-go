package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// http://golang.org/ref/spec#Struct_types
type readOp struct {
    key int
    resp chan int
}

// - a struct us a sequence of named elements, called fields, each of which has a name and a type.
// - field names must be unique
// - A field declared with a type but no explicit field name is an anonymous field
type writeOp struct {
    key int
    val int
    resp chan bool
}

func main() {

    var ops int64 = 0

    // create a read channel that accepts an instance of readOp
    reads  := make(chan *readOp)
    // create a write channel that accepts an instance of writeOp
    writes := make(chan *writeOp)

    go func() {
        // create a private map to handle state
        var state = make(map[int]int)
        for {
            /**
             * a select statement chooses which of a set of possible communications will proceed.
             * It looks similar to a "switch" statement but with the cases all referring to
             * communication operations
             */
            select {
            /**
             * All cases in a select statement are evaluated in top-to-bottom order, along with
             * any expressions that appear on the right ahdn side of the send statements
             */
            case read := <-reads:
                read.resp <- state[read.key]
            // The receive case may declare one or two new variables using a short variable declaration.
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp {
                    key: rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
            }
        }()
    }

    // create 10 goroutines
    for w := 0; w < 10; w++ {
        go func() {
            // for loop without a condition loops repeatedly
            for {
                /**
                 * create a new write operation using a struct
                 */
                write := &writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(100),
                    resp: make(chan bool)}

                // write to channel
                writes <- write
                // block until write.resp equals true
                <-write.resp
                // increment ops atomically
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

}
