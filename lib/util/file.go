package util

import (
	"log"
	"os"
)

func WriteToFile(content string, filePath string) {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(content)

	if err != nil {
		log.Fatal(err)
	}
}
