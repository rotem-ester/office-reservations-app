package store

var s Store

type Store struct {
	DataFilePath string
	RevenueRequestParams []string
}

func Get() *Store {
	return &s
}

func init() {
	s.DataFilePath = "./rent_data.txt"
	s.RevenueRequestParams = []string{ "year", "month" }
}