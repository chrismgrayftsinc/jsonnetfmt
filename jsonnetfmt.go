package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-jsonnet/formatter"
)

func main() {
	contentChanged := false
	args := os.Args
	for i := 1; i < len(args); i++ {
		contentBytes, err := ioutil.ReadFile(args[i])
		if err != nil {
			log.Fatal(err)
		}
		content := string(contentBytes)
		newContent, err := formatter.Format(args[i], content, formatter.DefaultOptions())
		if err != nil {
			log.Fatal(err)
		}
		if content != newContent {
			err := ioutil.WriteFile(args[i], []byte(newContent), 0644)
			if err != nil {
				log.Fatal(err)
			}
			contentChanged = true
		}

	}
	if contentChanged == true {
		os.Exit(1)
	}
	os.Exit(0)
}
