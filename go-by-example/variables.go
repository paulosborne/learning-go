package main

import "fmt"

func main() {

  var a string = "initial"
  fmt.Println(a)

  // declares one or more variables at once
  var b, c int = 1, 2
  fmt.Println(b, c)

  // infer variable type when declared
  var d = true
  fmt.Println(d)

  // declares e with zero-value
  var e int
  fmt.Println(e)

  // use short syntax declaration
  // var f string = "short"
  f := "short"
  fmt.Println(f)

}
