package main

import "time"
import "fmt"

func main() {

    // create a new 2 second timer
    timer1 := time.NewTimer(time.Second * 2)

    // wait until time expires & print message
    <-timer1.C
    fmt.Println("timer 1 expired")

    // create a new 1 second timer
    timer2 := time.NewTimer(time.Second)

    // create a new closure
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()

    // stop second timer before it expires
    stop2 := timer2.Stop()

    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
