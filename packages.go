package main
  
import (
    "bytes"
    "fmt"
    "sort"
)
  
func main() {
  
    slice_1 := []byte{'*', 'K', 'a', 'm', 'l', 'e', 's',
        'h', 'K', 'u', 'm', 'a', 'r', 's', '^', '^'}
    slice_2 := []string{"Kam", "le", "shK", "uma", "rs"}
  
    fmt.Println("Original Slice:")
    fmt.Printf("Slice 1 : %s", slice_1)
    fmt.Println("\nSlice 2: ", slice_2) 
    res := bytes.Trim(slice_1, "*^")
    fmt.Printf("\nNew Slice : %s", res)
 
    sort.Strings(slice_2)
    fmt.Println("\nSorted slice:", slice_2)
}
