package main

import (
	"GoProjekt/src/handler"
	"GoProjekt/src/model"
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
	//cmd.Lauft()

	config.Init()
	var port1 = flag.Lookup("port1").Value.String()
	var port2 = flag.Lookup("port2").Value.String()
	var serverMuxA = http.NewServeMux()
	var serverMuxB = http.NewServeMux()

	locations := model.RegisterLocations().ToStrings()
	for i := 0; i < len(locations); i++ {
		log.Println("Register of /" + locations[i])
		serverMuxB.HandleFunc("/"+locations[i], handler.QrCodeCreate)
	}

	fileServer := http.FileServer(http.Dir("./html/"))
	serverMuxA.Handle("/html/", http.StripPrefix("/html", fileServer))
	serverMuxB.Handle("/html/", http.StripPrefix("/html", fileServer))
	serverMuxA.HandleFunc("/end", handler.End)
	serverMuxA.HandleFunc("/login", handler.LoginUser)
	serverMuxA.HandleFunc("/logout", handler.LogoutUser)
	serverMuxA.HandleFunc("/location", handler.SelectLocation)
	serverMuxB.HandleFunc("/qr", handler.QrCodeCreate)
	go func() {
		token.CreateAndUpdateTokenMap(locations)

		for now := range time.Tick(30 * time.Minute) {
			token.CreateAndUpdateTokenMap(locations)
			log.Println(now, "Token Updated")
		}
	}()

	go func() {
		log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/login?token=FPLLNGZIEYOH")
		err := http.ListenAndServeTLS(":"+port1, "server.crt", "server.key", serverMuxA)
		if err != nil {
			return
		} // port1 added needs to be tested
	}()

	log.Printf("About to listen on 8444. Go to https://127.0.0.1:8444/qr")
	err := http.ListenAndServeTLS(":"+port2, "server.crt", "server.key", serverMuxB)
	if err != nil {
		return
	}

}
