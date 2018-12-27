package main

import (
	"os"
	"testing"
)

type FakeReadWrite struct{}

func (rw *FakeReadWrite) Read(filepath string) (string, error) {
	return "fake_content", nil
}

func (rw *FakeReadWrite) Write(filepath string, content string) error {
	return nil
}

// Override the factory method to return FakeReadWrite which satisfies the ReaderWriter
// interface
func TestMain(m *testing.M) {
	oldReadWrite := CreateReadWrite

	CreateReadWrite = func() ReaderWriter {
		rw := &FakeReadWrite{}
		return rw
	}

	ret := m.Run()
	os.Exit(ret)

	CreateReadWrite = oldReadWrite
}

// Testget tests the get method
func TestGet(t *testing.T) {
	content, err := Get()
	if err != nil {
		t.Error(err)
	}

	if content != "fake_content" {
		t.Error(err)
	}
}

// Testput tests the put method
func TestPut(t *testing.T) {
	err := Put()
	if err != nil {
		t.Error(err)
	}
}
