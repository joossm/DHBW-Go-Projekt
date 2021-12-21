// 5807262
// 9899545
// 8622410

package analyzer

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

// AnlayzeReport is the main struct for the analyzer
func AnlayzeReport() {
	fmt.Println("NOTE: All special characters have to be replaced! For Example: ß -> ss, ä ->ae")
	fmt.Println("Please enter the day you are looking for. (Date Format: YYYY-MM-DD)")
	var dateString string
	for {
		dateString = requestInput()
		if validDateFormat(dateString) {
			if fileExists(dateString) {
				break
			}
			fmt.Println("There are no entries for this day. Please enter another day. (Date Format: YYYY-MM-DD)")
		} else {
			fmt.Println("Wrong Date Format. Please use: YYYY-MM-DD")
		}
	}
	file := readFileToStrings(dateString)
	fmt.Println("What would you like to do? Enter the number of your preferred task")
	fmt.Println("1: Look up the places a person visited on this day")
	fmt.Println("2: Extract data for a place as a CSV-file")
	fmt.Println("3: Show contacts of a person")
	chosenOption := requestInput()
	if strings.Compare(chosenOption, "1") == 0 {
		fmt.Println("Enter the name of the Person you are looking for:")
		personName := requestInput()
		fmt.Println("Enter the address of the Person you are looking for:")
		fmt.Println("(format: postcode city street house number)")
		fmt.Println("(example: 74821 Mosbach Lohrtalweg 10")
		personAddress := requestInput()
		if personExists(file, personName, personAddress) {
			for _, place := range findPlacesOfPerson(file, personName, personAddress) {
				fmt.Println(place)
			}
		} else {
			fmt.Println("No entry for this Person found.")
		}
	} else if strings.Compare(chosenOption, "2") == 0 {
		fmt.Println("Enter the name of the place you want to create a CSV-file for:")
		nameOfPlace := requestInput()
		if placeExists(file, nameOfPlace) {
			createCSV(file, nameOfPlace, dateString)
		} else {
			fmt.Println("Place does not exist")
		}
	} else if strings.Compare(chosenOption, "3") == 0 {
		fmt.Println("Enter the name of the Person you are looking for:")
		personName := requestInput()
		fmt.Println("Enter the address of the Person you are looking for:")
		fmt.Println("(format: postcode city street house number)")
		fmt.Println("(example: 74821 Mosbach Lohrtalweg 10)")
		personAddress := requestInput()

		if personExists(file, personName, personAddress) {
			for _, place := range findPlacesOfPerson(file, personName, personAddress) {
				if findPossibleContacts(place, personName, personAddress, file) {
					fmt.Printf("All contacts for %v printed.\n", place)
				} else {
					fmt.Printf("No contacts for %v found. \n", place)
				}
			}
		} else {
			fmt.Println("No entry for this Person found.")
		}

	} else {
		fmt.Println("No valid choice")
	}
}

// Checks if the input is in the correct format
func requestInput() string {
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "windows" {
		text = strings.Replace(text, "\r\n", "", -1)
	}
	if runtime.GOOS == "linux" {
		text = strings.Replace(text, "\n", "", -1)
	}
	return text
}

// Checks if the date is in the correct format
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

// Checks if the file exists
func fileExists(date string) bool {
	_, err := ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + date + ".txt")
	if err != nil {
		return false
	}
	return true
}

// Reads the file to a string
func readFileToStrings(date string) []string {
	file, err := ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + date + ".txt")
	if err != nil {
		panic(err)
	}
	fileString := strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}
	return fileString
}

// Checks if the person exists
func personExists(file []string, name string, address string) bool {
	for _, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[2], name) == 0 {
			if strings.Compare(row[3], address) == 0 {
				return true
			}
		}
	}
	return false
}

// Finds all places of a person
func findPlacesOfPerson(file []string, name string, address string) []string {
	var places []string
	for _, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[2], name) == 0 {
			if strings.Compare(row[3], address) == 0 {
				newPlace := true
				for _, placeName := range places {
					if strings.Compare(placeName, row[4]) == 0 {
						newPlace = false
					}
				}
				if newPlace {
					places = append(places, row[4])
				}
			}
		}
	}
	return places
}

// Checks if the place exists
func placeExists(file []string, place string) bool {
	for _, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[4], place) == 0 {
			return true
		}
	}
	return false
}

// creates a CSV-file for a place
func createCSV(file []string, place string, lookedForDate string) {
	csvFile, err := os.Create(flag.Lookup("logfilePath").Value.String() + lookedForDate + "_" + place + ".csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(csvFile)
	for _, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[4], place) == 0 && strings.Contains(row[0], "LOGIN") {
			var newRow []string
			newRow = append(newRow, row[2], row[3])
			err := writer.Write(newRow)
			if err != nil {
				return
			}
		}
	}
	writer.Flush()
	err = csvFile.Close()
	if err != nil {
		return
	}
}

// finds all contacts for a person at a specific place
func findPossibleContacts(place string, name string, address string, file []string) bool {
	startTimeMainPerson, endTimeMainPerson, startIndexMainPerson, endIndexMainPerson := findStartEndTimes(place, name, address, file)
	fmt.Printf("Contacts for %v \n", place)
	visitorsList := findVisitors(place, file, name, address)
	if visitorsList == nil {
		return false
	} else {
		var contactsExist bool
		for _, visitor := range visitorsList {
			contactInformation := strings.Split(visitor, ",")
			startTimeContact, endTimeContact, startIndexContact, endIndexContact := findStartEndTimes(place, contactInformation[0], contactInformation[1], file)
			contactsExist = calcContactTime(startTimeMainPerson, endTimeMainPerson, startIndexMainPerson, endIndexMainPerson, startTimeContact, endTimeContact, startIndexContact, endIndexContact, visitor)
		}
		if !contactsExist {
			return false
		}
	}
	fmt.Println("-------------------")
	return true
}

// finds the start and end time of a person at a specific place
func findStartEndTimes(place string, name string, address string, file []string) ([]time.Time, []time.Time, []int, []int) {
	var startTime, endTime []time.Time
	var startIndex, endIndex []int
	firstEntry := true
	lastEntryLogin := false
	for index, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[4], place) == 0 {
			if strings.Compare(row[2], name) == 0 && strings.Compare(row[3], address) == 0 {
				if strings.Contains(row[0], "LOGIN") {
					var bufferStartTime time.Time
					var err error
					bufferStartTime, err = time.Parse(time.RFC3339, row[1])
					//Hier!

					if err != nil {
						panic(err)
					}
					startTime = append(startTime, bufferStartTime)
					startIndex = append(startIndex, index)
					lastEntryLogin = true
					firstEntry = false
				} else {
					if firstEntry {
						fmt.Println(row[1])
						splittedDate := strings.Split(row[1], "T")
						splittedDate2 := strings.Split(splittedDate[1], "+")
						timeZoneRune := []rune(splittedDate2[1])
						splittedDate[1] = string(timeZoneRune[0]) + string(timeZoneRune[1]) + string(timeZoneRune[2]) + string(timeZoneRune[3]) + string(timeZoneRune[4])
						bufferStartTime, err := time.Parse(time.RFC3339, splittedDate[0]+"T00:00:00+"+splittedDate[1])
						if err != nil {
							panic(err)
						}
						startTime = append(startTime, bufferStartTime)
						startIndex = append(startIndex, -1)
						firstEntry = false
					}
					var bufferEndTime time.Time
					var err error
					bufferEndTime, err = time.Parse(time.RFC3339, row[1])
					if err != nil {
						panic(err)
					}
					endTime = append(endTime, bufferEndTime)
					endIndex = append(endIndex, index)
					lastEntryLogin = false
				}
			}
		}
	}

	if lastEntryLogin {
		timeToString := time.Time.String(startTime[0])
		splittedDate := strings.Split(timeToString, " ")
		timeZoneRune := []rune(splittedDate[2])
		splittedDate[2] = string(timeZoneRune[0]) + string(timeZoneRune[1]) + string(timeZoneRune[2]) + ":" + string(timeZoneRune[3]) + string(timeZoneRune[4])
		bufferEndTime, err := time.Parse(time.RFC3339, splittedDate[0]+"T23:59:59"+splittedDate[2])
		if err != nil {
			panic(err)
		}
		endTime = append(endTime, bufferEndTime)
		endIndex = append(endIndex, -1)
	}
	return startTime, endTime, startIndex, endIndex
}

// finds all visitors at a specific place
func findVisitors(place string, file []string, mainPerson string, mainAddress string) []string {
	var visitorsList []string
	mainPersonFull := mainPerson + "," + mainAddress
	for _, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[4], place) == 0 {
			newPerson := row[2] + "," + row[3]
			if strings.Compare(mainPersonFull, newPerson) != 0 {
				newVisitor := true
				for _, visitors := range visitorsList {
					if strings.Compare(visitors, newPerson) == 0 {
						newVisitor = false
					}
				}
				if newVisitor {
					visitorsList = append(visitorsList, newPerson)
				}
			}
		}
	}
	return visitorsList
}

// calculates the contact time of two persons
func calcContactTime(startTimeMainPerson []time.Time, endTimeMainPerson []time.Time, startIndexMainPerson []int, endIndexMainPerson []int, startTimeContact []time.Time, endTimeContact []time.Time, startIndexContact []int, endIndexContact []int, visitor string) bool {
	newContact := true
	numberOfContacts := 0
	var contactTime time.Time
	contactsExist := false

	for newStartIndexContact, oldStartIndexContact := range startIndexContact {
		for newStartIndexMainPerson, oldStartIndexMainPerson := range startIndexMainPerson {
			if oldStartIndexContact == -1 && oldStartIndexMainPerson == -1 {
				if endIndexMainPerson[0] <= endIndexContact[0] {
					newContact = newContactCheck(newContact, visitor)

					numberOfContacts++
					fmt.Printf("Contacts: %v\n", numberOfContacts)
					fmt.Println(endTimeMainPerson[0].Sub(startTimeMainPerson[0]))
					contactTime = contactTime.Add(endTimeMainPerson[0].Sub(startTimeMainPerson[0]))
				} else {
					newContact = newContactCheck(newContact, visitor)

					numberOfContacts++
					fmt.Printf("Contacts: %v\n", numberOfContacts)
					fmt.Println(endTimeContact[0].Sub(startTimeContact[0]))
					contactTime = contactTime.Add(endTimeContact[0].Sub(startTimeContact[0]))
				}
			}
			if oldStartIndexContact < oldStartIndexMainPerson && endIndexContact[newStartIndexContact] > oldStartIndexMainPerson {
				if endIndexContact[newStartIndexContact] <= endIndexMainPerson[newStartIndexMainPerson] {
					newContact = newContactCheck(newContact, visitor)

					//Start Main - end Kontakt
					numberOfContacts++
					fmt.Printf("Contacts: %v\n", numberOfContacts)
					fmt.Println(endTimeContact[newStartIndexContact].Sub(startTimeMainPerson[newStartIndexMainPerson]))
					contactTime = contactTime.Add(endTimeContact[newStartIndexContact].Sub(startTimeMainPerson[newStartIndexMainPerson]))
				} else {
					newContact = newContactCheck(newContact, visitor)

					numberOfContacts++
					fmt.Printf("Contacts: %v\n", numberOfContacts)
					//Start Main - end Main
					fmt.Println(endTimeMainPerson[newStartIndexMainPerson].Sub(startTimeMainPerson[newStartIndexMainPerson]))
					contactTime = contactTime.Add(endTimeMainPerson[newStartIndexMainPerson].Sub(startTimeMainPerson[newStartIndexMainPerson]))

				}
			} else if oldStartIndexContact < endIndexMainPerson[newStartIndexMainPerson] && oldStartIndexContact > oldStartIndexMainPerson {
				if endIndexContact[newStartIndexContact] <= endIndexMainPerson[newStartIndexMainPerson] {
					newContact = newContactCheck(newContact, visitor)
					numberOfContacts++
					fmt.Printf("Contacts: %v\n", numberOfContacts)
					fmt.Println(endTimeContact[newStartIndexContact].Sub(startTimeContact[newStartIndexContact]))
					//Start Kontakt - end kontakt
					contactTime = contactTime.Add(endTimeContact[newStartIndexContact].Sub(startTimeContact[newStartIndexContact]))
				} else {
					newContact = newContactCheck(newContact, visitor)
					numberOfContacts++
					fmt.Printf("Contacts: %v\n", numberOfContacts)
					fmt.Println(endTimeMainPerson[newStartIndexMainPerson].Sub(startTimeContact[newStartIndexContact]))
					//Start Kontakt - end Main
					contactTime = contactTime.Add(endTimeMainPerson[newStartIndexMainPerson].Sub(startTimeContact[newStartIndexContact]))
				}

			}
		}
	}
	if numberOfContacts != 0 {
		time.Time.String(contactTime)
		timeString := strings.Split(time.Time.String(contactTime), " ")
		fmt.Printf("Full contact time: %v\n", timeString[1])
		contactsExist = true
	}
	return contactsExist

}

// checks if a new contact is needed
func newContactCheck(newContact bool, visitor string) bool {
	if newContact {
		fmt.Println("--------------------------")
		fmt.Println(visitor)
		return false
	}
	return false
}
