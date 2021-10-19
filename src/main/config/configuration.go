package config

import (
	"flag"
	"fmt"
)

var port1 int
var port2 int
var tokenDuration int
var XmlFilePath * string
//Other Paths


func init(){
	flag.IntVar(&port1,"port1",8444,"4-digit number for the port regarding the qrCodeScan")
	flag.IntVar(&port2,"port2", 8444, "4-digit number for the login Page")
	flag.IntVar(&tokenDuration,"tokenDuration", 8444, "The life duration of the token")
	XmlFilePath = flag.String("xmlFilePath", "../../assets/locations.xml", "Path to the XML File containing the locations")
	flag.Parse()
	fmt.Println("Path is: " , *XmlFilePath)
}


