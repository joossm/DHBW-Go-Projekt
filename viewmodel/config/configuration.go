// 5807262
// 9899545
// 8622410

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
func init() {
	SetPrimaryFlags("8443", "8444", 5*time.Minute)

	SetFlagsForPaths("model/log/files/", "view/logoutPage.html", "view/loginPage.html",
		"view/reLoginPage.html", "view/locationOverview.html", "view/wrongInput.html",
		"server.crt", "server.key", "./view/", "view/qrCodes/", "model/assets/locations.xml")

	SetFlagsForLinks("/", "/login", "/logout", "/location", "/end", "/view/")
	flag.Parse()
}

//SetPrimaryFlags sets the flags for both servers and the token duration
func SetPrimaryFlags(pPort1 string, pPort2 string, pTokenDur time.Duration) {
	flag.StringVar(&port1, "port1", pPort1, "4-digit number for the port regarding the qrCodeScan")
	flag.StringVar(&port2, "port2", pPort2, "4-digit number for the login Page")
	flag.DurationVar(&tokenDuration, "tokenDuration", pTokenDur, "The life duration of the token")

}

// SetFlagsForPaths sets the flags for the paths
func SetFlagsForPaths(pLogfilePath string, pLogoutPagePath string, pLoginPagePath string, pEndPagePath string,
	pLocationOverviewPath string, pWrongInputPath string, pCertFilePath string,
	pKeyFilePath string, pFileServerPath string, pQrCodePath string, pXmlPath string) {
	flag.StringVar(&logfilePath, "logfilePath", pLogfilePath, "Where the logs are stored")
	flag.StringVar(&logoutPagePath, "logoutPagePath", pLogoutPagePath, "Where the logout.view is stored")
	flag.StringVar(&loginPagePath, "loginPagePath", pLoginPagePath, "Where the login.view is stored")
	flag.StringVar(&endPagePath, "endPagePath", pEndPagePath, "Where the end.view is stored")
	flag.StringVar(&locationOverviewPath, "locationOverviewPath", pLocationOverviewPath, "Where the locationOverview.view is stored")
	flag.StringVar(&wrongInputPath, "wrongInputPath", pWrongInputPath, "Where the wrongInput.view is stored")
	flag.StringVar(&certFilePath, "certFilePath", pCertFilePath, "Where the certFile is stored")
	flag.StringVar(&keyFilePath, "keyFilePath", pKeyFilePath, "Where the keyFile is stored")
	flag.StringVar(&fileServerPath, "fileServerPath", pFileServerPath, "Where the fileServer is stored")
	flag.StringVar(&qrCodePath, "qrCodePath", pQrCodePath, "Where the fileServer is stored")
	flag.StringVar(&xmlPath, "xmlPath", pXmlPath, "Path to the xmlFile storing the locations")
}

// SetFlagsForLinks sets the flags for the links in the server
func SetFlagsForLinks(pStandardUrl string, pLoginUrl string, pLogoutUrl string, pLocationUrl string,
	pEndUrl string, pFileServerUrl string) {
	flag.StringVar(&standardUrl, "standardUrl", pStandardUrl, "Where the login Page is accessible")
	flag.StringVar(&loginUrl, "loginUrl", pLoginUrl, "Where the login Page is accessible")
	flag.StringVar(&logoutUrl, "logoutUrl", pLogoutUrl, "Where the logout Page is accessible")
	flag.StringVar(&locationUrl, "locationUrl", pLocationUrl, "Where the location overview Page is accessible")
	flag.StringVar(&endUrl, "endUrl", pEndUrl, "Where the end Page is accessible")
	flag.StringVar(&fileServerUrl, "fileServerUrl", pFileServerUrl, "Where the end Page is accessible")
}

// GetPath returns the path
func GetPath() string {
	return xmlPath
}

// GetTokenDuration returns the tokenDuration
func GetTokenDuration() time.Duration {
	return tokenDuration
}

// InitForTesting initializes the flags for testing, because the file locations need to be addressed differently
func InitForTesting() {
	err := flag.Set("endPagePath", "../../view/reLoginPage.html")
	if err != nil {
		return
	}
	err = flag.Set("loginPagePath", "../../view/loginPage.html")
	if err != nil {
		return
	}
	err = flag.Set("logoutPagePath", "../../view/logoutPage.html")
	if err != nil {
		return
	}
	err = flag.Set("logfilePath", "../../model/log/files/")
	if err != nil {
		return
	}
	err = flag.Set("locationOverviewPath", "../../view/locationOverview.html")
	if err != nil {
		return
	}
	err = flag.Set("wrongInputPath", "../../view/wrongInput.html")
	if err != nil {
		return
	}
	err = flag.Set("certFilePath", "../../server.crt")
	if err != nil {
		return
	}
	err = flag.Set("xmlPath", "assets/locations.xml")
	if err != nil {
		return
	}
}
