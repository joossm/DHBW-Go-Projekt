package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var path = "../../assets/locations.xml" // absolute Path cause the relative is not working right now
var locations Locations

type Locations struct {
	XMLName   xml.Name   `xml:"locations"`
	Locations []Location `xml:"location"`
}
type Location struct {
	XMLName xml.Name `xml:"location"`
	Name    string   `xml:"name"`
}

func main() {
	locations = read()
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
func read() Locations {
	xmlFile, _ := os.Open(path)
	var locations Locations
	fmt.Println("Successfully Opened " + path)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &locations)
	for i := 0; i < len(locations.Locations); i++ {
		fmt.Println(locations.Locations[i].Name) //for now
	}

	return locations

}
