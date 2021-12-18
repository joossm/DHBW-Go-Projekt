package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var locationsList = LocationsList{}

type LocationsList struct {
	XMLName   xml.Name    `xml:"locations"`
	Locations []*Location `xml:"location"`
}
type Location struct {
	XMLName xml.Name `xml:"location"`
	Name    string   `xml:"name"`
}
type AllLocations struct {
	Location []*Location
}

func GetList() LocationsList {
	return locationsList
}

func (l LocationsList) equals(list LocationsList) bool {
	lenA := len(l.Locations)
	lenB := len(list.Locations)
	if lenA != lenB {
		return false
	} else {
		for i := 0; i < lenA; i++ {
			if l.Locations[i].Name != list.Locations[i].Name {
				return false
			}
		}
		return true
	}
}
func (l LocationsList) getLength() int {
	return len(l.Locations)
}
func (l LocationsList) ToStrings() []string {
	locStrings := make([]string, 0)
	for i := 0; i < len(l.Locations); i++ {
		locStrings = append(locStrings, l.Locations[i].Name)
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

func (l LocationsList) ShowAllLoc() []*Location {
	return l.Locations
}
func init() {
}

func ReadXmlFile(path string) LocationsList {
	xmlFile, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	errorHandling(err)
	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {

		}
	}(xmlFile)
	byteValue, err := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &locationsList)
	return locationsList
}
