package handler

import (
	report "GoProjekt/src/log"
	"GoProjekt/src/model"
	"GoProjekt/src/token"
	"flag"
	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"
	"log"
	"net/http"
	"regexp"
	_ "strings"
)

func QrCodeCreate(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		log.Println(urlBuilder(request))
		log.Println(getLocation(request))
		var filename = flag.Lookup("qrCodePath").Value.String() + getLocation(request) + ".png"
		_ = qrcode.WriteFile(urlBuilder(request), qrcode.Medium, 256, filename)
		responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
		AddForm := "<head><title>QR Code " + getLocation(request) + "</title><meta http-equiv=\"refresh\" content=\"5\"></head><body><div style=\"text-align: center;\"><br><br><br><br><br><br><br><br><h1>" + getLocation(request) + "</h1><br><br><br><br><img alt=\"" + urlBuilder(request) + "\" src=\"html/qrCodes/" + getLocation(request) + ".png\"></div><body>"
		_, err := fmt.Fprint(responseWriter, AddForm)
		if err != nil {
			return
		}
		return
	}
}

func parseAndExecuteWebsite(filename string, responseWriter http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(filename)
	errorHandling(err)
	err = t.Execute(responseWriter, data)
	errorHandling(err)
}

func validateInputNumber(request *http.Request, forms ...string) (erg bool, str string) {
	for _, form := range forms {
		matchString, _ := regexp.MatchString("^[0-9]+$", request.FormValue(form))
		if matchString == false {
			return false, "Please use only numbers for zip code and house number. Thank you very much!"
		}

	}
	return true, ""
}
func validateInputLetter(request *http.Request, forms ...string) (erg bool, str string) {
	for _, form := range forms {
		matchString, _ := regexp.MatchString("[a-zA-z- ]+", request.FormValue(form))
		if matchString == false {
			return false, "Please use only upper and lower case letters for first and last name as well as city and street. Thank you very much!" + form
		}

	}
	return true, ""
}
func LoginUser(responseWriter http.ResponseWriter, request *http.Request) {
	if alreadyLoggedIn(request) == true {
		var name = informationsFromCookies("name", request)
		var address = informationsFromCookies("address", request)
		var location = proofIfLoginInSameLocation(request)

		report.WriteToFile(true, combineText(name, address, location))

		setCookie(responseWriter, "name", name)
		setCookie(responseWriter, "address", address)
		setCookie(responseWriter, "location", location)

		http.Redirect(responseWriter, request, flag.Lookup("logoutUrl").Value.String(), 301)
	} else {

		if request.Method == "GET" {

			if request.URL.Query().Get("token") != "" {

				var tokenParameter = request.URL.Query().Get("token")
				var locationParameter = request.URL.Query().Get("location")

				if token.ValidateTokenByLocation(tokenParameter, locationParameter) == true {

					log.Println("Token Valid")
					setCookie(responseWriter, "location", request.URL.Query().Get("location"))

					parseAndExecuteWebsite(flag.Lookup("loginPagePath").Value.String(), responseWriter, nil)

				} else {

					log.Println("Token Invalid")
					http.Redirect(responseWriter, request, flag.Lookup("standardUrl").Value.String(), 403)
				}

			} else {

				log.Println("No Token")
				http.Redirect(responseWriter, request, flag.Lookup("standardUrl").Value.String(), 403)
			}

		} else {
			resBool, errStr := validateInputNumber(request, "zipCode", "houseNumber")
			resBool, errStr = validateInputLetter(request, "firstName", "lastName", "cityName", "streetName")
			if resBool == false {
				parseAndExecuteWebsite(flag.Lookup("wrongInputPath").Value.String(), responseWriter, errStr)
				return
			}
			name := request.FormValue("firstName") + " " + request.FormValue("lastName")
			address := request.FormValue("zipCode") + " " + request.FormValue("cityName") + " " + request.FormValue("streetName") + " " + request.FormValue("houseNumber")
			var location = informationsFromCookies("location", request)

			report.WriteToFile(true, combineText(name, address, location))

			setCookie(responseWriter, "name", name)
			setCookie(responseWriter, "address", address)

			http.Redirect(responseWriter, request, flag.Lookup("logoutUrl").Value.String(), 301)
		}
	}
}

func LogoutUser(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		parseAndExecuteWebsite(flag.Lookup("logoutPagePath").Value.String(), responseWriter, nil)
	} else {
		var name = informationsFromCookies("name", request)
		var address = informationsFromCookies("address", request)
		var location = informationsFromCookies("location", request)
		report.WriteToFile(false, combineText(name, address, location))
		http.Redirect(responseWriter, request, flag.Lookup("endUrl").Value.String(), 301)
	}
}

func End(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		parseAndExecuteWebsite(flag.Lookup("endPagePath").Value.String(), responseWriter, nil)
	} else {
		http.Redirect(responseWriter, request, flag.Lookup("endUrl").Value.String(), 301)
	}
}

func SelectLocation(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		locations := model.GetList().ShowAllLoc()
		parseAndExecuteWebsite(flag.Lookup("locationOverviewPath").Value.String(), responseWriter, locations)
	} else {
		http.Redirect(responseWriter, request, flag.Lookup("locationUrl").Value.String(), 301)
	}
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func alreadyLoggedIn(request *http.Request) bool {
	var loggedInWithName bool
	var loggedInWithAddress bool
	for _, cookie := range request.Cookies() {
		if cookie.Name == "name" {
			loggedInWithName = true
		}
		if cookie.Name == "address" {
			loggedInWithAddress = true
		}
	}
	if loggedInWithName == true && loggedInWithAddress == true {
		return true
	} else {
		return false
	}
}

func informationsFromCookies(value string, request *http.Request) string {
	for _, cookie := range request.Cookies() {
		if cookie.Name == "name" && value == "name" {
			return cookie.Value
		}
		if cookie.Name == "address" && value == "address" {
			return cookie.Value
		}
		if cookie.Name == "location" && value == "location" {
			return cookie.Value
		}
	}
	return "NO INFORMATION"
}

func combineText(name string, address string, location string) string {
	return name + ", " + address + ", " + location
}

func proofIfLoginInSameLocation(request *http.Request) string {
	if request.URL.Query().Get("location") != "" {
		return request.URL.Query().Get("location")
	} else {
		return informationsFromCookies("location", request)
	}
}

func setCookie(responseWriter http.ResponseWriter, name string, value string) {
	cookieToStore := http.Cookie{Name: name, Value: value}
	http.SetCookie(responseWriter, &cookieToStore)
}

func urlBuilder(request *http.Request) string {

	return "https://127.0.0.1:8443/login?token=" + token.GetTokenByLocation(getLocation(request)) + "&location=" + getLocation(request)
}
func getLocation(request *http.Request) string {
	return request.URL.String()[1 : len(request.URL.String())-1]
}
