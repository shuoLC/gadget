package main

import (
	"log"
	"tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln("cmd.Execute err: %v", err)
	}
}