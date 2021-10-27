package log

import (
	"io/ioutil"
	"os"
	"time"
)

func WriteToFile(login bool, textToWrite string) {
	var today = getDateInFormat()
	if doesFileExists("src/log/files/" + today + ".txt") {
		readAndWriteAndSave(login, textToWrite, today)
	} else {
		createFile(today)
		readAndWriteAndSave(login, textToWrite, today)
	}
}

func getDateInFormat() string {
	var date = time.Now().Format(time.RFC3339)
	today := date[0:10]
	return today
}

func getTimeStamp() string {
	var timestamp = time.Now().Format(time.Stamp)
	return timestamp
}

func createText(login bool, textToWrite string) string {
	if login {
		var text = "LOGIN, " + getTimeStamp() + ", " + textToWrite + ";\n"
		return text
	} else {
		var text = "LOGOUT, " + getTimeStamp() + ", " + textToWrite + ";\n"
		return text
	}
}

func panicHandling(err error) {
	if err != nil {
		panic(err)
	}
}

func doesFileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readAndWriteAndSave(login bool, textToWrite string, today string) {
	file, err := ioutil.ReadFile("src/log/files/" + today + ".txt")
	panicHandling(err)

	file = append(file, []byte(createText(login, textToWrite))...)

	err = ioutil.WriteFile("src/log/files/"+today+".txt", file, 0644)
	panicHandling(err)
}

func createFile(today string) {
	createFile, err := os.Create("src/log/files/" + today + ".txt")
	panicHandling(err)

	err = createFile.Close()
	panicHandling(err)
}
