package functions

import "fmt"

func multipleReturns() (int, int) {
	return 5, 2
}

func CallMultipleReturns() {
	result0, result1 := multipleReturns()
	fmt.Printf("Result 1: %d Result 2: %d \n", result0, result1)
}

func variadic(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Total: %d \n", total)
}

func CallVariadic() {
	numbers := []int{4, 4, 5, 1}
	variadic(numbers...)
}

func valueFunc(value int) {
	value = 0
}

func pointerFunc(value *int) {
	*value = 0
}

func CallPointerValue() {
	i := 1

	valueFunc(i)
	fmt.Printf("Updated value by reference: %d \n", i)

	pointerFunc(&i)
	fmt.Printf("Updated value by pointer: %d \n", i)
}

