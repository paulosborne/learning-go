package main

import "sort"
import "fmt"

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}


