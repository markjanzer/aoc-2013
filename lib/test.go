package lib

import "fmt"

func AssertNoError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func AssertEqual[T comparable](expected, actual T) {
	if expected != actual {
		fmt.Println("Test failed \n\texpected: ", expected, " got: ", actual)
	} else {
		fmt.Println("Test passed")
	}
}
