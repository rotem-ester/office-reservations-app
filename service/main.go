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
	ors.ParseData(data)
	fmt.Print(ors.Reservations)
}