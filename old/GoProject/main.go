package main

/*//func model() {
	fileServer := http.FileServer(http.Dir("GoProject"))
	http.Handle("/", fileServer)
	http.HandleFunc("/addnewuser/", handlers.AddNewUserFunc)
	http.HandleFunc("/notsucceded", handlers.NotSucceded)

	http.HandleFunc("/deleted", handlers.DeletedFunc)
	http.HandleFunc("/deleteuser/deleted", handlers.DeleteUserFunc)
	http.HandleFunc("/deleteuser/", handlers.DeleteUserServe)
	http.HandleFunc("/deleteuser/notsuccededdelete", handlers.NotSuccededDelete)

	http.HandleFunc("/", handlers.IndexFunc)

	http.HandleFunc("/showuser/show", handlers.ShowUserFunc)
	http.HandleFunc("/showuser/", handlers.ShowUser)
	http.HandleFunc("/showuser/notsuccededshow/", handlers.NotSuccededShow)

	http.ListenAndServe(":8080", nil)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}*/
