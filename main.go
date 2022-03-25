package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Envelope struct {
	XmlName xml.Name `xml:"Envelope"`
	Subject string   `xml:"subject"`
	Sender  Sender   `xml:"Sender"`
	Cube    Cube     `xml:"Cube"`
}

type Sender struct {
	XmlName xml.Name `xml:"Sender"`
	Name    string   `xml:"name"`
}

type Cube struct {
	XmlName xml.Name `xml:"Cube"`
	Cube    []Cube1  `xml:"Cube"`
}

type Cube1 struct {
	XmlName xml.Name `xml:"Cube"`
	Time    string   `xml:"time,attr"`
	Cube    []Cube2  `xml:"Cube"`
}

type Cube2 struct {
	Xmlname  xml.Name `xml:"Cube"`
	Currency xml.Attr `xml:"currency"`
	Rate     xml.Attr `xml:"rate"`
}

func getdata() (Envelope, error) {
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		return Envelope{}, err
	}

	defer resp.Body.Close()

	var envelope Envelope

	if err := xml.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return envelope, err
	}

	return envelope, nil

}

func main() {
	envelope, err := getdata()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("Subject: ", envelope.Subject)
	fmt.Print("Sender: ", envelope.Sender)

	fmt.Print("Sender: ", envelope.Cube.Cube)
	for _, cube := range envelope.Cube.Cube {
		fmt.Print("Time:", cube.Time)
	}
}
