// Go uses a default seed value of 1 which can lead to predictable numbers being generated. We can generate random numbers by using different seed values.
package main

import (
	"fmt"
	"math/rand" // Importing math/rand for random number generation
	"time"      // Importing time to get the current time for seeding
)

// Default Seed: Without seeding, you might see different numbers by chance, but it’s not reliable for applications requiring true randomness.
func main() {
	// Seeding the random number generator using the current time in nanoseconds
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	// Generating and printing a random number between 0 and 99
	fmt.Println("Random number:", rand.Intn(100))
}

/*
Imagine you have a magic box that can give you a random number, but it needs a special key to start. If you use the same key every time, you get the same number. That’s not very fun!

To make sure you get a different number each time you open the magic box, you need to use a different key each time.

Here’s how we use time as a key:

Get the current time: This is like saying, "What time is it right now?" The time changes every moment, so it’s always different.
Use that time as the key: This key helps the magic box pick a different number every time you open it.
*/
