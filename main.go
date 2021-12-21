// 5807262
// 9899545
// 8622410

package main

import (
	"GoProjekt/src/analyzer"
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
	// TODO questioning the user if he wants to use the Analyzer or start the server
	if true {
		analyzer.AnlayzeReport()
	} else {
		// initialize both server
		var serverMuxA = http.NewServeMux()
		var serverMuxB = http.NewServeMux()

		// handle the locations from the QRCode
		locations := model.ReadXmlFile(config.GetPath()).ToStrings()
		for i := 0; i < len(locations); i++ {
			log.Println("Register of /" + locations[i])
			serverMuxB.HandleFunc("/"+locations[i], handler.QrCodeCreate)
		}

		// prepare for handle the static files
		fileServer := http.FileServer(http.Dir(flag.Lookup("fileServerPath").Value.String()))

		// initialize the serverMuxA with all the handlers
		serverMuxA.Handle(flag.Lookup("fileServerUrl").Value.String(), http.StripPrefix("/html", fileServer))
		serverMuxA.HandleFunc(flag.Lookup("endUrl").Value.String(), handler.ReLogin)
		serverMuxA.HandleFunc(flag.Lookup("loginUrl").Value.String(), handler.LoginUser)
		serverMuxA.HandleFunc(flag.Lookup("logoutUrl").Value.String(), handler.LogoutUser)
		serverMuxA.HandleFunc(flag.Lookup("locationUrl").Value.String(), handler.SelectLocation)

		// initialize the serverMuxB with the fileServer
		serverMuxB.Handle("/html/", http.StripPrefix("/html", fileServer))
		// TODO neccessary?
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
}
