package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const csvFileName = "problemes.csv"
const jsonFileName = "data.json"

// Problemes struct which contains an array of problemes
type Problemes struct {
	Probleme []Probleme `json:"probleme"`
}

// Probleme struct which contains a question and a reponse
type Probleme struct {
	Question string `json:"question"`
	Reponse  int    `json:"reponse"`
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	byteValue := readJSONFile(jsonFileName)

	var problemes Problemes
	json.Unmarshal(byteValue, &problemes)

	createCSVFile(problemes)
	readCSVFile()

}
func createCSVFile(problemes Problemes) (resultat Problemes) {

	csvFile, err := os.Create(csvFileName)

	checkError(err)

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"Question", "Reponse"})
	for _, probleme := range problemes.Probleme {
		writer.Write([]string{probleme.Question, strconv.Itoa(probleme.Reponse)})
		writer.Error()
	}
	return problemes
}
func readCSVFile() {
	var input string
	var score int
	openCSVFile, err := ioutil.ReadFile(csvFileName)
	checkError(err)

	fmt.Println("Opération ouverture " + csvFileName + " reussi")

	reader := csv.NewReader(strings.NewReader(string(openCSVFile)))
	// on lit la premiere ligne
	reader.Read()
	for i := 1; i < compterLigneFichier(csvFileName); i++ {
		monProbleme, err := reader.Read()
		checkError(err)

		fmt.Println(monProbleme[0])
		fmt.Println("Enter text: ")

		go timer(i, &i)

		fmt.Scanln(&input)
		fmt.Println("Le resultat est " + monProbleme[1])

		if monProbleme[1] == input {
			score++
		}
	}
	fmt.Println("Votre score est de " + strconv.Itoa(score) + "/" + strconv.Itoa(compterLigneFichier(csvFileName)-1))
}
func compterLigneFichier(csvFileName string) int {
	file, _ := os.Open(csvFileName)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}
func timer(debutIndex int, finIndex *int) {
	timerQuestion := time.NewTimer(5 * time.Second)
	<-timerQuestion.C
	if debutIndex != *finIndex {
		timerQuestion.Stop()
	} else {
		fmt.Println(timerQuestion)
		os.Exit(1)
	}
}
func readJSONFile(jsonFileName string) []byte {
	jsonFile, error := os.Open(jsonFileName)

	checkError(error)
	fmt.Println("Opération ouverture " + jsonFileName + " reussi")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}
