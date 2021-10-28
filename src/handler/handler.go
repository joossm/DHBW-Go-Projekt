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

func QrCodeCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println(urlBuilder(r))
		log.Println(getLocation(r))
		var filename = flag.Lookup("qrCodePath").Value.String() + getLocation(r) + ".png"
		_ = qrcode.WriteFile(urlBuilder(r), qrcode.Medium, 256, filename)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		AddForm := "<head><title>QR Code " + getLocation(r) + "</title><meta http-equiv=\"refresh\" content=\"5\"></head><body><div style=\"text-align: center;\"><br><br><br><br><br><br><br><br><h1>" + getLocation(r) + "</h1><br><br><br><br><img alt=\"" + urlBuilder(r) + "\" src=\"html/qrCodes/" + getLocation(r) + ".png\"></div><body>"
		fmt.Fprint(w, AddForm)
		return
	}
}

func parseAndExecuteWebsite(filename string, w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(filename)
	errorHandling(err)
	err = t.Execute(w, data)
	errorHandling(err)
}
func validateInputZipAndHouseNumber(w http.ResponseWriter, r *http.Request, forms ...string) (erg bool, str string) {
	for _, form := range forms {
		matchString, _ := regexp.MatchString("^[0-9]+$", r.FormValue(form))
		if matchString == false {
			return false, "Please use only upper and lower case letters for full name and address. Thank you."
		}

	}
	return true, ""
}
func validateInput(w http.ResponseWriter, r *http.Request, forms ...string) (erg bool, str string) {
	for _, form := range forms {
		matchString, _ := regexp.MatchString("[a-zA-z- ]+", r.FormValue(form))
		if matchString == false {
			return false, "Please use only upper and lower case letters for full name and address. Thank you." + form
		}

	}
	return true, ""
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) == true {
		var name = informationsFromCookies("name", r)
		var address = informationsFromCookies("address", r)
		var location = proofIfLoginInSameLocation(r)

		report.WriteToFile(true, combineText(name, address, location))

		setCookie(w, "name", name)
		setCookie(w, "address", address)
		setCookie(w, "location", location)

		http.Redirect(w, r, flag.Lookup("logoutUrl").Value.String(), 301)
	} else {

		if r.Method == "GET" {

			if r.URL.Query().Get("token") != "" {

				var tokenParameter = r.URL.Query().Get("token")
				var locationParameter = r.URL.Query().Get("location")

				if token.ValidateTokenByLocation(tokenParameter, locationParameter) == true {

					log.Println("Token Valid")
					setCookie(w, "location", r.URL.Query().Get("location"))

					parseAndExecuteWebsite(flag.Lookup("loginPagePath").Value.String(), w, nil)

				} else {

					log.Println("Token Invalid")
					http.Redirect(w, r, flag.Lookup("standardUrl").Value.String(), 403)
				}

			} else {

				log.Println("No Token")
				http.Redirect(w, r, flag.Lookup("standardUrl").Value.String(), 403)
			}

		} else {
			resBool, errStr := validateInputZipAndHouseNumber(w, r, "zipCode", "houseNumber")
			resBool, errStr = validateInput(w, r, "firstName", "lastName", "cityName", "streetName")
			if resBool == false {
				parseAndExecuteWebsite(flag.Lookup("wrongInputPath").Value.String(), w, errStr)

				return
			}
			name := r.FormValue("firstName") + " " + r.FormValue("lastName")
			address := r.FormValue("zipCode") + " " + r.FormValue("cityName") + " " + r.FormValue("streetName") + " " + r.FormValue("houseNumber")
			var location = informationsFromCookies("location", r)

			report.WriteToFile(true, combineText(name, address, location))

			setCookie(w, "name", name)
			setCookie(w, "address", address)

			http.Redirect(w, r, flag.Lookup("logoutUrl").Value.String(), 301)
		}
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parseAndExecuteWebsite(flag.Lookup("logoutPagePath").Value.String(), w, nil)
	} else {
		var name = informationsFromCookies("name", r)
		var address = informationsFromCookies("address", r)
		var location = informationsFromCookies("location", r)
		report.WriteToFile(false, combineText(name, address, location))
		http.Redirect(w, r, flag.Lookup("endUrl").Value.String(), 301)
	}
}

func End(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parseAndExecuteWebsite(flag.Lookup("endPagePath").Value.String(), w, nil)
	} else {
		http.Redirect(w, r, flag.Lookup("endUrl").Value.String(), 301)
	}
}

func SelectLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		locations := model.GetList().ShowAllLoc()
		parseAndExecuteWebsite(flag.Lookup("locationOverviewPath").Value.String(), w, locations)
	} else {
		http.Redirect(w, r, flag.Lookup("locationUrl").Value.String(), 301)
	}
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func alreadyLoggedIn(r *http.Request) bool {
	var loggedInWithName bool
	var loggedInWithAddress bool
	for _, cookie := range r.Cookies() {
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

func informationsFromCookies(value string, r *http.Request) string {
	for _, cookie := range r.Cookies() {
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

func proofIfLoginInSameLocation(r *http.Request) string {
	if r.URL.Query().Get("location") != "" {
		return r.URL.Query().Get("location")
	} else {
		return informationsFromCookies("location", r)
	}
}

func setCookie(w http.ResponseWriter, name string, value string) {
	cookieToStore := http.Cookie{Name: name, Value: value}
	http.SetCookie(w, &cookieToStore)
}

func urlBuilder(r *http.Request) string {

	return "https://127.0.0.1:8443/login?token=" + token.GetTokenByLocation(getLocation(r)) + "&location=" + getLocation(r)
}
func getLocation(r *http.Request) string {
	return r.URL.String()[1 : len(r.URL.String())-1]
}
