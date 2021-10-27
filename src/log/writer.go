package log

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func exists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WriteLoginToFile(textToWrite string) {
	var date = time.Now().Format(time.RFC3339)
	today := date[0:10]
	log.Print(today)

	if exists("src/log/files/" + today + ".txt") {
		fmt.Println("Example file exists")
		// read the whole file at once
		file, err := ioutil.ReadFile("src/log/files/" + today + ".txt")
		if err != nil {
			panic(err)
		}

		var timestamp = time.Now().Format(time.Stamp)
		var text = "LOGIN, " + timestamp + ", " + textToWrite + ";\n"

		file = append(file, []byte(text)...)

		// write the whole body at once
		err = ioutil.WriteFile("src/log/files/"+today+".txt", file, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Example file does not exist (or is a directory)")
		emptyFile, err := os.Create("src/log/files/" + today + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(emptyFile)
		emptyFile.Close()
		// read the whole file at once
		file, err := ioutil.ReadFile("src/log/files/" + today + ".txt")
		if err != nil {
			panic(err)
		}

		var timestamp = time.Now().Format(time.Stamp)
		var text = "LOGIN, " + timestamp + ", " + textToWrite + ";\n"

		file = append(file, []byte(text)...)

		// write the whole body at once
		err = ioutil.WriteFile("src/log/files/"+today+".txt", file, 0644)
		if err != nil {
			panic(err)
		}
	}

}

func WriteLogoutToFile(textToWrite string) {
	var date = time.Now().Format(time.RFC3339)
	today := date[0:10]
	log.Print(today)

	if exists("src/log/files/" + today + ".txt") {
		fmt.Println("Example file exists")
		// read the whole file at once
		file, err := ioutil.ReadFile("src/log/files/" + today + ".txt")
		if err != nil {
			panic(err)
		}

		var timestamp = time.Now().Format(time.Stamp)
		var text = "LOGOUT, " + timestamp + ", " + textToWrite + ";\n"

		file = append(file, []byte(text)...)

		// write the whole body at once
		err = ioutil.WriteFile("src/log/files/"+today+".txt", file, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Example file does not exist (or is a directory)")
		emptyFile, err := os.Create("src/log/files/" + today + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(emptyFile)
		emptyFile.Close()
		// read the whole file at once
		file, err := ioutil.ReadFile("src/log/files/" + today + ".txt")
		if err != nil {
			panic(err)
		}

		var timestamp = time.Now().Format(time.Stamp)
		var text = "LOGIN, " + timestamp + ", " + textToWrite + ";\n"

		file = append(file, []byte(text)...)

		// write the whole body at once
		err = ioutil.WriteFile("src/log/files/"+today+".txt", file, 0644)
		if err != nil {
			panic(err)
		}
	}
}
