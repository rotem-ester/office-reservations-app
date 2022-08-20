package store

var s Store

type Store struct {
	DataFilePath string
	RequestParams []string
}

func Get() *Store {
	return &s
}

func init() {
	s.DataFilePath = "./rent_data.txt"
	s.RequestParams = []string{ "year", "month" }
}