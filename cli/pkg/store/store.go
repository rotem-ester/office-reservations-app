package store

var s Store

type Store struct {
	BinaryName string
	ServerUrl string
}

func Get() *Store {
	return &s
}

func init() {
	s.BinaryName = "ofre"
	s.ServerUrl = "http://localhost:8080"
}