// 5807262
// 9899545
// 8622410

package config

import (
	"flag"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	if port1 != "8443" {
		t.Error("port 1 is ", port1, "but it should be :8443")
	}
	if port2 != "8444" {
		t.Error("port 2 is ", port2, "but it should be :8444")
	}
	if tokenDuration != time.Minute*5 {
		t.Error("token duration is ", tokenDuration, "but it should be 60")
	}
	if xmlPath == "" {
		t.Error("path is empty")
	}
	//check if Flag is parsed
	assert.Equal(t, flag.Parsed(), true)
}

//check primary Flags != nil
func TestPrimaryFlagsNotNil(t *testing.T) {
	assert.False(t, flag.Lookup("port1") == nil)
	assert.False(t, flag.Lookup("port2") == nil)
	assert.False(t, flag.Lookup("tokenDuration") == nil)
}

//check Path Flags != nil
func TestPathFlagsNotNil(t *testing.T) {
	assert.False(t, flag.Lookup("logfilePath") == nil)
	assert.False(t, flag.Lookup("logoutPagePath") == nil)
	assert.False(t, flag.Lookup("loginPagePath") == nil)
	assert.False(t, flag.Lookup("endPagePath") == nil)
	assert.False(t, flag.Lookup("locationOverviewPath") == nil)
	assert.False(t, flag.Lookup("wrongInputPath") == nil)
	assert.False(t, flag.Lookup("certFilePath") == nil)
	assert.False(t, flag.Lookup("keyFilePath") == nil)
	assert.False(t, flag.Lookup("fileServerPath") == nil)
	assert.False(t, flag.Lookup("qrCodePath") == nil)
	assert.False(t, flag.Lookup("xmlPath") == nil)
}

//check Link Flags != nil
func TestLinkFlagsNotNil(t *testing.T) {
	assert.False(t, flag.Lookup("standardUrl") == nil)
	assert.False(t, flag.Lookup("loginUrl") == nil)
	assert.False(t, flag.Lookup("logoutUrl") == nil)
	assert.False(t, flag.Lookup("locationUrl") == nil)
	assert.False(t, flag.Lookup("endUrl") == nil)
	assert.False(t, flag.Lookup("fileServerUrl") == nil)
}

//check that flag has the same values as config (check if the flag is reachable)
//primary Flags
func TestPrimaryFlagsReachable(t *testing.T) {
	assert.Equal(t, port1, flag.Lookup("port1").Value.String(), true)
	assert.Equal(t, port2, flag.Lookup("port2").Value.String(), true)
	assert.Equal(t, tokenDuration.String(), flag.Lookup("tokenDuration").Value.String(), true)

}

//path Flags
func TestPathFlagsReachable(t *testing.T) {
	assert.Equal(t, logfilePath, flag.Lookup("logfilePath").Value.String(), true)
	assert.Equal(t, logoutPagePath, flag.Lookup("logoutPagePath").Value.String(), true)
	assert.Equal(t, loginPagePath, flag.Lookup("loginPagePath").Value.String(), true)
	assert.Equal(t, endPagePath, flag.Lookup("endPagePath").Value.String(), true)
	assert.Equal(t, locationOverviewPath, flag.Lookup("locationOverviewPath").Value.String(), true)
	assert.Equal(t, wrongInputPath, flag.Lookup("wrongInputPath").Value.String(), true)
	assert.Equal(t, certFilePath, flag.Lookup("certFilePath").Value.String(), true)
	assert.Equal(t, keyFilePath, flag.Lookup("keyFilePath").Value.String(), true)
	assert.Equal(t, fileServerPath, flag.Lookup("fileServerPath").Value.String(), true)
	assert.Equal(t, qrCodePath, flag.Lookup("qrCodePath").Value.String(), true)
	assert.Equal(t, xmlPath, flag.Lookup("xmlPath").Value.String(), true)
}

//Link Flags
func TestLinkFlagsReachable(t *testing.T) {
	assert.Equal(t, standardUrl, flag.Lookup("standardUrl").Value.String(), true)
	assert.Equal(t, loginUrl, flag.Lookup("loginUrl").Value.String(), true)
	assert.Equal(t, logoutUrl, flag.Lookup("logoutUrl").Value.String(), true)
	assert.Equal(t, locationUrl, flag.Lookup("locationUrl").Value.String(), true)
	assert.Equal(t, endUrl, flag.Lookup("endUrl").Value.String(), true)
	assert.Equal(t, fileServerUrl, flag.Lookup("fileServerUrl").Value.String(), true)
}

func TestInitForTesting(t *testing.T) {
	assert.False(t, flag.Lookup("endPagePath").Value.String() == "../../view/reLoginPage.html")
	assert.False(t, flag.Lookup("loginPagePath").Value.String() == "../../view/loginPage.html")
	assert.False(t, flag.Lookup("logoutPagePath").Value.String() == "../../view/logoutPage.html")
	InitForTesting()
	assert.True(t, flag.Lookup("endPagePath").Value.String() == "../../view/reLoginPage.html")
	assert.True(t, flag.Lookup("loginPagePath").Value.String() == "../../view/loginPage.html")
	assert.True(t, flag.Lookup("logoutPagePath").Value.String() == "../../view/logoutPage.html")

}

func TestGetPath(t *testing.T) {
	assert.True(t, flag.Lookup("xmlPath").Value.String() == GetPath())
}
func TestGetTokenDuration(t *testing.T) {
	assert.True(t, flag.Lookup("tokenDuration").Value.String() == GetTokenDuration().String())
}
