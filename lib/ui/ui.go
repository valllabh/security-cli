package ui

import (
	"fmt"
	"log"
)

func Step(name string) {
	fmt.Println("Step: " + name)
}

func Fatal(err error) {
	log.Fatal(err)
}
