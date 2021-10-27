package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Lauft() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the day you are looking for. (Day Format: YYYY-MM-DD)")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if validDateFormat(text) {
		fmt.Println("wahr")
	}
}

func validDateFormat(date string) bool {
	var compareOne []string
	compareOne = strings.Split(date, "-")
	for index, element := range compareOne {
		if index == 1 || index == 2 {
			if len(element) != 2 {
				return false
			}
			for _, innerElement := range element {
				if innerElement < 48 || innerElement > 57 {
					return false
				}
			}
		} else if index == 0 {
			if len(element) != 4 {
				return false
			}
			for _, innerElement := range element {
				if innerElement < 48 || innerElement > 57 {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}

func readFile(date string) {
	file, err := ioutil.ReadFile("src/log/files/" + date + ".txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}
