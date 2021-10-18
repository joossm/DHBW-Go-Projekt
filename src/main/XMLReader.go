package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var path = "assets/locations.xml"

type Locations struct{
	XMLName xml.Name      `xml:"locations"`
	Locations  []Location `xml:"location"`
}
type Location struct{
	XMLName xml.Name `xml:"location"`
	Name string `xml:"name"`
}

func main(){
	xmlFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + path)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var locations Locations

	xml.Unmarshal(byteValue, &locations)

	for i := 0; i < len(locations.Locations); i++{
		fmt.Println("Location: " + locations.Locations[i].Name)
	}
}

