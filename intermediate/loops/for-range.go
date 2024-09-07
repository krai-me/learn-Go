//Range-based for loop
//This type of for loop is specifically used to iterate over a collection (like slices, arrays, maps, or strings).

package main

import "fmt"

func main() {
	str := "Hello"
	for index, value := range str {
		//Use %c when you specifically want to print a character representation of a byte or rune (e.g., when iterating over strings).
		fmt.Printf("Index: %d, Value: %c\n", index, value)
	}

}
