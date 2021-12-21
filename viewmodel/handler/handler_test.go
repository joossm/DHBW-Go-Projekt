// 5807262
// 9899545
// 8622410

package handler

import (
	"GoProjekt/viewmodel/config"
	"GoProjekt/viewmodel/token"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestLoginWithNoToken(t *testing.T) {
	request, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusForbidden {
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
	config.InitForTesting()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	request, err := http.NewRequest("GET", "/login?token=FPLLNGZIEYOH&location=Mosbach", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestLogout(t *testing.T) {
	config.InitForTesting()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	request, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LogoutUser)
	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code: %v want %v", status, http.StatusOK)
	}

	request, err = http.NewRequest("POST", "/end", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCombineText(t *testing.T) {
	var name = "Max Mustermann"
	var address = "74081 Heilbronn"
	var location = "Mosbach"
	assert.Equal(t, name+", "+address+", "+location, combineText(name, address, location))
}
func TestGetLocation(t *testing.T) {
	request, err := http.NewRequest("GET", "/Mosbach?", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, "Mosbach", getLocation(request))
}
func TestEnd(t *testing.T) {
	config.InitForTesting()
	request, err := http.NewRequest("GET", "/end", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ReLogin)

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}

	request, err = http.NewRequest("POST", "/end", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}

}
func TestSelectLocation(t *testing.T) {
	config.InitForTesting()
	request, err := http.NewRequest("GET", "/location", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SelectLocation)

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}

	request, err = http.NewRequest("POST", "/location", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}
}
func TestLoginUserForm(t *testing.T) {
	config.InitForTesting()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	form := url.Values{}
	form.Set("firstName", "Emil")
	form.Set("lastName", "Bartoldus")
	form.Set("zipCode", "34414")
	form.Set("cityName", "Germete")
	form.Set("streetName", "Am Waldwinkel")
	form.Set("houseNumber", "12")
	var timeStamp = time.Now().Format(time.RFC3339)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/login?token=FPLLNGZIEYOH&location=Mosbach", nil)
	request.PostForm = form

	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(recorder, request)

	var before, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + time.Now().Format(time.RFC3339)[0:10] + ".txt")
	beforeLine := strings.Split(string(before), "\n")

	var check, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + time.Now().Format(time.RFC3339)[0:10] + ".txt")
	assert.Equal(t, []byte("LOGIN, " + timeStamp + ", Emil Bartoldus, 34414 Germete Am Waldwinkel 12, NO INFORMATION;\n")[0:3], check[0:3])
	assert.FileExists(t, flag.Lookup("logfilePath").Value.String()+time.Now().Format(time.RFC3339)[0:10]+".txt")
	lines := strings.Split(string(check), "\n")

	var data []byte
	for i := 0; i < len(beforeLine)-2; i++ {
		data = append(data, []byte(lines[i]+"\n")...)
	}

	_ = ioutil.WriteFile(flag.Lookup("logfilePath").Value.String()+time.Now().Format(time.RFC3339)[0:10]+".txt", data, 0644)
}

func TestLoginUserFormWrongInput(t *testing.T) {
	config.InitForTesting()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	form := url.Values{}
	form.Set("firstName", "123!")
	form.Set("lastName", "Bartoldus")
	form.Set("zipCode", "34414M!")
	form.Set("cityName", "Germete")
	form.Set("streetName", "Am Waldwinkel")
	form.Set("houseNumber", "12")

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/login?token=FPLLNGZIEYOH&location=Mosbach", nil)
	request.PostForm = form

	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(recorder, request)

	var before, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + time.Now().Format(time.RFC3339)[0:10] + ".txt")
	beforeLine := strings.Split(string(before), "\n")

	var check, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + time.Now().Format(time.RFC3339)[0:10] + ".txt")
	assert.FileExists(t, flag.Lookup("logfilePath").Value.String()+time.Now().Format(time.RFC3339)[0:10]+".txt")
	lines := strings.Split(string(check), "\n")

	var data []byte
	for i := 0; i < len(beforeLine)-1; i++ {
		data = append(data, []byte(lines[i]+"\n")...)
	}

	_ = ioutil.WriteFile(flag.Lookup("logfilePath").Value.String()+time.Now().Format(time.RFC3339)[0:10]+".txt", data, 0644)
}
func TestQrCodeCreate(t *testing.T) {
	config.InitForTesting()

	request, err := http.NewRequest("GET", "/Mosbach?", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(QrCodeCreate)

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}

	request, err = http.NewRequest("POST", "/location", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}

}

func TestInformationsFromCookies(t *testing.T) {
	config.InitForTesting()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	// Create a new HTTP Recorder (implements http.ResponseWriter)
	recorder := httptest.NewRecorder()

	// Drop a cookie on the recorder.
	http.SetCookie(recorder, &http.Cookie{Name: "name", Value: "expected"})
	http.SetCookie(recorder, &http.Cookie{Name: "address", Value: "expected"})
	http.SetCookie(recorder, &http.Cookie{Name: "location", Value: "expected"})
	// Copy the Cookie over to a new Request
	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}

	// Extract the dropped cookie from the request.
	cookie, err := request.Cookie("name")

	require.NoError(t, err, "Failed to read 'test' Cookie: %v", err)
	require.Equal(t, cookie.Value, "expected")
	cookie, err = request.Cookie("address")

	require.NoError(t, err, "Failed to read 'test' Cookie: %v", err)
	require.Equal(t, cookie.Value, "expected")
	cookie, err = request.Cookie("location")

	require.NoError(t, err, "Failed to read 'test' Cookie: %v", err)
	require.Equal(t, cookie.Value, "expected")
	assert.Equal(t, "expected", informationsFromCookies("name", request))
	assert.Equal(t, "expected", informationsFromCookies("address", request))
	assert.Equal(t, "expected", informationsFromCookies("location", request))

}
func TestProofIfInSameLocation(t *testing.T) {
	recorder := httptest.NewRecorder()
	config.InitForTesting()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	request, err := http.NewRequest("GET", "/login?token=FPLLNGZIEYOH&location=Mosbach", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(recorder, request)
	var location = proofIfLoginInSameLocation(request)
	assert.Equal(t, "Mosbach", location)

	newRequest, err := http.NewRequest("GET", "/login?token=FPLLNGZIEYOH&location=", nil)
	if err != nil {
		t.Fatal(err)
	}
	location = proofIfLoginInSameLocation(newRequest)
	assert.Equal(t, "NO INFORMATION", location)
}

/*func TestLoginWithCookies(t *testing.T) {
	config.InitForTesting()
	locations := []string{"Mosbach"}
	token.CreateAndUpdateTokenMap(locations)

	rr := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/login?token=FPLLNGZIEYOH&location=Mosbach", nil)
	if err != nil {
		t.Fatal(err)
	}
	http.SetCookie(rr, &http.Cookie{Name: "name", Value: "Mosbach"})
	http.SetCookie(rr, &http.Cookie{Name: "address", Value: "Mosbach"})
	http.SetCookie(rr, &http.Cookie{Name: "location", Value: "Mosbach"})
	request = &http.Request{Header: http.Header{"Cookie": rr.HeaderMap["Set-Cookie"]}}


	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, request)


}
*/
