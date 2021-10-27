package handler

import (
	writer "GoProjekt/src/log"
	"GoProjekt/src/model"
	"GoProjekt/src/token"
	_ "GoProjekt/src/token"

	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"

	"log"
	"net/http"
)

func QrCodeCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println(r.URL.Query().Get("location"))
		log.Println(r.URL)
		var input string = string(r.URL.String())
		var site string = input[1 : len(input)-1]
		log.Println(site)

		/*t, err := template.ParseFiles("html/qrPage.html")
		errorHandling(err)
		err = t.Execute(w, nil)
		errorHandling(err)*/

		log.Println(r.URL.Query().Get("location"))
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
		location := r.URL.Query().Get("location")
		log.Println("Location ist:::::" + location)
		cookieLocation := http.Cookie{Name: "location", Value: location}
		http.SetCookie(w, &cookieLocation)
		writer.WriteLoginToFile(name + ", " + address + ", " + location)
		log.Println("Name: " + name + ", Address:" + address + ", Location:" + location)

		cookieName := http.Cookie{Name: "name", Value: name}
		cookieAddress := http.Cookie{Name: "address", Value: address}

		http.SetCookie(w, &cookieName)
		http.SetCookie(w, &cookieAddress)

		http.Redirect(w, r, "/logout", 301)
	} else {
		if r.Method == "GET" {
			log.Println(r.URL.Query().Get("token"))
			if r.URL.Query().Get("token") != "" {
				if token.ValidateTokenByLocation(r.URL.Query().Get("token"), r.URL.Query().Get("location")) == "true" {
					log.Println("Token Valid")
					location := r.URL.Query().Get("location")
					log.Println("Location ist:::::" + location)
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
			log.Println(r.URL)
			for _, cookie := range r.Cookies() {
				if cookie.Name == "location" {
					location = cookie.Value
				}

				fmt.Println("Found a cookie named:", cookie.Name)
			}

			writer.WriteLoginToFile(name + ", " + address + ", " + location)
			log.Println("Name: " + name + ", Address:" + address + ", Location:" + location)

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

			fmt.Println("Found a cookie named:", cookie.Name)
		}
		writer.WriteLogoutToFile(name + ", " + address + ", " + location)
		log.Println("Name: " + name + ", Address:" + address + ", Location:" + location)

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
