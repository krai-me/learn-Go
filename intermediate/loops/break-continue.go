// The break keyword allows the programmer to stop the loop at the current iteration.
// The continue keyword works similarly, allowing the loop to skip to the next iteration.
package main

import (
	"fmt"
)

func main() {

	for count := 0; count < 20; count++ {
		// WRITE CONTINUE STATEMENT IF COUNT EQUALS 8 BELOW THIS LINE
		if count == 8 {
			continue
		}
		// WRITE BREAK STATEMENT IF COUNT EQUALS 15 BELOW THIS LINE
		if count == 15 {
			break
		}
		fmt.Println(count)
	}

}
