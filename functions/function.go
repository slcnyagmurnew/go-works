package functions

import "fmt"

func multipleReturns() (int, int) {
	return 5, 2
}

func CallMultipleReturns() {
	result0, result1 := multipleReturns()
	fmt.Printf("Result 1: %d Result 2: %d \n", result0, result1)
}
