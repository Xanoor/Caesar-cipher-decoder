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

var wordList = "words.txt" // wordlist to be used
var numberOfTest int = 26  // number of tests (max length of the alphabet used)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var specials string = " .,?!:;()[]{}«»\"'’-_/\\$€£¥₣₽+−×÷=≠><≥≤%^√π∞@#*&^~|¿¡§¶éèëüçâêîôû"

// Don't modify this !
var currentStep int = numberOfTest
var textVariants []string
var ShiftList []int
var occurenceList []int

// Color codes
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

// GET CHARACTER BY TRANSLATING ALPHABET (KEY2, A->C)
func getChar(startingIndex int) string {
	var currentIndex = startingIndex
	for i := currentStep; i > 0; i-- {
		currentIndex = (currentIndex - 1 + len(alphabet)) % len(alphabet)
	}
	return alphabet[currentIndex]
}

// GENERATE A NEW ALPHABET BASED ON THE CAESAR KEY
func newAlphabet() []string {
	var newAlphabet []string

	for i := range alphabet {
		newAlphabet = append(newAlphabet, getChar(i))
	}
	return newAlphabet
}

// RETURN INDEX OF THE REAL ALPHABET
func getIndex(char string) int {
	for i := range alphabet {
		if strings.Contains(alphabet[i], specials) {
			return -1
		}
		if char == alphabet[i] {
			return i
		}
	}
	return -1
}

// RETURN MAX VALUE OF AN ARRAY (MAX OCCURRENCE)
func getMax(arr []int) (int, int) {
	var max = 0
	var maxIndex = 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}
	return maxIndex, max
}

// SORT PROBABILITY ARRAY (FROM STRONGEST TO WEAKEST)
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

// VERIFY IF A FILE WITH THE GIVEN PATH EXISTS
func isFileValid(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func main() {
	ASCII, err := getFileData("ASCIIart.txt")
	if err == nil {
		fmt.Println(ASCII)
	}

	wordsListData, err := getFileData(wordList)
	if err != nil {
		return
	}

	fmt.Println(Green + "Enter your encrypted text: [raw text / file name (txt)]" + Reset)
	reader := bufio.NewReader(os.Stdin)
	encryptedText, err := reader.ReadString('\n')
	encryptedText = strings.TrimSpace(encryptedText)

	// If a file with encryptedText name exist, read content
	if isFileValid(encryptedText) && len(encryptedText) < 255 {
		encryptedText, err = getFileData(encryptedText)
	}

	fmt.Println("\nResult:")

	if err != nil {
		fmt.Println("An error occurred!")
		return
	}

	// FOR EVERY TEST, DO THIS:
	for range numberOfTest {
		newAlph := newAlphabet() // CREATE A NEW ALPHABET
		var newText string
		for i := range encryptedText {
			currentChar := string(encryptedText[i])
			alphIndex := getIndex(strings.ToLower(currentChar)) // RETURN INDEX OF encryptedText[i]
			if alphIndex == -1 {
				newText += string(encryptedText[i])
			} else {
				if currentChar == strings.ToUpper(currentChar) { // Detect uppercase
					newText += strings.ToUpper(newAlph[alphIndex])
				} else {
					newText += newAlph[alphIndex]
				}
			}
		}
		ShiftList = append(ShiftList, currentStep)
		textVariants = append(textVariants, newText)
		currentStep--
	}

	// PERFORM OCCURRENCE TEST
	for i := range textVariants {
		var occurence = 0
		splittedText := strings.Split(textVariants[i], " ")
		for w := range splittedText {
			if strings.Contains(string(wordsListData), " "+strings.ToLower(splittedText[w])+" ") {
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
	var maxProb int
	//Show shift by probability
	for i := range probArr {
		if i == 0 {
			maxProb = probArr[i].Probability
		}
		if probArr[i].Probability >= 20 || i == 0 || probArr[i].Probability == maxProb {
			fmt.Printf("%sProbability: %-5dShift %d%s\n", Red, probArr[i].Probability, probArr[i].Shift, Reset)
		} else {
			fmt.Printf("Probability: %-5dShift %d\n", probArr[i].Probability, probArr[i].Shift)
		}
	}

	strongerShiftIndex, probValue := getMax(occurenceList)
	fmt.Printf("%s\nThe best response is [Shift used: %d, Probability: %d]:%s \n%s", Green, ShiftList[strongerShiftIndex], probValue, Reset, textVariants[strongerShiftIndex])
	fmt.Printf("%s\n----------------------END----------------------%s", Green, Reset)
}
