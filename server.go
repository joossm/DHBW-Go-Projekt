package main

import (
	"GoProjekt/src/handler"
	"GoProjekt/src/model/config"
	token "GoProjekt/src/token"
	"flag"
	_ "github.com/skip2/go-qrcode"
	_ "html/template"
	"log"
	"net/http"
	"time"
	_ "time"
)

func main() {
	config.Init()
	var port1 = flag.Lookup("port1").Value.String()
	var port2 = flag.Lookup("port2").Value.String()
	var serverMuxA = http.NewServeMux()
	var serverMuxB = http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./html/"))
	serverMuxA.Handle("/html/", http.StripPrefix("/html", fileServer))
	serverMuxB.Handle("/html/", http.StripPrefix("/html", fileServer))

	serverMuxA.HandleFunc("/end", handler.End)
	serverMuxA.HandleFunc("/login", handler.LoginUser)
	serverMuxA.HandleFunc("/logout", handler.LogoutUser)
	serverMuxA.HandleFunc("/token", token.CreateToken)
	serverMuxA.HandleFunc("/location", handler.SelectLocation)
	serverMuxB.HandleFunc("/qrSite", handler.QrCodeCreate)
	go func() {
		var websiteToken string
		websiteToken = token.CreateAndUpdateToken()
		if token.ValidateToken(websiteToken) == "true" {
			log.Println("Token successful Updated " + websiteToken)
		} else {
			log.Println("Token Updated not successful")
		}
		for now := range time.Tick(5 * time.Minute) {
			log.Println(now, "Token Update started")

			websiteToken = token.CreateAndUpdateToken()
			if token.ValidateToken(websiteToken) == "true" {
				log.Println(now, "Token successful Updated")
			} else {
				log.Println(now, "Token Updated not successful")
			}
		}
	}()

	go func() {
		log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/login?token=FPLLNGZIEYOH")
		http.ListenAndServeTLS(":"+port1, "server.crt", "server.key", serverMuxA) // port1 added needs to be tested
	}()

	log.Printf("About to listen on 8444. Go to https://127.0.0.1:8444/qrSite")
	http.ListenAndServeTLS(":"+port2, "server.crt", "server.key", serverMuxB)

}
