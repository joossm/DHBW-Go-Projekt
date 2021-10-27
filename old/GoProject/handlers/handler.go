package handlers

import (
	"GoProject/model"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//function to check correct user adding input (regular expression and non-empty field input)
func checkFormValue(w http.ResponseWriter, r *http.Request, forms ...string) (res bool, errStr string) {
	for _, form := range forms {
		m, _ := regexp.MatchString("^[a-zA-Z]+$", r.FormValue(form))
		if r.FormValue(form) == "" {
			return false, "All forms must be completed"
		}
		if m == false {
			return false, "Use only english letters if firstname,lastname forms"
		}

	}
	return true, ""
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/showUser.html")
}

func NotSuccededShow(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/notSuccededShow.html")

}

//handler to show user with id input
func ShowUserFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/showUserPage.html")
		t.Execute(w, nil)

	} else {

		id, err := strconv.Atoi(r.FormValue("id"))
		checkError(err)
		var allUsers model.AllUsers
		file, err := os.OpenFile("list.json", os.O_RDONLY, 0666)
		checkError(err)
		b, err := ioutil.ReadAll(file)
		checkError(err)
		json.Unmarshal(b, &allUsers.Users)

		var allID []int
		for _, usr := range allUsers.Users {
			allID = append(allID, usr.Id)
		}
		for _, usr := range allUsers.Users {
			if model.IsValueInSlice(allID, id) != true {
				http.Redirect(w, r, "/showuser/notsuccededshow/", 302)
				return
			}
			if usr.Id != id {
				continue
			} else {
				t, err := template.ParseFiles("templates/showUserPage.html")
				checkError(err)
				t.Execute(w, usr)
			}

		}
	}
}

//function to handle page with successful deletion
func DeletedFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/deleted.html")
}

//serving file with error (add function:empty field input or uncorrect input)
func NotSucceded(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/notSucceded.html")
}

//function,which serve html file,when deleting was not succesful(id input is not correct)
func NotSuccededDelete(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/notSuccededDelete.html")
}

//function,which serve page with delete information input
func DeleteUserServe(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/deleteUser.html")

}

//function to delete user
func DeleteUserFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/deleteUser.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		id, err := strconv.Atoi(r.FormValue("id"))
		checkError(err)

		//open file with users
		file, err := os.OpenFile("list.json", os.O_RDWR|os.O_APPEND, 0666)
		defer file.Close()

		//read file and unmarshall json to []users
		b, err := ioutil.ReadAll(file)
		var alUsrs model.AllUsers
		err = json.Unmarshal(b, &alUsrs.Users)
		checkError(err)

		var allID []int
		for _, usr := range alUsrs.Users {
			allID = append(allID, usr.Id)
		}
		for i, usr := range alUsrs.Users {
			if model.IsValueInSlice(allID, id) != true {
				http.Redirect(w, r, "/deleteuser/notsuccededdelete", 302)
				return
			}
			if usr.Id != id {
				continue
			} else {
				alUsrs.Users = append(alUsrs.Users[:i], alUsrs.Users[i+1:]...)
			}

		}
		newUserBytes, err := json.MarshalIndent(&alUsrs.Users, "", " ")
		checkError(err)
		ioutil.WriteFile("list.json", newUserBytes, 0666)
		http.Redirect(w, r, "/deleted", 301)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//function to add user
func AddNewUserFunc(w http.ResponseWriter, r *http.Request) {

	//creating new instance and checking method
	newUser := &model.User{}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/addNewUser.html")
		t.Execute(w, nil)

	} else {
		resBool, errStr := checkFormValue(w, r, "name", "address")
		if resBool == false {
			t, err := template.ParseFiles("templates/wrongInput.html")
			checkError(err)
			t.Execute(w, errStr)
			return
		}
		newUser.Name = r.FormValue("name")
		newUser.Address = r.FormValue("address")
		//cookie := session.CreateToken(newUser)

		var err error
		checkError(err)

		//open file
		file, err := os.OpenFile("list.json", os.O_RDWR, 0644)
		checkError(err)
		defer file.Close()

		//read file and unmarshall json file to slice of users
		b, err := ioutil.ReadAll(file)
		var allUsers model.AllUsers
		err = json.Unmarshal(b, &allUsers.Users)
		checkError(err)
		max := 0

		//generation of id(last id at the json file+1)
		for _, usr := range allUsers.Users {
			if usr.Id > max {
				max = usr.Id
			}
		}
		id := max + 1
		newUser.Id = id

		//appending newUser to slice of all Users and rewrite json file
		allUsers.Users = append(allUsers.Users, newUser)
		newUserBytes, err := json.MarshalIndent(&allUsers.Users, "", " ")
		checkError(err)
		ioutil.WriteFile("list.json", newUserBytes, 0666)
		http.Redirect(w, r, "/", 301)

	}

}

//Index page handler
func IndexFunc(responseWriter http.ResponseWriter, request *http.Request) {
	au := model.ShowAllUsers()
	t, err := template.ParseFiles("templates/indexPage.html")
	checkError(err)
	t.Execute(responseWriter, au)
}

func QrCodeCreate(res http.ResponseWriter, r *http.Request) {

	var png []byte
	png, err := qrcode.Encode("https://example.org ", qrcode.Medium, 256)
	if err != nil {
	}
	res.Write(png)
	/*session.GetCookie()

	http.ServeFile(res, r, "templates/qrSite.html")*/

}
