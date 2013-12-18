package main

import "fmt"
import "math"

const s string = "constant"

func main() {

  // print the value if s
  fmt.Println(s)

  // const can be defined anywhere a var statement can
  const n = 50000000

  // constant expressions perform arithmatic with arbitrary precision
  const d = 3e20 / n
  fmt.Println(d)

  // numeric constant has no type until it's given one
  fmt.Println(int64(d))

  // a number can be given a type by using it in a context that requires one
  fmt.Println(math.Sin(n))
}
