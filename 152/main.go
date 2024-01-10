package main

import (
	"fmt"
	"log"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

func writeLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	b := book{
		title: "Star Wars",
	}
	writeLog(b)
}
