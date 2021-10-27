package log

import (
	"io/ioutil"
	"time"
)

func WriteLoginToFile(textToWrite string) {
	// read the whole file at once
	file, err := ioutil.ReadFile("src/log/tmp/logOfLogins.txt")
	if err != nil {
		panic(err)
	}

	var timestamp = time.Now().Format(time.Stamp)
	var text = "LOGIN, " + timestamp + ", " + textToWrite + ";\n"

	file = append(file, []byte(text)...)

	// write the whole body at once
	err = ioutil.WriteFile("src/log/tmp/logOfLogins.txt", file, 0644)
	if err != nil {
		panic(err)
	}
}

func WriteLogoutToFile(textToWrite string) {
	// read the whole file at once
	file, err := ioutil.ReadFile("src/log/tmp/logOfLogins.txt")
	if err != nil {
		panic(err)
	}
	var timestamp = time.Now().Format(time.Stamp)

	var text = "LOGOUT, " + timestamp + ", " + textToWrite + ";\n"
	file = append(file, []byte(text)...)

	// write the whole body at once
	err = ioutil.WriteFile("src/log/tmp/logOfLogins.txt", file, 0644)
	if err != nil {
		panic(err)
	}
}
