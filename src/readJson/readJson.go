package readJson

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
)

// Structure to read a Turing Machine specification from a JSON file.
type JsonTuringMachine struct {
    Estados        []string
    Entrada        []string
    Cinta          []string
    Inicial          string
    Blanco           string
    Finales        []string
    Transiciones [][]string
}

// Reads a Turing Machine specification from a JSON file.
func Read(fileName string) JsonTuringMachine{
    file, fileError := ioutil.ReadFile(fileName)
    if fileError != nil {
        fmt.Println("Error. Could not read from file.")
    }
	turingMachine := JsonTuringMachine{}

	jsonError := json.Unmarshal([]byte(file), &turingMachine)
    if jsonError != nil {
        fmt.Println("Error. Could not read from JSON file.")
    }
    return turingMachine
}
