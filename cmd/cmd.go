package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Lauft() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the day you are looking for. (Day Format: DD.MM.YYYY)")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hello, Yourself")
			validDateFormat("23.13.2323")
			break
		}
	}

}

func validDateFormat(date string) bool {
	result := true
	var compareOne []string
	compareOne = strings.Split(date, ".")
	for index, element := range compareOne {
		if index == 0 || index == 1 {
			if len(element) != 2 {
				return false
			}
			for _, innerElement := range element {
				if innerElement < 48 || innerElement > 57 {
					return false
				}
			}
		} else if index == 3 {
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

	return result

}
