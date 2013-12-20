package main

import "time"
import "fmt"

func main() {

    // create a new ticker that ticks every 500ms
    ticker := time.NewTicker(time.Millisecond * 500)

    go func() {
        // every time the ticker ticks
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
