package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Option 1: Most basic file reading
	file, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))

	// Option 2: os lib has greater control in reading
	f, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	data := make([]byte, 100)
	count, err := f.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Option2: read %d bytes: %s\n", count, string(data))

	// Option 3: Use bufio.NewReader
	// Rewind handle to 0,0 first coz it was read from above example
	_, err = f.Seek(0, 0)

	fileRead := bufio.NewReader(f)
	filePeek, err := fileRead.Peek(120)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Option3:\n%s", string(filePeek))
	f.Close()
}
