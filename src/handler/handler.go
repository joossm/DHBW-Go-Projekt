package handler

import (
	report "GoProjekt/src/log"
	"GoProjekt/src/model"
	"GoProjekt/src/token"
	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"
	"log"
	"net/http"
)

func QrCodeCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println(urlBuilder(r))
		var png []byte
		png, err := qrcode.Encode(urlBuilder(r), qrcode.Medium, 256)
		errorHandling(err)
		_, err = w.Write(png)
		errorHandling(err)
	}
}

func parseAndExecuteWebsite(filename string, w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(filename)
	errorHandling(err)
	err = t.Execute(w, data)
	errorHandling(err)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) == true {
		var name = informationsFromCookies("name", r)
		var address = informationsFromCookies("address", r)
		var location = proofIfLoginInSameLocation(r)

		report.WriteLoginToFile(combineText(name, address, location))

		setCookie(w, "name", name)
		setCookie(w, "address", address)
		setCookie(w, "location", location)

		http.Redirect(w, r, "/logout", 301)
	} else {

		if r.Method == "GET" {

			if r.URL.Query().Get("token") != "" {

				var tokenParameter = r.URL.Query().Get("token")
				var locationParameter = r.URL.Query().Get("location")

				if token.ValidateTokenByLocation(tokenParameter, locationParameter) == "true" {

					log.Println("Token Valid")
					setCookie(w, "location", r.URL.Query().Get("location"))

					parseAndExecuteWebsite("html/loginPage.html", w, nil)

				} else {

					log.Println("Token Invalid")
					http.Redirect(w, r, "/", 403)
				}

			} else {

				log.Println("No Token")
				http.Redirect(w, r, "/", 403)
			}

		} else {
			name := r.FormValue("name")
			address := r.FormValue("address")
			var location = informationsFromCookies("location", r)

			report.WriteLoginToFile(combineText(name, address, location))

			setCookie(w, "name", name)
			setCookie(w, "address", address)

			http.Redirect(w, r, "/logout", 301)
		}
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parseAndExecuteWebsite("html/logoutPage.html", w, nil)
	} else {
		var name = informationsFromCookies("name", r)
		var address = informationsFromCookies("address", r)
		var location = informationsFromCookies("location", r)

		report.WriteLogoutToFile(name + ", " + address + ", " + location)
		http.Redirect(w, r, "/end", 301)
	}
}

func End(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parseAndExecuteWebsite("html/endPage.html", w, nil)
	} else {
		http.Redirect(w, r, "/end", 301)
	}
}

func SelectLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		locations := model.ShowAllLocations()
		parseAndExecuteWebsite("html/locationOverview.html", w, locations)
	} else {
		http.Redirect(w, r, "/location", 301)
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
	var input = r.URL.String()
	return "https://127.0.0.1:8443/login?token=" + token.GetTokenByLocation(input[1:len(input)-1]) + "&location=" + (input[1 : len(input)-1])
}
