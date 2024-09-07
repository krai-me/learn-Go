//if a import is done/var is created but never used
//it will raise err

package main

import "fmt"

//./err-1.go:5:8: "fmt" imported and not used

func main() {
	var amount int32 //./err-1.go:8:6: amoutn declared and not used

	var pokemon string = "pika"

	pokemon = 12 // throws err as we  try to defining int to a string variables

	var int1 = 12
	var float1 float64 = 12.12
	fmt.Println(int1 + float1) //throws err as we can sum onyl same datatypes even its int8 and int16 u 	cant

}
