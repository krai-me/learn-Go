package main

import "fmt"

func main() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	changeLastElement(myTutors, "Bobby")
	fmt.Println(myTutors)
}

func changeLastElement(array []string, value string) {
	len := len(array)
	if len > 0 {
		array[len-1] = value
	}
}
