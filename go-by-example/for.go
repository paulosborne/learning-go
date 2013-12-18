package main

import "fmt"

func main() {
  i := 1
  // most basic type, with a single condition
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // a classic initial/condition/after
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  // a for condition will loop repeatedly until you break out of the loop
  // or return from the enclosing function
  for {
    fmt.Println("loop")
    break
  }
}
