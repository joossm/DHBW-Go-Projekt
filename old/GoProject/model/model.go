package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func IsValueInSlice(slice []int, value int) (result bool) {
	for _, n := range slice {
		if n == value {
			return true
		}

	}
	return false

}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type AllUsers struct {
	Users []*User
}

func ShowAllUsers() (au *AllUsers) {
	file, err := os.OpenFile("list.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	b, err := ioutil.ReadAll(file)
	var alUsrs AllUsers
	json.Unmarshal(b, &alUsrs.Users)
	checkError(err)
	return &alUsrs
}

type Places struct {
	XMLName xml.Name `xml:"places"`
	Place   []Places `xml:"place"`
}
