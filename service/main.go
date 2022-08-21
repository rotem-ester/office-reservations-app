package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	or "github.com/rotem-ester/office-reservation-app/service/office_reservation"
	fileUtil "github.com/rotem-ester/office-reservation-app/service/pkg/file_util"
	"github.com/rotem-ester/office-reservation-app/service/pkg/store"
)

func main(){
	var ors or.OfficeReservationService
	port := store.Get().Port
	dataFilePath := store.Get().DataFilePath

	args := os.Args[1:]
	if len(args) > 0 {
		port = args[0]
	}

	if len(args) > 1 {
		dataFilePath = args[1]
	}

	data, err := fileUtil.LoadCsv(dataFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = ors.ParseData(data)
	if err != nil {
		log.Fatalf("failed to load data from csv file: %s", err.Error())
	}
	log.Println("data was loaded from file successfully")

	http.HandleFunc("/revenue", ors.RevenueHandler)
	http.HandleFunc("/capacity", ors.CapacityHandler)

	log.Printf("Starting to listen on port %s\n", port)
	formatPort := fmt.Sprintf(":%s", port)
	err = http.ListenAndServe(formatPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}