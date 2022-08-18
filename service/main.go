package main

import (
	"fmt"
	"log"
)

func main(){
	reservations, err := LoadCsv("./rent_data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(reservations)
}