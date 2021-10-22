package main

import (
	"GoPruefungsaufgabe/src/main/config"
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
func main(){
	config.Init()
	read()
}
func read() Locations {
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
