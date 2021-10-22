package config

import (
	"flag"
)

var port1 int
var port2 int
var tokenDuration int
var xmlPath string
//Other Paths



func Init(){
	flag.IntVar(&port1,"port1",8443,"4-digit number for the port regarding the qrCodeScan")
	flag.IntVar(&port2,"port2", 8444, "4-digit number for the login Page")
	flag.IntVar(&tokenDuration,"tokenDuration", 60, "The life duration of the token")
	flag.StringVar(&xmlPath,"xmlPath","../../assets/locations.xml","Path to the xmlFile storing the locations")
	flag.Parse()
}


