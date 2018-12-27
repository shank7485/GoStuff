package main

import (
	"fmt"
)

// Get is used to read a file
func Get() (string, error) {
	rw := CreateReadWrite()
	content, err := rw.Read("sample.txt")
	if err != nil {
		return "", err
	}
	return content, nil
}

// Put is used to write to a file
func Put() error {
	rw := CreateReadWrite()
	err := rw.Write("sample.txt", "This is the content!")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Put()
	if err != nil {
		fmt.Printf("put errored: %s", err.Error())
	}

	content, err := Get()
	if err != nil {
		fmt.Printf("get errored: %s", err.Error())
	}

	fmt.Println("#############")
	fmt.Println(content)
	fmt.Println("#############")
}
