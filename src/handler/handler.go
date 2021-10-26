package handler

import (
	writer "GoPruefungsaufgabe/src/log"
	"GoPruefungsaufgabe/src/model"
	"GoPruefungsaufgabe/src/token"
	_ "GoPruefungsaufgabe/src/token"

	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"

	"log"
	"net/http"
)

func QrCodeCreate(res http.ResponseWriter, r *http.Request) {

	var websiteUrl = "https://127.0.0.1:8444/qrsite"
	var websiteParameter = "?token="
	var websiteToken = token.GetActiveToken()
	var completeUrl = websiteUrl + websiteParameter + websiteToken

	var png []byte
	png, err := qrcode.Encode(completeUrl, qrcode.Medium, 256)
	errorHandling(err)
	res.Write(png)

}

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	//creating new instance and checking method
	var alreadyAname bool
	var alreadyAaddress bool
	for _, cookie := range r.Cookies() {
		if cookie.Name == "name" {
			alreadyAname = true
		}
		if cookie.Name == "address" {
			alreadyAaddress = true
		}

		fmt.Println("Found a cookie named:", cookie.Name)
	}
	if alreadyAname == true && alreadyAaddress == true {
		var name, address string
		for _, cookie := range r.Cookies() {
			if cookie.Name == "name" {
				name = cookie.Value
			}
			if cookie.Name == "address" {
				address = cookie.Value
			}

			fmt.Println("Found a cookie named:", cookie.Name)
		}
		writer.WriteLoginToFile((name + ", " + address + "\n"))
		log.Println("Name: " + name + ", Address:" + address)

		cookieName := http.Cookie{Name: "name", Value: name}
		cookieAddress := http.Cookie{Name: "address", Value: address}

		http.SetCookie(w, &cookieName)
		http.SetCookie(w, &cookieAddress)

		http.Redirect(w, r, "/logout", 301)
	} else {
		if r.Method == "GET" {
			log.Println(r.URL.Query().Get("token"))
			if r.URL.Query().Get("token") != "" {
				if token.ValidateToken(r.URL.Query().Get("token")) == "true" {
					log.Println("Token Valid")
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
			writer.WriteLoginToFile((name + ", " + address + "\n"))
			log.Println("Name: " + name + ", Address:" + address)

			cookieName := http.Cookie{Name: "name", Value: name}
			cookieAddress := http.Cookie{Name: "address", Value: address}

			http.SetCookie(w, &cookieName)
			http.SetCookie(w, &cookieAddress)

			http.Redirect(w, r, "/logout", 301)
		}
	}

}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	//creating new instance and checking method
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/logoutPage.html")
		errorHandling(err)
		err = t.Execute(w, nil)
		errorHandling(err)

	} else {
		var name string
		var address string
		for _, cookie := range r.Cookies() {
			if cookie.Name == "name" {
				name = cookie.Value
			}
			if cookie.Name == "address" {
				address = cookie.Value
			}

			fmt.Println("Found a cookie named:", cookie.Name)
		}
		writer.WriteLogoutToFile((name + ", " + address + "\n"))
		log.Println("Name: " + name + ", Address:" + address)

		http.Redirect(w, r, "/end", 301)

	}

}

func End(w http.ResponseWriter, r *http.Request) {
	//creating new instance and checking method
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
