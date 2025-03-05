package main

import (
	"fmt"
	"net/http"
)

var Kontakte = []*Kontakt{}

func main() {
	Kontakte = append(Kontakte, &Kontakt{Vorname: "Thorsten", Nachname: "Klöhn"})

	router := http.NewServeMux()
    fmt.Println("Server läuft auf http://localhost:8080")
	_= http.ListenAndServe(":8080", router)

}