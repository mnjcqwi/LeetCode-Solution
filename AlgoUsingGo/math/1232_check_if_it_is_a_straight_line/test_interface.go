package main

//Retriver need to implement get method
type Retriver interface {
	Get(url string) string
}

