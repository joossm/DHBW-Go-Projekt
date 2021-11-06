package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func AnlayzeReport() {
	fmt.Println("NOTE: All special characters have to be replaced! For Example: ß -> ss, ä ->ae")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the day you are looking for. (Date Format: YYYY-MM-DD)")
	var text string
	for {
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if validDateFormatReport(text) {
			if fileExists(text) {
				break
			}
			fmt.Println("There are no entries for this day. Please enter another day. (Date Format: YYYY-MM-DD)")
		} else {
			fmt.Println("Wrong Date Format. Please use: YYYY-MM-DD")
		}
	}
	file := readFileToStrings(text)
	lookedForDate := text
	fmt.Println("What would you like to do? Enter the number of your preferred task")
	fmt.Println("1: Look up the places a person visited on this day")
	fmt.Println("2: Extract data for a place as a CSV-file")
	fmt.Println("3: Show contacts of a person")
	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.Compare(text, "1") == 0 {
		fmt.Println("Enter the name of the Person you are looking for:")
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		personName := text
		fmt.Println("Enter the address of the Person you are looking for:")
		fmt.Println("(format: postcode city street house number)")
		fmt.Println("(example: 74821 Mosbach Lohrtalweg 10")
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		personAddress := text
		if personExists(file, personName, personAddress) {
			for _, place := range findPlacesOfPerson(file, personName, personAddress) {
				fmt.Println(place)
			}
		} else {
			fmt.Println("No entry for this Person found.")
		}
	} else if strings.Compare(text, "2") == 0 {
		fmt.Println("Enter the name of the place you want to create a CSV-file for:")
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if placeExists(file, text) {
			createCSV(file, text, lookedForDate)
		} else {
			fmt.Println("Place does not exist")
		}
	} else if strings.Compare(text, "3") == 0 {
		fmt.Println("Enter the name of the Person you are looking for:")
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		personName := text
		fmt.Println("Enter the address of the Person you are looking for:")
		fmt.Println("(format: postcode city street house number)")
		fmt.Println("(example: 74821 Mosbach Lohrtalweg 10)")
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		personAddress := text
		if personExists(file, personName, personAddress) {
			for _, place := range findPlacesOfPerson(file, personName, personAddress) {
				findPossibleContacts(place, personName, personAddress, file)
			}
		} else {
			fmt.Println("No entry for this Person found.")
		}
	} else {
		fmt.Println("No valid choice")
	}
}

func validDateFormatReport(date string) bool {
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

func fileExists(date string) bool {
	_, err := ioutil.ReadFile("src/log/files/" + date + ".txt")
	if err != nil {
		return false
	}
	return true
}

func readFileToStrings(date string) []string {
	file, err := ioutil.ReadFile("src/log/files/" + date + ".txt")
	if err != nil {
		panic(err)
	}
	fileString := strings.Split(string(file), ";")
	if len(fileString) > 0 {
		fileString = fileString[:len(fileString)-1]
	}
	return fileString
}

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

func placeExists(file []string, place string) bool {
	for _, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[4], place) == 0 {
			return true
		}
	}
	return false
}

func createCSV(file []string, place string, lookedForDate string) {
	csvFile, err := os.Create(lookedForDate + "_" + place + ".csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(csvFile)
	for _, element := range file {
		row := strings.Split(element, ", ")
		if strings.Compare(row[4], place) == 0 && strings.Contains(row[0], "LOGIN") {
			var newRow []string
			newRow = append(newRow, row[2], row[3])
			writer.Write(newRow)
		}
	}
	writer.Flush()
	csvFile.Close()
}

func findPossibleContacts(place string, name string, address string, file []string) {
	startTimeMainPerson, endTimeMainPerson, startIndexMainPerson, endIndexMainPerson := findStartEndTimes(place, name, address, file)
	fmt.Printf("Contacts for %v \n", place)
	visitorsList := findVisitors(place, file, name, address)
	if visitorsList == nil {
		fmt.Println("No contact at this place")
	} else {
		for _, visitor := range visitorsList {
			contactInformation := strings.Split(visitor, ",")
			startTimeContact, endTimeContact, startIndexContact, endIndexContact := findStartEndTimes(place, contactInformation[0], contactInformation[1], file)
			calcContactTime(startTimeMainPerson, endTimeMainPerson, startIndexMainPerson, endIndexMainPerson, startTimeContact, endTimeContact, startIndexContact, endIndexContact, visitor)
		}
	}
	fmt.Println("-------------------")
}

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
					bufferStartTime, _ = time.Parse(time.RFC3339, row[1])
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
					bufferEndTime, _ = time.Parse(time.RFC3339, row[1])
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

func findVisitors(place string, file []string, mainPerson string, mainAddress string) []string {
	var visitorsList []string
	mainPersonFull := mainPerson + "," + mainAddress
	fmt.Println(mainPersonFull)
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

func calcContactTime(startTimeMainPerson []time.Time, endTimeMainPerson []time.Time, startIndexMainPerson []int, endIndexMainPerson []int, startTimeContact []time.Time, endTimeContact []time.Time, startIndexContact []int, endIndexContact []int, visitor string) {
	newContact := true
	numberOfContacts := 0
	var contactTime time.Time

	for newStartIndexContact, oldStartIndexContact := range startIndexContact {
		for newStartIndexMainPerson, oldStartIndexMainPerson := range startIndexMainPerson {
			if oldStartIndexContact == -1 && oldStartIndexMainPerson == -1 {
				if endIndexMainPerson[0] <= endIndexContact[0] {
					if newContact {
						newContact = false
						fmt.Println("--------------------------")
						fmt.Println(visitor)
					}
					numberOfContacts++
					fmt.Printf("Contact %v\n", numberOfContacts)
					fmt.Println(endTimeMainPerson[0].Sub(startTimeMainPerson[0]))
					contactTime = contactTime.Add(endTimeMainPerson[0].Sub(startTimeMainPerson[0]))
				} else {
					if newContact {
						newContact = false
						fmt.Println("--------------------------")
						fmt.Println(visitor)
					}
					numberOfContacts++
					fmt.Printf("Contact %v\n", numberOfContacts)
					fmt.Println(endTimeContact[0].Sub(startTimeContact[0]))
					contactTime = contactTime.Add(endTimeContact[0].Sub(startTimeContact[0]))
				}
			}
			if oldStartIndexContact < oldStartIndexMainPerson && endIndexContact[newStartIndexContact] > oldStartIndexMainPerson {
				if endIndexContact[newStartIndexContact] <= endIndexMainPerson[newStartIndexMainPerson] {
					if newContact {
						newContact = false
						fmt.Println("--------------------------")
						fmt.Println(visitor)
					}
					//Start Main - end Kontakt
					numberOfContacts++
					fmt.Printf("Contact %v\n", numberOfContacts)
					fmt.Println(endTimeContact[newStartIndexContact].Sub(startTimeMainPerson[newStartIndexMainPerson]))
					contactTime = contactTime.Add(endTimeContact[newStartIndexContact].Sub(startTimeMainPerson[newStartIndexMainPerson]))
				} else {
					if newContact {
						newContact = false
						fmt.Println("--------------------------")
						fmt.Println(visitor)
					}
					numberOfContacts++
					fmt.Printf("Contact %v\n", numberOfContacts)
					//Start Main - end Main
					fmt.Println(endTimeMainPerson[newStartIndexMainPerson].Sub(startTimeMainPerson[newStartIndexMainPerson]))
					contactTime = contactTime.Add(endTimeMainPerson[newStartIndexMainPerson].Sub(startTimeMainPerson[newStartIndexMainPerson]))

				}
			} else if oldStartIndexContact < endIndexMainPerson[newStartIndexMainPerson] && oldStartIndexContact > oldStartIndexMainPerson {
				if endIndexContact[newStartIndexContact] <= endIndexMainPerson[newStartIndexMainPerson] {
					if newContact {
						newContact = false
						fmt.Println("--------------------------")
						fmt.Println(visitor)
					}
					numberOfContacts++
					fmt.Printf("Contact %v\n", numberOfContacts)
					fmt.Println(endTimeContact[newStartIndexContact].Sub(startTimeContact[newStartIndexContact]))
					//Start Kontakt - end kontakt
					contactTime = contactTime.Add(endTimeContact[newStartIndexContact].Sub(startTimeContact[newStartIndexContact]))
				} else {
					if newContact {
						newContact = false
						fmt.Println("--------------------------")
						fmt.Println(visitor)
					}
					numberOfContacts++
					fmt.Printf("Contact %v\n", numberOfContacts)
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
	}

}
