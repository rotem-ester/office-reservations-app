package main

import (
	"fmt"
	"log"
	"os"

	or "github.com/rotem-ester/office-reservation-app/service/office_reservation"
	fileUtil "github.com/rotem-ester/office-reservation-app/service/pkg/file_util"
	"github.com/rotem-ester/office-reservation-app/service/pkg/store"
)

func main(){
	var ors or.OfficeReservationService

	f, err := os.Open(store.Get().DataFilePath)
	if err != nil {
		log.Fatalf("failed to open csv file. error: %s", err.Error())
	}

	data, err := fileUtil.LoadCsv(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
	err = ors.ParseData(data)
	if err != nil {
		log.Fatalf("failed to load data from csv file: %s", err.Error())
	}
	log.Println("data was loaded from file successfully")

	fmt.Print(ors.Reservations)
}