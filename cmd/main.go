package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// struct correspondiente al json esperado
type data := struct {
	Origin string
	User string
	Active bool
}

func main(){
	content, err := ioutil.ReadFile(./training.json)
	if err != nil {
		log.Fatal("Error, could not read file")
	}

	var payload
}