package main

import "time"
import "fmt"

func main() {

    // create requests channel
    requests := make(chan int, 5)

    // add 5 requests to requests channel
    for i := 1; i <= 5; i++ {
        requests <- i
    }

    // close request channel
    close(requests)

    // create a rate limiter
    limiter := time.Tick(time.Millisecond * 200)

    for req := range requests {
        // wait for a tick before printing request
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 100)

    for i := 0; i < 3; i++ {
        // add 3 items to burstyLimiter channel
        burstyLimiter <- time.Now()
    }

    go func() {
        // every 200ms add a new value to burstyLimiter
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 100)

    for i := 1; i <= 100; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
