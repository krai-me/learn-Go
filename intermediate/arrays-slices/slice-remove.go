//slice dont comes with inbuilt remove of an element we can do by logic

package main

import "fmt"

func main() {
	num := []int{23, 1, 10, 12} //let say we have remove 10
	fmt.Println("Initial slices: ", num)
	remove(num, 89)
}
func check(array []int, value int) {
	for i := 0; i < len(arrayarray); i++ {

	}

}
func remove(array []int, value int) {
	for index, values := range array {
		if value != array[index] {
			fmt.Println("pass an appropriate value")
		} else if values != value {
			fmt.Println(value)
		}
	}
}
