package model

import (
	"GoProjekt/src/model/config"
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
func (this Locations) toStrings() []string {
	locStrings := make([]string, 0)
	for i := 0; i < len(this.Locations); i++ {
		locStrings = append(locStrings, this.Locations[i].Name)
	}
	return locStrings
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
		println(locations.Locations[i].Name)
	}
	return locations
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ShowAllLocations() (au *AllLocations) {
	file, err := os.OpenFile("assets/locations.xml", os.O_RDWR|os.O_APPEND, 0666)
	errorHandling(err)
	all, err := ioutil.ReadAll(file)
	var allLoc AllLocations
	err = xml.Unmarshal(all, &allLoc.Location)
	if err != nil {
		return nil
	}
	return &allLoc
}

type AllLocations struct {
	Location []*Locations
}

func RegisterLocations() []string {
	var locations Locations
	var path = flag.Lookup("xmlPath").Value.String()
	xmlFile, _ := os.Open(path)
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &locations)
	return locations.toStrings()
}
