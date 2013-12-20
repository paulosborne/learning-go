package main

import "fmt"

func main() {

    // create a new channel, which holds 2 values
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    // closing the channel prevents the range from locking
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}
