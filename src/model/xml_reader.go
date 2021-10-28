package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var locationsList = LocationsList{}
type LocationsList struct {
	XMLName   xml.Name   `xml:"locations"`
	Locations []*Location `xml:"location"`
}
type Location struct {
	XMLName xml.Name `xml:"location"`
	Name    string   `xml:"name"`
}
type AllLocations struct {
	Location []*Location
}


func (this LocationsList) getLength() int {
	return len(this.Locations)
}
func (this LocationsList) ToStrings() []string {
	locStrings := make([]string, 0)
	for i := 0; i < len(this.Locations); i++ {
		locStrings = append(locStrings, this.Locations[i].Name)
	}
	return locStrings
}
func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
/*
func ShowAllLocations() (au *AllLocations) {
	file, err := os.OpenFile(flag.Lookup("xmlPath").Value.String(), os.O_RDWR|os.O_APPEND, 0666)
	errorHandling(err)
	all, err := ioutil.ReadAll(file)
	var allLoc AllLocations
	err = xml.Unmarshal(all, &allLoc.Location)
	if err != nil {
		return nil
	}
	return &allLoc
}
*/


func ShowAllLoc() []*Location{
	return locationsList.Locations
}


func RegisterLocations(path string) LocationsList {
	fmt.Println(path)
	xmlFile, err := os.OpenFile(path,os.O_RDWR|os.O_APPEND, 0666)
	errorHandling(err)
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &locationsList)
	return locationsList
}
