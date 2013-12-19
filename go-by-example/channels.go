package main

import "fmt"

func main() {

    // create a new messages channel
    messages := make(chan string)

    // send a string into the message channel
    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}
