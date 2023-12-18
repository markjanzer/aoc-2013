package lib

import (
	"io"
	"os"
)

func GetDataString(fileName string) string {
	file, err := os.Open(fileName)
	AssertNoError((err))
	defer file.Close()

	// Read file content into a byte slice
	byteContent, err := io.ReadAll(file)
	AssertNoError(err)

	return string(byteContent)
}
