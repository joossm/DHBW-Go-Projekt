// 5807262
// 9899545
// 8622410

package main

import (
	"GoProjekt/model"
	"GoProjekt/viewmodel/analyzer"
	"GoProjekt/viewmodel/config"
	"GoProjekt/viewmodel/handler"
	"GoProjekt/viewmodel/token"
	"bufio"
	"flag"
	_ "github.com/skip2/go-qrcode"
	_ "html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
	_ "time"
)

func main() {
	// read command line arguments
	println("Please choose a mode:\n 1. Analyzer\n 2. Server\n")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		text = strings.Replace(text, "\r\n", "", -1)
	}
	if runtime.GOOS == "linux" {
		text = strings.Replace(text, "\n", "", -1)
	}
	// TODO questioning the user if he wants to use the Analyzer or start the server
	if text == "1" {
		analyzer.AnlayzeReport()
	} else if text == "2" {
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
		serverMuxA.Handle(flag.Lookup("fileServerUrl").Value.String(), http.StripPrefix("/view", fileServer))
		serverMuxA.HandleFunc(flag.Lookup("endUrl").Value.String(), handler.ReLogin)
		serverMuxA.HandleFunc(flag.Lookup("loginUrl").Value.String(), handler.LoginUser)
		serverMuxA.HandleFunc(flag.Lookup("logoutUrl").Value.String(), handler.LogoutUser)
		serverMuxA.HandleFunc(flag.Lookup("locationUrl").Value.String(), handler.SelectLocation)

		// initialize the serverMuxB with the fileServer
		serverMuxB.Handle("/view/", http.StripPrefix("/view", fileServer))
		// TODO neccessary?
		//serverMuxB.HandleFunc("/qr", handler.QrCodeCreate)

		// start the TokenUpdater
		go func() {
			token.CreateAndUpdateTokenMap(locations)
			for now := range time.Tick(config.GetTokenDuration()) {
				token.CreateAndUpdateTokenMap(locations)
				log.Println(now, "Token Updated")
			}
		}()

		// start the server with the serverMuxA
		go func() {
			log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/location")
			err := http.ListenAndServeTLS(":"+flag.Lookup("port1").Value.String(),
				flag.Lookup("certFilePath").Value.String(), flag.Lookup("keyFilePath").Value.String(), serverMuxA)
			if err != nil {
				return
			} // port1 added needs to be tested
		}()

		// start the server with the serverMuxB
		log.Printf("About to listen on 8444.")
		err := http.ListenAndServeTLS(":"+flag.Lookup("port2").Value.String(),
			flag.Lookup("certFilePath").Value.String(), flag.Lookup("keyFilePath").Value.String(), serverMuxB)
		if err != nil {
			return
		}
	}
}
