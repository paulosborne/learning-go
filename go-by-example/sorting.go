/**
 * Go's sprt package implements sorting for builtins and user-defined types. We'll
 * look at sorting for builtins first
 */
package main

import "fmt"
import "sort"

func main() {

    strs := []string{"c","a","b"}

    // sorts a slice of strings in increasing order.
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7,2,4}
    // sorts a slice of ints in increasing order.
    sort.Ints(ints)
    fmt.Println("Ints:", ints)

    // IntsAreSorted tests whether a slice of ints is sorted in increasing order.
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
