// A map is an unordered collection of keys and values.
package main

import "fmt"

func main() {
	//to create an empty map we can use
	//variableName := make(map[keyType]valueType)
	prices := make(map[string]float32)
	fmt.Println(prices) //map[]

	/*
			variableName := map[keyType]valueType{
		    name1: value1,
		    name2: value2,
		    name3: value3,
		}
	*/
	age := map[string]int{
		"ash":   10,
		"brock": 15,
		"Misty": 10,
	}
	fmt.Println(age) //map[Misty:10 ash:10 brock:15]

	//to access values in the map
	//variable := yourMap[keyValue]
	//fmt.Println(variable)
	ash_age := age["ash"]
	fmt.Println(ash_age) //10
}
