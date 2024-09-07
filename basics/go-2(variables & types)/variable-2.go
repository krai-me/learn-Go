package main

import "fmt"

func main() {
	//we can define mutiple varibale of same type or diff at same timeby
	var string1, string2 string
	string1 = "hello"
	string2 = "bye"
	fmt.Println(string1, string2) // hello bye

	//for multiple type
	number, boolean := 12, true
	fmt.Println(number, boolean) // 12 true
}
