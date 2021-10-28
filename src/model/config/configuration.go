package config

import (
	"flag"
	"time"
)

var port1 string
var port2 string
var tokenDuration time.Duration
var xmlPath string

var standardUrl string
var loginUrl string
var endUrl string
var locationUrl string
var logoutUrl string
var fileServerUrl string
var logfilePath string
var logoutPagePath string
var loginPagePath string
var endPagePath string
var locationOverviewPath string
var wrongInputPath string
var certFilePath string
var keyFilePath string
var fileServerPath string
var qrCodePath string

//Other Paths
func init(){
	SetPrimaryFlags("8443", "8444", 5*time.Minute)
	SetFlagsForPaths("src/log/files/", "html/logoutPage.html", "html/loginPage.html",
		"html/endPage.html", "html/locationOverview.html", "html/wrongInput.html",
		"server.crt", "server.key", "./html/", "html/qrCodes/", "assets/locations.xml")
	SetFlagsForLinks("/", "/login", "/logout", "/location", "/end", "/html/")
}

func Init() {
}
func InitForTest() {
	SetPrimaryFlags("8443", "8444", 5*time.Minute)
	//TODO change Paths for Tests
	SetFlagsForPaths("src/log/files/", "html/logoutPage.html", "html/loginPage.html",
		"html/endPage.html", "html/locationOverview.html", "html/wrongInput.html",
		"server.crt", "server.key", "./html/", "html/qrCodes/", "assets/locations.xml")
	SetFlagsForLinks("/", "/login", "/logout", "/location", "/end", "/html/")
}
func SetPrimaryFlags(pPort1 string, pPort2 string, pTokenDur time.Duration) {
	flag.StringVar(&port1, "port1", pPort1, "4-digit number for the port regarding the qrCodeScan")
	flag.StringVar(&port2, "port2", pPort2, "4-digit number for the login Page")
	flag.DurationVar(&tokenDuration, "tokenDuration", pTokenDur, "The life duration of the token")

}
func SetFlagsForPaths(pLogfilePath string, pLogoutPagePath string, pLoginPagePath string, pEndPagePath string,
	pLocationOverviewPath string, pWrongInputPath string, pCertFilePath string,
	pKeyFilePath string, pFileServerPath string, pQrCodePath string, pXmlPath string) {
	flag.StringVar(&logfilePath, "logfilePath", pLogfilePath, "Where the logs are stored")
	flag.StringVar(&logoutPagePath, "logoutPagePath", pLogoutPagePath, "Where the logout.html is stored")
	flag.StringVar(&loginPagePath, "loginPagePath", pLoginPagePath, "Where the login.html is stored")
	flag.StringVar(&endPagePath, "endPagePath", pEndPagePath, "Where the end.html is stored")
	flag.StringVar(&locationOverviewPath, "locationOverviewPath", pLocationOverviewPath, "Where the locationOverview.html is stored")
	flag.StringVar(&wrongInputPath, "wrongInputPath", pWrongInputPath, "Where the wrongInput.html is stored")
	flag.StringVar(&certFilePath, "certFilePath", pCertFilePath, "Where the certFile is stored")
	flag.StringVar(&keyFilePath, "keyFilePath", pKeyFilePath, "Where the keyFile is stored")
	flag.StringVar(&fileServerPath, "fileServerPath", pFileServerPath, "Where the fileServer is stored")
	flag.StringVar(&qrCodePath, "qrCodePath", pQrCodePath, "Where the fileServer is stored")
	flag.StringVar(&xmlPath, "xmlPath", pXmlPath, "Path to the xmlFile storing the locations")
}
func SetFlagsForLinks(pStandardUrl string, pLoginUrl string, pLogoutUrl string, pLocationUrl string,
	pEndUrl string, pFileServerUrl string) {
	flag.StringVar(&standardUrl, "standardUrl", pStandardUrl, "Where the login Page is accessible")
	flag.StringVar(&loginUrl, "loginUrl", pLoginUrl, "Where the login Page is accessible")
	flag.StringVar(&logoutUrl, "logoutUrl", pLogoutUrl, "Where the logout Page is accessible")
	flag.StringVar(&locationUrl, "locationUrl", pLocationUrl, "Where the location overview Page is accessible")
	flag.StringVar(&endUrl, "endUrl", pEndUrl, "Where the end Page is accessible")
	flag.StringVar(&fileServerUrl, "fileServerUrl", pFileServerUrl, "Where the end Page is accessible")
}
func GetPath() string {
	return xmlPath
}
func GetTokenDuration() time.Duration {
	return tokenDuration
}
