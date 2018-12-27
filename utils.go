package main

import (
	"bufio"
	"fmt"
	"os"
)

// ReaderWriter describes the Reading and Writing characteristic
type ReaderWriter interface {
	Read(string) (string, error)
	Write(string, string) error
}

// ReadWrite is concrete implementation of ReaderWriter interface
type ReadWrite struct{}

// Read reads content from file
func (rw *ReadWrite) Read(filepath string) (string, error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(file)

	version, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	return string(version), nil
}

// Write writes content into file
func (rw *ReadWrite) Write(filepath string, content string) error {
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}

	fmt.Fprintf(file, content)
	return nil
}

// CreateReadWrite is a factory which returns the ReaderWriter type interface
var CreateReadWrite = func() ReaderWriter {
	rw := &ReadWrite{}
	return rw
}
