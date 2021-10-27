package token

import (
	"math/rand"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var newTokenMap = map[string]string{}
var oldTokenMap = map[string]string{}

func CreateAndUpdateToken() string {
	tokenLength := 12
	tokenRune := make([]rune, tokenLength)
	for i := range tokenRune {
		tokenRune[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(tokenRune)
}

func CreateAndUpdateTokenMap(loc []string) {
	for i := 0; i < len(newTokenMap); i++ {
		oldTokenMap[loc[i]] = newTokenMap[loc[i]]
	}
	for i := 0; i < len(loc); i++ {
		newTokenMap[loc[i]] = CreateAndUpdateToken()
	}
}

func GetTokenByLocation(loc string) string {
	return newTokenMap[loc]
}

func ValidateTokenByLocation(token string, location string) string {
	if newTokenMap[location] == token || oldTokenMap[location] == token {
		return "true"
	} else {
		return "false"
	}
}
