package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Example Watch
func ExampleRun() {
	ParseFlags()
	FC.Daemon = false
	Run()
	// Output:
	// Log file config path: watch.json5
	// Flag `daemon` set to false.  Running commands in config and exiting.
}

// Create a file, write to it, and delete it.
func touch() {
	fileName := "temp.txt"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create("temp.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(fileName, currentTime, currentTime)
		if err != nil {
			fmt.Println(err)
		}
	}

	b := []byte("hello watch!")
	err = ioutil.WriteFile(fileName, b, 0644)
	if err != nil {
		fmt.Println("Could not write to: " + fileName)
		return
	}

	e := os.Remove(fileName)
	if e != nil {
		fmt.Println("Could not remove: " + fileName)
	}
}
