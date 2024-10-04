/*
     _____                           _____ _       _               _____                     _
    / ____|                         / ____(_)     | |             |  __ \                   | |
   | |     __ _  ___  ___  __ _ _ _| (___  _ _ __ | |__   ___ _ __| |  | | ___  ___ ___   __| | ___ _ __
   | |    / _` |/ _ \/ __|/ _` | '__\___ \| | '_ \| '_ \ / _ \ '__| |  | |/ _ \/ __/ _ \ / _` |/ _ \ '__|
   | |___| (_| |  __/\__ \ (_| | |  ____) | | |_) | | | |  __/ |  | |__| |  __/ (_| (_) | (_| |  __/ |
    \_____\__,_|\___||___/\__,_|_| |_____/|_| .__/|_| |_|\___|_|  |_____/ \___|\___\___/ \__,_|\___|_|
                                            | |
                                            |_|
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var specials = []string{" ", ".", ",", ";", "!", "?"}
var numberOfTest int = 20
var currentStep int = numberOfTest
var wordList = "words.txt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

type probabilityStruct struct {
	Probability int
	Shift       int
}

// RETURN CONTENT OF WORDLIST
func getFileData(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "Error", err
	}
	return string(data), nil
}

// GET CHAR BY TRANSLATING ALPHABET (KEY2, A->C)
func getChar(startingIndex int) string {
	var currentIndex = startingIndex
	for i := currentStep; i > 0; i-- {
		currentIndex = (currentIndex - 1 + len(alphabet)) % len(alphabet)
	}
	return alphabet[currentIndex]
}

// GENERATE NEW ALPHABET BASE ON CAESAR KEY
func newAlphabet() []string {
	var newAlphabet []string

	for i := range alphabet {
		newAlphabet = append(newAlphabet, getChar(i))
	}
	return newAlphabet
}

// VERIFY IF A CARACTER IS SPECIAL
func isSpecial(char string) bool {
	for i := range specials {
		if char == specials[i] {
			return true
		}
	}
	return false
}

// RETURN INDEX OF REAL ALPHABET
func getIndex(char string) int {
	for i := range alphabet {
		if isSpecial(char) {
			return -1
		}
		if char == alphabet[i] {
			return i
		}
	}
	return -1
}

// RETURN MAX VALUE OF ARR (MAX OCCURENCE)
func getMax(arr []int) int {
	var max = 0
	var maxIndex = 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}
	return maxIndex
}

// SORT PROBABILITY ARRAY (STRONGER TO WEAKEST)
func getProbabilityArr(arr []probabilityStruct) []probabilityStruct {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].Probability < arr[j+1].Probability {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

var textVariants []string
var ShiftList []int
var occurenceList []int

func main() {
	wordsListData, err := getFileData(wordList)
	if err != nil {
		return
	}

	fmt.Println("Enter your hashed text:")
	reader := bufio.NewReader(os.Stdin)
	hashText, err := reader.ReadString('\n')
	fmt.Println("\n\nResult:")

	if err != nil {
		fmt.Println("An error occurred!")
		return
	}

	//FOR EVERY TEST, DO THIS:
	for range numberOfTest {
		newAlph := newAlphabet() //CREATE NEW ALPHABET
		var newText string
		for i := range hashText {
			alphIndex := getIndex(strings.ToLower(string(hashText[i]))) //RETURN INDEX OF hashText[i]
			if alphIndex == -1 {
				newText += string(hashText[i])
			} else {
				newText += newAlph[alphIndex]
			}
		}
		ShiftList = append(ShiftList, currentStep)
		textVariants = append(textVariants, newText)
		currentStep--
	}

	//PERFORM OCCURENCE TEST
	for i := range textVariants {
		var occurence = 0
		splittedText := strings.Split(textVariants[i], " ")
		for w := range splittedText {
			if strings.Contains(string(wordsListData), " "+splittedText[w]) {
				occurence++
			}
		}
		occurenceList = append(occurenceList, occurence)
	}

	var probability_shift_list []probabilityStruct
	if len(occurenceList) == len(ShiftList) {
		for i := range occurenceList {
			probability_shift_list = append(probability_shift_list, probabilityStruct{Probability: occurenceList[i], Shift: ShiftList[i]})
		}
	}

	probArr := getProbabilityArr(probability_shift_list)
	for i := range probArr {
		if probArr[i].Probability >= 20 || i == 0 {
			fmt.Printf("%sProbability: %-5dShift %d%s\n", Red, probArr[i].Probability, probArr[i].Shift, Reset)
		} else {
			fmt.Printf("Probability: %-5dShift %d\n", probArr[i].Probability, probArr[i].Shift)
		}
	}

	strongerShiftIndex := getMax(occurenceList)
	fmt.Printf("%s\nThe best response is [Shift used: %d]:%s \n%s", Green, ShiftList[strongerShiftIndex], Reset, textVariants[strongerShiftIndex])
}
