package store

var s Store

type Store struct {
	DataFilePath string
}

func Get() *Store {
	return &s
}

func init() {
	s.DataFilePath = "./rent_data.txt"
}