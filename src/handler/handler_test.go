package handler

import (
	"GoProjekt/src/model/config"
	"GoProjekt/src/token"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginWithNoToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("status code: got %v want %v", status, http.StatusForbidden)
	}

}
func TestLoginWithWrongToken(t *testing.T) {
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)
	request, err := http.NewRequest("GET", "/login?token=FALSCH&location=Mosbach", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(recorder, request)

	if code := recorder.Code; code != http.StatusForbidden {
		t.Errorf("expected %v got %v", http.StatusForbidden, code)
	}
}
func TestLoginWithRightToken(t *testing.T) {
	config.InitByMatthias()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	req, err := http.NewRequest("GET", "/login?token=FPLLNGZIEYOH&location=Mosbach", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestLogout(t *testing.T) {
	config.InitByMatthias()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogoutUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: %v want %v", status, http.StatusOK)
	}

}
func TestAlreadyLoggedIn(t *testing.T) {

}
func TestCombineText(t *testing.T) {
	var name = "Max Mustermann"
	var address = "74081 Heilbronn"
	var location = "Mosbach"
	assert.Equal(t, name+", "+address+", "+location, combineText(name, address, location))
}
func TestGetLocation(t *testing.T) {
	req, err := http.NewRequest("GET", "/Mosbach?", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, "Mosbach", getLocation(req))
}
