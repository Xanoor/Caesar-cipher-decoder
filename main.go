package main

import (
	"fmt"
	"os"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var specials = []string{" ", ".", ",", ";", "!", "?"}
var numberOfTest int = 100
var currentStep int = numberOfTest

var hashText string = "Ot kdgsototm znk yuioukiutusoi xgsoloigzouty ul mruhgrofgzout, utk sayz giqtucrkjmk znk vgxgjud zngz cnork oz ngy lgiorozgzkj atvxkikjktzkj kiutusoi mxuczn gtj iarzaxgr kdingtmk, oz ngy gryu kdgikxhgzkj otkwagrozoky huzn coznot gtj hkzckkt tgzouty. Znk xgvoj lruc ul igvozgr gtj otluxsgzout, cnork hktkloiogr lux ikxzgot ykizuxy, ulzkt rkgbky sgxmotgrofkj iussatozoky laxznkx joyktlxgtinoykj, ixkgzotm g ieirk ul vubkxze zngz oy jolloiarz zu kyigvk. Iutykwaktzre, g nuroyzoi gvvxugin zngz ktiusvgyyky lgox zxgjk vxgizoiky, yayzgotghrk jkbkruvsktz, gtj kwaozghrk xkyuaxik joyzxohazout oy kyyktzogr zu ngxtkyy znk hktklozy ul mruhgrofgzout coznuaz ygixoloiotm znk ckrr-hkotm ul znk suyz bartkxghrk vuvargzouty."

func getFileData(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "Error", err
	}
	return string(data), nil
}

func getChar(startingIndex int) string {
	var currentIndex = startingIndex
	for i := currentStep; i > 0; i-- {
		currentIndex = (currentIndex - 1 + len(alphabet)) % len(alphabet)
	}
	return alphabet[currentIndex]
}

func newAlphabet() []string {
	var newAlphabet []string

	for i := range alphabet {
		newAlphabet = append(newAlphabet, getChar(i))
	}
	return newAlphabet
}

func isSpecial(char string) bool {
	for i := range specials {
		if char == specials[i] {
			return true
		}
	}
	return false
}

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

func getPotentials(arr []int) []string {
	var maxVal = getMax(arr)
	var potentials []string
	for i := range arr {
		if occurenceList[i] > maxVal-1 {
			potentials = append(potentials, textVariants[i])
		}
	}
	return potentials
}

var textVariants []string
var occurenceList []int

func main() {
	wordsListData, err := getFileData("words.txt")
	if err != nil {
		return
	}

	for range numberOfTest {
		newAlph := newAlphabet()
		var newText string
		for i := range hashText {
			alphIndex := getIndex(strings.ToLower(string(hashText[i])))
			if alphIndex == -1 {
				newText += string(hashText[i])
			} else {
				newText += newAlph[alphIndex]
			}
		}
		currentStep--
		textVariants = append(textVariants, newText)
	}

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

	fmt.Println("The good sentence is \n" + textVariants[getMax(occurenceList)])

}
