package errors

import (
	"errors"
	"fmt"
)

func newError(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("arg can not be less than zero")
	}
	return arg * 2, nil
}

type funcError struct {
	arg    int
	reason string
}

func newerError(arg int) (int, error) {
	if arg < 0 {
		return -1, &funcError{arg, "arg is not proper"}
	}
	return arg + 3, nil
}

func (e *funcError) Error() string {
	return fmt.Sprintf("error code: %d error reason: %s \n", e.arg, e.reason)
}

func CallErrorFunctions() {
	for _, i := range []int{7, -3} {
		if r, e := newError(i); e != nil {
			fmt.Println("err func failed:", e)
		} else {
			fmt.Println("err func worked:", r)
		}
	}

	for _, i := range []int{7, -3} {
		if r, e := newerError(i); e != nil {
			fmt.Println("newer err func failed:", e)
		} else {
			fmt.Println("newer err func worked:", r)
		}
	}
}
