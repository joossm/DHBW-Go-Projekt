package handler

import (
	writer "GoProjekt/src/log"
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
		var input = r.URL.String()
		var site = input[1 : len(input)-1]
		var websiteUrl = "https://127.0.0.1:8443/login?token="
		var websiteParameterLocation = "&location="
		var websiteToken = token.GetTokenByLocation(site)
		var websiteLocation = site
		var completeUrl = websiteUrl + websiteToken + websiteParameterLocation + websiteLocation
		log.Println(completeUrl)
		var png []byte
		png, err := qrcode.Encode(completeUrl, qrcode.Medium, 256)
		errorHandling(err)
		w.Write(png)
	}
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
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
		var name, address string
		for _, cookie := range r.Cookies() {
			if cookie.Name == "name" {
				name = cookie.Value
			}
			if cookie.Name == "address" {
				address = cookie.Value
			}
		}
		location := r.URL.Query().Get("location")
		cookieLocation := http.Cookie{Name: "location", Value: location}
		http.SetCookie(w, &cookieLocation)
		writer.WriteLoginToFile(name + ", " + address + ", " + location)
		cookieName := http.Cookie{Name: "name", Value: name}
		cookieAddress := http.Cookie{Name: "address", Value: address}
		http.SetCookie(w, &cookieName)
		http.SetCookie(w, &cookieAddress)
		http.Redirect(w, r, "/logout", 301)
	} else {
		if r.Method == "GET" {
			if r.URL.Query().Get("token") != "" {
				if token.ValidateTokenByLocation(r.URL.Query().Get("token"), r.URL.Query().Get("location")) == "true" {
					log.Println("Token Valid")
					location := r.URL.Query().Get("location")
					cookieLocation := http.Cookie{Name: "location", Value: location}
					http.SetCookie(w, &cookieLocation)
					t, err := template.ParseFiles("html/loginPage.html")
					errorHandling(err)
					err = t.Execute(w, nil)
					errorHandling(err)
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
			var location string
			for _, cookie := range r.Cookies() {
				if cookie.Name == "location" {
					location = cookie.Value
				}
			}
			writer.WriteLoginToFile(name + ", " + address + ", " + location)
			cookieName := http.Cookie{Name: "name", Value: name}
			cookieAddress := http.Cookie{Name: "address", Value: address}
			http.SetCookie(w, &cookieName)
			http.SetCookie(w, &cookieAddress)
			http.Redirect(w, r, "/logout", 301)
		}
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/logoutPage.html")
		errorHandling(err)
		err = t.Execute(w, nil)
		errorHandling(err)
	} else {
		var name string
		var address string
		var location string
		for _, cookie := range r.Cookies() {
			if cookie.Name == "name" {
				name = cookie.Value
			}
			if cookie.Name == "address" {
				address = cookie.Value
			}
			if cookie.Name == "location" {
				location = cookie.Value
			}
		}
		writer.WriteLogoutToFile(name + ", " + address + ", " + location)
		http.Redirect(w, r, "/end", 301)
	}
}

func End(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/endPage.html")
		errorHandling(err)
		err = t.Execute(w, nil)
		errorHandling(err)
	} else {
		http.Redirect(w, r, "/end", 301)
	}
}

func SelectLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		locations := model.ShowAllLocations()
		t, err := template.ParseFiles("html/selectLocation.html")
		errorHandling(err)
		err = t.Execute(w, locations)
		errorHandling(err)
	} else {
		http.Redirect(w, r, "/location", 301)
	}
}
