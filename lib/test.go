package lib

import "fmt"

func AssertNoError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func AssertEqual(expected, actual int) {
	if expected != actual {
		fmt.Println(fmt.Sprintf("Test failed \n\texpected: %d, got: %d", expected, actual))
	} else {
		fmt.Println("Test passed")
	}
}
