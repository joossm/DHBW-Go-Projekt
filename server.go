package main

import (
	"GoPruefungsaufgabe/src/handler"
	_ "github.com/skip2/go-qrcode"
	_ "html/template"
	"log"
	"net/http"
	_ "time"
)

func main() {

	var serverMuxA = http.NewServeMux()
	var serverMuxB = http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./html/"))
	serverMuxA.Handle("/html/", http.StripPrefix("/html", fileServer))
	serverMuxB.Handle("/html/", http.StripPrefix("/html", fileServer))

	serverMuxB.HandleFunc("/qrSite", handler.QrCodeCreate)
	go func() {
		log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/Emil")
		http.ListenAndServeTLS(":8443", "server.crt", "server.key", serverMuxA)
	}()

	log.Printf("About to listen on 8444. Go to https://127.0.0.1:8444/qrSite")
	http.ListenAndServeTLS(":8444", "server.crt", "server.key", serverMuxB)

}
