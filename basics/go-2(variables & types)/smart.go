package main

import "fmt"

func main() {
	//to avoid using var we use Inferring Type
	//There is a way to declare a variable without explicitly stating its type
	//using the short declaration `:=` operator.

	ukid := false       //defines general bool
	snack := "snack"    //it defines  general string type
	inthuh := 1212      //defines general int
	floathuh := 9023.23 //defines general float64
	fmt.Println(ukid, snack, inthuh, floathuh)
	//we can define var without stating types
	var wekid = false
	fmt.Println(wekid)

	//It’s recommended to use int unless there’s a reason to specify
	//the size of the int (like knowing that value will be larger than the default type, or needing to optimize the amount of space used).
	var number int = 10
	fmt.Println(number)

	var basketTotal float64
	spinachPrice := 1.50
	basketTotal += spinachPrice
	fmt.Println(basketTotal) // Prints: 2.25

}
