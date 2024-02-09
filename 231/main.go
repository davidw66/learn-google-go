package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
		fmt.Println("check log.txt file in the directory")
	}
	defer f2.Close()

}
