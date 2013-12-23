package main

import (
    "fmt"
    "flag"
    s "strings"
)

func save (p int) {
    fmt.Println(p)
}

func main() {
    // create a slice to hold the attendees
    people := make([]string, 10)

    // register a flag
    attendees := flag.String("att","foo","A comma seperated list of standup attendees")

    // parse command line flags
    flag.Parse()

    // save the list of attendees

    fmt.Println("args:", flag.Args())
    fmt.Println(people)
}
