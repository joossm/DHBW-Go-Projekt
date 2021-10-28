package config

import (
	"flag"
	"time"
)

var port1 string
var port2 string
var tokenDuration time.Duration
var xmlPath string
var logfilePath string
var loginUrl string
var endUrl string
var locationUrl string
var logoutUrl string
var logoutPagePath string
var loginPagePath string
var endPagePath string
var locationOverviewPath string
var wrongInputPath string
var runeForTokens string
var certFile string
var keyFile string

//Other Paths

func Init() {
	SetFlags("8443", "8444", 5*time.Minute, "assets/locations.xml")
}

func SetFlags(pPort1 string, pPort2 string, pTokenDur time.Duration, pXmlPath string) {
	flag.StringVar(&port1, "port1", pPort1, "4-digit number for the port regarding the qrCodeScan")
	flag.StringVar(&port2, "port2", pPort2, "4-digit number for the login Page")
	flag.DurationVar(&tokenDuration, "tokenDuration", pTokenDur, "The life duration of the token")
	flag.StringVar(&xmlPath, "xmlPath", pXmlPath, "Path to the xmlFile storing the locations")
}
func GetPath() string {
	return xmlPath
}
func GetTokenDuration() time.Duration {
	return tokenDuration
}
