package main

func mian() {
	var mixedOperators bool

	mixedOperators = true || false
	mixedOperators = mixedOperators && true
	mixedOperators = !mixedOperators
	//The first true || false evaluates to true. Then true && true is true. Then !true is false.

}
