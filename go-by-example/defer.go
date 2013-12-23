/**
 * DEFER
 *
 * A defer statement pushes a function call onto a list. The list of
 * saved calls is executed after the surrounding function returns.
 *
 * Rules
 * 1. A deferred functions arguments are evaluated when the defer statement is evaluated
 * 2. Deferred function calls are executed Last In First Out order after the surrounding function returns
 * 3. Deferred functions may read and assign the returning functions named return values.
 * 
 * Resources
 * 1. http://blog.golang.org/defer-panic-and-recover
 */
package main

import "fmt"
import "os"

func main() {
    f := createFile("./defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
