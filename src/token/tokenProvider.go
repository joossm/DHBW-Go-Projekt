package token

import (
	"math/rand"
	"net/http"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var m = map[int]string{}

func CreateToken(res http.ResponseWriter, r *http.Request) {
	n := 12
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	saveTokenToMap(string(b))
	//res.Write([]byte(string(b)))
	var s string = ValidateToken("3H3X9KCEGTA5")
	res.Write([]byte("Token Neu:" + m[0] + " Token Alt:" + m[1] + " Token validiert: " + s))
}
func CreateAndUpdateToken() string {
	tokenLength := 12
	tokenRune := make([]rune, tokenLength)
	for i := range tokenRune {
		tokenRune[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	saveTokenToMap(string(tokenRune))
	return string(tokenRune)
}
func saveTokenToMap(token string) {
	if len(m) >= 1 {
		var s = map[int]string{}
		s[0] = m[0]
		s[1] = token
		m[0] = s[1]
		m[1] = s[0]
	} else {
		m[0] = token
	}
}
func ValidateToken(token string) string {
	if m[0] == token || m[1] == token {
		return "true"
	} else {
		return "false"
	}
}
func GetActiveToken() string {
	return m[0]
}
