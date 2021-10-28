package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAndUpdateToken(t *testing.T) {
	var oldT, newT string
	var same bool
	oldT = CreateAndUpdateToken()
	newT = CreateAndUpdateToken()
	if oldT == newT {
		same = true
	} else {
		same = false
	}
	assert.Equal(t, false, same)
}
func TestCreateAndUpdateTokenMap(t *testing.T) {
	var locations = []string{"TestLocation"}
	var oldT, newT string
	var same bool
	CreateAndUpdateTokenMap(locations)
	oldT = GetTokenByLocation(locations[0])
	CreateAndUpdateTokenMap(locations)
	newT = GetTokenByLocation(locations[0])
	if oldT == newT {
		same = true
	} else {
		same = false
	}
	assert.Equal(t, false, same)

}
func TestGetTokenByLocation(t *testing.T) {
	var locations = []string{"TestLocation"}
	var rigthParameter, wrongParameter string
	CreateAndUpdateTokenMap(locations)
	rigthParameter = GetTokenByLocation(locations[0])
	wrongParameter = GetTokenByLocation("LocationTest")
	var i = len(rigthParameter)
	assert.Equal(t, 12, i)
	assert.Equal(t, "", wrongParameter)
}
func TestValidateTokenByLocation(t *testing.T) {
	var locations = []string{"TestLocation", "LocationTest"}
	var token string
	CreateAndUpdateTokenMap(locations)
	token = GetTokenByLocation(locations[0])
	var falseToken = "O1PL5L64WDRT"
	assert.Equal(t, true, ValidateTokenByLocation(token, locations[0]))
	assert.Equal(t, false, ValidateTokenByLocation(token, locations[1]))
	assert.Equal(t, false, ValidateTokenByLocation(falseToken, locations[0]))
	assert.Equal(t, false, ValidateTokenByLocation(falseToken, locations[0]))

}
