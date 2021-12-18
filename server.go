package main

import (
	"GoProjekt/src/handler"
	"GoProjekt/src/model"
	"GoProjekt/src/model/config"
	"GoProjekt/src/token"
	"flag"
	_ "github.com/skip2/go-qrcode"
	_ "html/template"
	"log"
	"net/http"
	"time"
	_ "time"
)

func main() {
	//cmd.AnlayzeReport()
	var serverMuxA = http.NewServeMux()
	var serverMuxB = http.NewServeMux()

	locations := model.ReadXmlFile(config.GetPath()).ToStrings()
	for i := 0; i < len(locations); i++ {
		log.Println("Register of /" + locations[i])
		serverMuxB.HandleFunc("/"+locations[i], handler.QrCodeCreate)
	}

	fileServer := http.FileServer(http.Dir(flag.Lookup("fileServerPath").Value.String()))

	serverMuxA.Handle(flag.Lookup("fileServerUrl").Value.String(), http.StripPrefix("/html", fileServer))
	serverMuxA.HandleFunc(flag.Lookup("endUrl").Value.String(), handler.End)
	serverMuxA.HandleFunc(flag.Lookup("loginUrl").Value.String(), handler.LoginUser)
	serverMuxA.HandleFunc(flag.Lookup("logoutUrl").Value.String(), handler.LogoutUser)
	serverMuxA.HandleFunc(flag.Lookup("locationUrl").Value.String(), handler.SelectLocation)

	serverMuxB.Handle("/html/", http.StripPrefix("/html", fileServer))
	//serverMuxB.HandleFunc("/qr", handler.QrCodeCreate)
	go func() {
		token.CreateAndUpdateTokenMap(locations)
		for now := range time.Tick(config.GetTokenDuration()) {
			token.CreateAndUpdateTokenMap(locations)
			log.Println(now, "Token Updated")
		}
	}()

	go func() {
		log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/location")
		err := http.ListenAndServeTLS(":"+flag.Lookup("port1").Value.String(),
			flag.Lookup("certFilePath").Value.String(), flag.Lookup("keyFilePath").Value.String(), serverMuxA)
		if err != nil {
			return
		} // port1 added needs to be tested
	}()

	log.Printf("About to listen on 8444.")
	err := http.ListenAndServeTLS(":"+flag.Lookup("port2").Value.String(),
		flag.Lookup("certFilePath").Value.String(), flag.Lookup("keyFilePath").Value.String(), serverMuxB)
	if err != nil {
		return
	}
}
