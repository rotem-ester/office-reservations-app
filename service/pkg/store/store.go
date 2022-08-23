package store

var s store

type store struct {
	Port          string
	DataFilePath  string
	RequestParams []string
}

func Get() *store {
	return &s
}

func init() {
	s.Port = "8080"
	s.DataFilePath = "./rent_data.txt"
	s.RequestParams = []string{"year", "month"}
}
