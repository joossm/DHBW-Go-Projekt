package main

import (
	"GoPruefungsaufgabe/src/handler"
	"GoPruefungsaufgabe/src/main/config"
	"flag"
	_ "github.com/skip2/go-qrcode"
	_ "html/template"
	"log"
	"net/http"
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

	serverMuxB.HandleFunc("/qrSite", handler.QrCodeCreate)
	go func() {
		log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/Emil")
		http.ListenAndServeTLS(":" + port1, "server.crt", "server.key", serverMuxA) // port1 added needs to be tested
	}()

	log.Printf("About to listen on 8444. Go to https://127.0.0.1:8444/qrSite")
	http.ListenAndServeTLS(":" + port2, "server.crt", "server.key", serverMuxB)

}
