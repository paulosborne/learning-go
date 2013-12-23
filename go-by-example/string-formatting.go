package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}

    fmt.Printf("%v\n", p)       // {1 2}
    fmt.Printf("%+v\n", p)      // {x:1 y:2}
    fmt.Printf("%#v\n", p)      // main.point{x:1, y:2}
    fmt.Printf("%T\n", p)       // main.point

    // formatting booleans
    fmt.Printf("%t\n", true)    // true

    // There are many options for formatting integers. Use %d for standard, base-10 formatting.
    fmt.Printf("%d\n", 123)     // 123

    // This prints a binary representation.
    fmt.Printf("%b\n", 14)      // 1110

    // this prints the character corresponding to the given integer
    fmt.Printf("%c\n", 33)      // !

    // %x provides hex encoding
    fmt.Printf("%x\n", 456)     // 1c8

    // There are also several formatting options for floats.
    fmt.Println("%f\n", 78.9)

    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    fmt.Printf("%s\n", "\"string\"")
    fmt.Printf("%q\n", "\"string\"")

    fmt.Printf("%x\n", "hex this")

    fmt.Printf("%p\n", &p)
    
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)    // |  1.20|  3.45|

    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
