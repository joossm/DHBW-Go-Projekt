// 5807262
// 9899545
// 8622410

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

// AllLocations is a struct that contains all the locations
type AllLocations struct {
	Location []*Location
}

// GetList returns a list of all locations
func GetList() LocationsList {
	return locationsList
}

// equals is checking if two lists are equal
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

// getLength returns the length of the list
func (l LocationsList) getLength() int {
	return len(l.Locations)
}

// ToString returns a string representation of the list
func (l LocationsList) ToStrings() []string {
	locStrings := make([]string, 0)
	for i := 0; i < len(l.Locations); i++ {
		locStrings = append(locStrings, l.Locations[i].Name)
	}
	return locStrings
}

// error handling
func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Gets the Locationslist and returns a LocationsArray
func (l LocationsList) ShowAllLoc() []*Location {
	return l.Locations
}

// Reads the XML file and returns a LocationsList
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
