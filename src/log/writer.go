package log

import (
	"io/ioutil"
	"time"
)

func WriteLoginToFile(textToWrite string) {
	// read the whole file at once
	b, err := ioutil.ReadFile("src/log/tmp/logOfLogins.txt")
	if err != nil {
		panic(err)
	}
	var timestamp = time.Now().Format(time.Stamp)

	var c = "LOGIN, " + timestamp + ", " + textToWrite
	b = append(b, []byte(c)...)

	// write the whole body at once
	err = ioutil.WriteFile("src/log/tmp/logOfLogins.txt", b, 0644)
	if err != nil {
		panic(err)
	}
}

func WriteLogoutToFile(textToWrite string) {
	// read the whole file at once
	b, err := ioutil.ReadFile("src/log/tmp/logOfLogins.txt")
	if err != nil {
		panic(err)
	}
	var timestamp = time.Now().Format(time.Stamp)

	var c = "LOGOUT, " + timestamp + ", " + textToWrite
	b = append(b, []byte(c)...)

	// write the whole body at once
	err = ioutil.WriteFile("src/log/tmp/logOfLogins.txt", b, 0644)
	if err != nil {
		panic(err)
	}
}
