package config

import "flag"

var port1 string
var port2 string
var tokenDuration string
var xmlPath string

//Other Paths

func Init() {
	SetFlags("8443","8444","60","assets/locations.xml")
}

func SetFlags(pPort1 string,pPort2 string, pTokenDur string, pXmlPath string,){
	flag.StringVar(&port1, "port1", pPort1, "4-digit number for the port regarding the qrCodeScan")
	flag.StringVar(&port2, "port2", pPort2, "4-digit number for the login Page")
	flag.StringVar(&tokenDuration, "tokenDuration", pTokenDur, "The life duration of the token")
	flag.StringVar(&xmlPath, "xmlPath", pXmlPath, "Path to the xmlFile storing the locations")

}
func GetPath() string {
	return xmlPath
}
