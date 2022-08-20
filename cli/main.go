package main

import (
	"log"

	"github.com/rotem-ester/office-reservations-app/cli/cmd"
)

func main() {
	c := cmd.NewRoot()
	err := c.Execute()
	if err != nil {
		log.Fatal(err)
	}
}