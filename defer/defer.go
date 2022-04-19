package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	defer fmt.Println("Hello")
	defer fmt.Println("Hello 2")

}

func CreateFile(path string) error {
	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		return err
	}

	return nil
}

func WriteFile(path string, content []byte) error {

	if err := os.WriteFile(path, content, os.ModeAppend); err != nil {

		return err
	}

	return nil
}

func readFile(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
