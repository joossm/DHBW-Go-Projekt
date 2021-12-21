// 5807262
// 9899545
// 8622410

package analyzer

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestValidateDateFormat(t *testing.T) {
	assert.Equal(t, false, validDateFormat("01/01/2019"))
	assert.Equal(t, false, validDateFormat("01-01-2019"))
	assert.Equal(t, false, validDateFormat("01.01.2019"))

	assert.Equal(t, false, validDateFormat("2019/01/01"))
	assert.Equal(t, false, validDateFormat("2019.01.01"))
	assert.Equal(t, false, validDateFormat("2019.01.A1"))

	assert.Equal(t, false, validDateFormat("20190-01-01"))
	assert.Equal(t, false, validDateFormat("2019-001-01"))
	assert.Equal(t, false, validDateFormat("2019-01-001"))
	assert.Equal(t, false, validDateFormat("201a-01-01"))
	assert.Equal(t, false, validDateFormat("2019-0a-01"))
	assert.Equal(t, false, validDateFormat("2019-01-0a"))
	assert.Equal(t, false, validDateFormat("2019-01-01-M"))

	assert.Equal(t, true, validDateFormat("2019-01-01"))
}
func TestNewContactCheck(t *testing.T) {
	assert.Equal(t, false, newContactCheck(true, "Max Mustermann"))
	assert.Equal(t, false, newContactCheck(false, "Max Mustermann"))
}

func TestCreateCSV(t *testing.T) {
	InitForTesting()

	var fileN = fileTest()
	createCSV(fileN, "Mosbach", "2001-01-01")
	assert.FileExists(t, "../../src/log/files/2001-01-01_Mosbach.csv")
	assert.Equal(t, true, fileExists("2001-01-01"))
	assert.Equal(t, false, fileExists("../../party.go"))
	_ = os.Remove("../../src/log/files/2001-01-01_Mosbach.csv")
}

func TestFindStartEndTimes(t *testing.T) {
	InitForTesting()
	file, err := ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + "2001-01-01" + ".txt")
	if err != nil {
		panic(err)
	}
	fileString := strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}

	var startTime, endTime []time.Time
	var startIndex, endIndex []int

	startTime, endTime, startIndex, endIndex = findStartEndTimes("Mosbach", "Max Mustermann", "89150 Laichingen Bei der Kirche 9", fileString)
	assert.Equal(t, "2001-01-01 11:14:30 +0100 CET", startTime[0].String())
	assert.Equal(t, "2001-01-01 11:19:38 +0100 CET", endTime[0].String())
	assert.Equal(t, 0, startIndex[0])
	assert.Equal(t, 5, endIndex[0])

	file, err = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + "2001-01-02" + ".txt")
	if err != nil {
		panic(err)
	}
	fileString = strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}

	startTime, endTime, startIndex, endIndex = findStartEndTimes("Mosbach", "Max Mustermann", "89150 Laichingen Bei der Kirche 9", fileString)
	assert.Equal(t, "2001-01-01 11:14:30 +0100 CET", startTime[0].String())
	assert.Equal(t, "2001-01-01 11:19:38 +0100 CET", endTime[0].String())
	assert.Equal(t, 0, startIndex[0])
	assert.Equal(t, 5, endIndex[0])

	file, err = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + "2001-01-03" + ".txt")
	if err != nil {
		panic(err)
	}
	fileString = strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}

	startTime, endTime, startIndex, endIndex = findStartEndTimes("Mosbach", "Max Mustermann", "89150 Laichingen Bei der Kirche 9", fileString)
	assert.Equal(t, "2001-01-01 00:00:00 +0100 CET", startTime[0].String())
	assert.Equal(t, "2001-01-01 01:19:38 +0100 CET", endTime[0].String())
	assert.Equal(t, -1, startIndex[0])
	assert.Equal(t, 0, endIndex[0])

}

func InitForTesting() {
	err := flag.Set("endPagePath", "../../html/reLoginPage.html")
	if err != nil {
		return
	}
	err = flag.Set("loginPagePath", "../../html/loginPage.html")
	if err != nil {
		return
	}
	err = flag.Set("logoutPagePath", "../../html/logoutPage.html")
	if err != nil {
		return
	}
	err = flag.Set("logfilePath", "../../src/log/files/")
	if err != nil {
		return
	}
	err = flag.Set("locationOverviewPath", "../../html/locationOverview.html")
	if err != nil {
		return
	}
	err = flag.Set("wrongInputPath", "../../html/wrongInput.html")
	if err != nil {
		return
	}
	err = flag.Set("certFilePath", "../../server.crt")
	if err != nil {
		return
	}
}

func fileTest() []string {
	var file [6]string
	file[0] = "LOGIN, 2001-01-01T11:14:30+01:00, Max Mustermann, 89150 Laichingen Bei der Kirche 9, Mosbach;\n"
	file[1] = "LOGIN, 2001-01-01T11:15:34+01:00, Manuel Neuer, 74081 Heilbronn Robert-Bosch-Strasse 23, Mosbach;\n"
	file[2] = "LOGIN, 2001-01-01T11:16:35+01:00, Serge Gnabry, 74082 Heilbronn NRW-City 25, Mosbach;\n"
	file[3] = "LOGOUT, 2001-01-01T11:17:36+01:00, Serge Gnabry, 74082 Heilbronn NRW-City 25, Mosbach;\n"
	file[4] = "LOGOUT, 2001-01-01T11:18:37+01:00, Manuel Neuer, 74081 Heilbronn Robert-Bosch-Strasse 23, Mosbach;\n"
	file[5] = "LOGOUT, 2001-01-01T11:19:38+01:00, Max Mustermann, 89150 Laichingen Bei der Kirche 9, Mosbach;\n"
	return file[0:6]
}

func TestFindVisitor(t *testing.T) {
	InitForTesting()
	file, err := ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + "2001-01-01" + ".txt")
	if err != nil {
		panic(err)
	}
	fileString := strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}

	var visitorsList []string
	visitorsList = findVisitors("Mosbach", fileString, "Max Mustermann", "89150 Laichingen Bei der Kirche 9")
	assert.Equal(t, "Manuel Neuer,74081 Heilbronn Robert-Bosch-Strasse 23", visitorsList[0])
	assert.Equal(t, "Serge Gnabry,74082 Heilbronn NRW-City 25", visitorsList[1])
}

func TestReadFileToStrings(t *testing.T) {
	InitForTesting()
	var list []string

	list = readFileToStrings("2001-01-01")
	assert.Equal(t, "LOGIN, 2001-01-01T11:14:30+01:00, Max Mustermann, 89150 Laichingen Bei der Kirche 9, Mosbach", list[0])
	assert.Equal(t, "\nLOGIN, 2001-01-01T11:15:34+01:00, Manuel Neuer, 74081 Heilbronn Robert-Bosch-Strasse 23, Mosbach", list[1])
	assert.Equal(t, "\nLOGIN, 2001-01-01T11:16:35+01:00, Serge Gnabry, 74082 Heilbronn NRW-City 25, Mosbach", list[2])
	assert.NotEqual(t, "LOGIN, 2001-01-01T11:14:30+01:00, Max Mustermann, Mosbach", list[0])
	assert.NotEqual(t, "\nLOGIN, 2001-01-01T11:15:34+01:00, 74081 Heilbronn Robert-Bosch-Strasse 23, Mosbach", list[1])
	assert.NotEqual(t, "\nLOGIN, Serge Gnabry, 74082 Heilbronn NRW-City 25, Mosbach", list[2])
}

func TestPersonExisits(t *testing.T) {
	InitForTesting()
	var list []string
	list = readFileToStrings("2001-01-01")
	assert.Equal(t, true, personExists(list, "Max Mustermann", "89150 Laichingen Bei der Kirche 9"))
	assert.Equal(t, false, personExists(list, "Tim Mustermann", "89150 Laichingen Bei der Kirche 9"))
	assert.Equal(t, false, personExists(list, "Max Mustermann", "Laichingen Bei der Kirche 9"))
	assert.Equal(t, false, personExists(list, "Max Mustermann", "89150 Laichingen Bei der Kirche"))
	assert.Equal(t, false, personExists(list, "Max Mustermann", "89150 Bei der Kirche 9"))
}

func TestCalcContactTime(t *testing.T) {
	InitForTesting()
	file, err := ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + "2001-01-06" + ".txt")
	if err != nil {
		panic(err)
	}
	fileString := strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}

	var startTime, endTime []time.Time
	var startIndex, endIndex []int
	var startTime2, endTime2 []time.Time
	var startIndex2, endIndex2 []int

	startTime, endTime, startIndex, endIndex = findStartEndTimes("Mosbach", "Max Mustermann", "89150 Laichingen Bei der Kirche 9", fileString)
	startTime2, endTime2, startIndex2, endIndex2 = findStartEndTimes("Mosbach", "Manuel Neuer", "74081 Heilbronn Robert-Bosch-Strasse 23", fileString)

	assert.Equal(t, true, calcContactTime(startTime, endTime, startIndex, endIndex, startTime2, endTime2, startIndex2, endIndex2, "Max Mustermann"))
	assert.Equal(t, true, calcContactTime(startTime, endTime, startIndex, endIndex, startTime, endTime, startIndex, endIndex, "Max Mustermann"))
	assert.Equal(t, true, calcContactTime(startTime2, endTime2, startIndex2, endIndex2, startTime, endTime, startIndex, endIndex, "Max Mustermann"))
	assert.Equal(t, true, calcContactTime(startTime2, endTime2, startIndex2, endIndex2, startTime2, endTime2, startIndex2, endIndex2, "Max Mustermann"))

	file, err = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + "2001-01-07" + ".txt")
	if err != nil {
		panic(err)
	}
	fileString = strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}

	startTime, endTime, startIndex, endIndex = findStartEndTimes("Mosbach", "Max Mustermann", "89150 Laichingen Bei der Kirche 9", fileString)
	startTime2, endTime2, startIndex2, endIndex2 = findStartEndTimes("Mosbach", "Manuel Neuer", "74081 Heilbronn Robert-Bosch-Strasse 23", fileString)

	assert.Equal(t, true, calcContactTime(startTime, endTime, startIndex, endIndex, startTime2, endTime2, startIndex2, endIndex2, "Max Mustermann"))
	assert.Equal(t, false, calcContactTime(startTime, endTime, startIndex, endIndex, startTime, endTime, startIndex, endIndex, "Max Mustermann"))
	assert.Equal(t, true, calcContactTime(startTime2, endTime2, startIndex2, endIndex2, startTime, endTime, startIndex, endIndex, "Max Mustermann"))
	assert.Equal(t, false, calcContactTime(startTime2, endTime2, startIndex2, endIndex2, startTime2, endTime2, startIndex2, endIndex2, "Max Mustermann"))
}
func TestPlaceExisits(t *testing.T) {
	InitForTesting()
	var list []string
	list = readFileToStrings("2001-01-01")
	assert.Equal(t, true, placeExists(list, "Mosbach"))
	assert.Equal(t, false, placeExists(list, "Mosbach DHBW"))
	assert.Equal(t, false, placeExists(list, "Monaco"))
	assert.Equal(t, false, placeExists(list, "Moskau"))
}
func TestPlacesOfPersons(t *testing.T) {
	InitForTesting()
	var list []string
	list = readFileToStrings("2001-01-01")
	placesOfPerson := findPlacesOfPerson(list, "Manuel Neuer", "74081 Heilbronn Robert-Bosch-Strasse 23")
	assert.Equal(t, "Mosbach", placesOfPerson[0])
	assert.NotEqual(t, "Monaco", placesOfPerson[0])
}

func TestFindPossibleContacts(t *testing.T) {
	InitForTesting()
	var list []string
	list = readFileToStrings("2001-01-01")
	assert.Equal(t, true, findPossibleContacts("Mosbach", "Manuel Neuer", "74081 Heilbronn Robert-Bosch-Strasse 23", list))
	assert.Equal(t, false, findPossibleContacts("Mosbach", "Max Berberich", "74081 Heilbronn Robert-Bosch-Strasse 23", list))
}
