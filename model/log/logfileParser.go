// 5807262
// 9899545
// 8622410

package log

import (
	"flag"
	"io/ioutil"
	"os"
	"time"
)

// WriteToFile writes the log to a file
func WriteToFile(login bool, textToWrite string) {
	var today = getDateInFormat()
	if doesFileExists(flag.Lookup("logfilePath").Value.String() + today + ".txt") {
		readAndWriteAndSave(login, textToWrite, today)
	} else {
		createFile(today)
		readAndWriteAndSave(login, textToWrite, today)
	}
}

// getDateInFormat returns the current date in the format YYYY-MM-DD
func getDateInFormat() string {
	return time.Now().Format(time.RFC3339)[0:10]
}

// getTimeInFormat returns the current time in the format HH:MM:SS
func getTimeStamp() string {
	return time.Now().Format(time.RFC3339)
}

// createText creates the text to write to the file
func createText(login bool, textToWrite string) string {
	if login {
		var text = "LOGIN, " + getTimeStamp() + ", " + textToWrite + ";\n"
		return text
	} else {
		var text = "LOGOUT, " + getTimeStamp() + ", " + textToWrite + ";\n"
		return text
	}
}

// panicHandler handles the panic that may occur when writing to the file
func panicHandling(err error) {
	if err != nil {
		panic(err)
	}
}

// doesFileExists checks if the file exists
func doesFileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// readAndWriteAndSave reads the file and writes to it, afterwards it saves the file
func readAndWriteAndSave(login bool, textToWrite string, today string) {
	file, err := ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + today + ".txt")
	panicHandling(err)

	file = append(file, []byte(createText(login, textToWrite))...)

	err = ioutil.WriteFile(flag.Lookup("logfilePath").Value.String()+today+".txt", file, 0644)
	panicHandling(err)
}

// createFile creates a new file
func createFile(today string) {
	createFile, err := os.Create(flag.Lookup("logfilePath").Value.String() + today + ".txt")
	panicHandling(err)

	err = createFile.Close()
	panicHandling(err)
}
