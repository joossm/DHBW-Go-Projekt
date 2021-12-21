// 5807262
// 9899545
// 8622410

package token

import (
	"math/rand"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var newTokenMap = map[string]string{}
var oldTokenMap = map[string]string{}

// CreateAndUpdateToken creates a new token
func CreateAndUpdateToken() string {
	tokenLength := 12
	tokenRune := make([]rune, tokenLength)
	for i := range tokenRune {
		tokenRune[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(tokenRune)
}

// CreateAndUpdateTokenMap creates a new token and updates the old token
func CreateAndUpdateTokenMap(loc []string) {
	for i := 0; i < len(newTokenMap); i++ {
		oldTokenMap[loc[i]] = newTokenMap[loc[i]]
	}
	for i := 0; i < len(loc); i++ {
		newTokenMap[loc[i]] = CreateAndUpdateToken()
	}
}

// GetTokenByLocation returns the token for a specific location
func GetTokenByLocation(loc string) string {
	return newTokenMap[loc]
}

// ValidateTokenByLocation validates a token for a specific location
func ValidateTokenByLocation(token string, location string) bool {
	if newTokenMap[location] == token || oldTokenMap[location] == token {
		return true
	} else {
		return false
	}
}
