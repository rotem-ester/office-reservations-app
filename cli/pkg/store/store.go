package store

var s store

type store struct {
	BinaryName string
	ServerUrl string
}

func Get() *store {
	return &s
}

func init() {
	s.BinaryName = "ofre"
	s.ServerUrl = "http://localhost:8080"
}