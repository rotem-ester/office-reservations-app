package store

var s Store

type Store struct {
	BinaryName string
	ServerHost string
}

func Get() *Store {
	return &s
}

func init() {
	s.BinaryName = "ofre"
	s.ServerHost = "localhost:8080"
}