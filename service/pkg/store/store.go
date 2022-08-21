package store

var s Store

type Store struct {
	Port          string
	DataFilePath  string
	RequestParams []string
}

func Get() *Store {
	return &s
}

func init() {
	s.Port = "8080"
	s.DataFilePath = "./rent_data.txt"
	s.RequestParams = []string{"year", "month"}
}
