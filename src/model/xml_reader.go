package model

import (
	"GoPruefungsaufgabe/src/model/config"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Locations struct {
	XMLName   xml.Name   `xml:"locations"`
	Locations []Location `xml:"location"`
}
type Location struct {
	XMLName xml.Name `xml:"location"`
	Name    string   `xml:"name"`
}

func (this Locations) GetLocations() []Location {
	return this.Locations
}
func (this Locations) getLength() int {
	return len(this.Locations)
}
func (Location) printName() string {
	return "To Do"
}
func main() {
	config.Init()
	Read()
}
func Read() Locations {
	var locations Locations
	var path = flag.Lookup("xmlPath").Value.String()
	xmlFile, _ := os.Open(path)
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &locations)

	for i := 0; i < len(locations.Locations); i++ {
		fmt.Println(locations.Locations[i].Name)
	}

	return locations

}
func ShowAllLocations() (au *Locations) {
	file, err := os.OpenFile("assets/locations.xml", os.O_RDWR|os.O_APPEND, 0666)
	errorHandling(err)
	b, err := ioutil.ReadAll(file)
	var locations Locations
	json.Unmarshal(b, &locations.Locations)

	return &locations
}
func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
