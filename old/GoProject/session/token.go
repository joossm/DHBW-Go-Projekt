package session

import (
	"GoProject/model"
	"net/http"
	"time"
)

var tokenEncodeString string = "something"

func CreateToken(user *model.User) http.Cookie {
	name := user.Name
	expiration := time.Now().Add(5 * time.Minute)
	cookie := http.Cookie{Name: name, Value: "astaxie", Expires: expiration}
	return cookie
}

func parseToken(cookie http.Cookie) bool {
	timeStamp := time.Now()
	timeStampCookie := cookie.Expires
	if timeStampCookie.After(timeStamp) {
		return true
	} else {
		return false
	}
}
func GetCookie() []*http.Cookie {
	client := http.Client{}
	var cookie []*http.Cookie
	url := ""
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req) //send request
	if err != nil {
		return nil
	}
	cookie = resp.Cookies()
	return cookie //save cookies
}
