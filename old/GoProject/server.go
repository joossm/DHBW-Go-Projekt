package main

import (
	"GoProject/handlers"
	"github.com/skip2/go-qrcode"
	_ "html/template"
	"log"
	"net/http"
	_ "time"
)

var old = "/addnewuser/"

func main() {
	/*var globalSessions *session.Manager
	// Then, initialize the session manager
	func
	init()
	{
		globalSessions = NewManager("memory", "gosessionid", 3600)
	}
	func(manager *Manager) sessionId()
	string{
		b := make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, b); err != nil{
		return ""
	}
		return base64.URLEncoding.EncodeToString(b)
	}*/

	var serverMuxA = http.NewServeMux()
	var serverMuxB = http.NewServeMux()
	content := "http://localhost:8080/addnewuser/?"
	fileLocation := "templates/qrCodes/qr2.png"
	err := qrcode.WriteFile(content, qrcode.Medium, 256, fileLocation)
	if err != nil {
	}
	fileServer := http.FileServer(http.Dir("./templates/qrCodes/"))
	serverMuxA.Handle("/qrCodes/", http.StripPrefix("/qrCodes", fileServer))
	serverMuxB.Handle("/qrCodes/", http.StripPrefix("/qrCodes", fileServer))

	serverMuxA.HandleFunc("/addnewuser/", handlers.AddNewUserFunc)
	serverMuxA.HandleFunc("/notsucceded", handlers.NotSucceded)

	serverMuxA.HandleFunc("/deleted", handlers.DeletedFunc)
	serverMuxA.HandleFunc("/deleteuser/deleted", handlers.DeleteUserFunc)
	serverMuxA.HandleFunc("/deleteuser/", handlers.DeleteUserServe)
	serverMuxA.HandleFunc("/deleteuser/notsuccededdelete", handlers.NotSuccededDelete)

	serverMuxA.HandleFunc("/start", handlers.IndexFunc)

	serverMuxA.HandleFunc("/showuser/show", handlers.ShowUserFunc)
	serverMuxA.HandleFunc("/showuser/", handlers.ShowUser)
	serverMuxA.HandleFunc("/showuser/notsuccededshow/", handlers.NotSuccededShow)
	//fmt.Printf("Starting server at port 8080\n")

	serverMuxB.HandleFunc("/qrSite", handlers.QrCodeCreate)

	//http.ListenAndServe(":8080", nil)

	//err = http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	//log.Fatal(err)
	/*serverMuxA.HandleFunc("/secret", session.Secret)
	serverMuxA.HandleFunc("/login", session.Login)
	serverMuxA.HandleFunc("/logout", session.Logout)
	serverMuxB.HandleFunc("/secret", session.Secret)
	serverMuxB.HandleFunc("/login", session.Login)
	serverMuxB.HandleFunc("/logout", session.Logout)*/

	go func() {
		log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/start")
		http.ListenAndServeTLS(":8443", "server.crt", "server.key", serverMuxA)
		/*go func() {
			var i = 0

			for range time.Tick(time.Minute * 1) {
				http.
				salt := string(rand.Int())
				serverMuxA
				link := ("/addnewuser/" + salt)
				serverMuxA.HandleFunc(link, handlers.AddNewUserFunc)
				old = link
			}
		}()*/
	}()

	log.Printf("About to listen on 8444. Go to https://127.0.0.1:8444/qrSite")
	http.ListenAndServeTLS(":8444", "server.crt", "server.key", serverMuxB)

	//finish := make(chan bool)
	//<-finish

}
