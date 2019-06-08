package main

import (
    "fmt"
    "../turingMachine"
    "../readJson"
)

func main()  {
    // The string for the JSON file
    var jsonFile string
    // The input for the Turing machine
    var input string

    fmt.Print("Archivo con la descripcion de la MT : ")
    fmt.Scanln(&jsonFile)
    fmt.Print("Inserte la cadena de entrada : ")
    fmt.Scanln(&input)

    // Gets a processed Turing machine specification.
    tmSpecification := readJson.Read(jsonFile)

    states := tmSpecification.Estados
    void := []rune(tmSpecification.Blanco)[0]
    initial := tmSpecification.Inicial
    transitions := tmSpecification.Transiciones
    finals := tmSpecification.Finales

    // Instantiates a new Turing machine.
    tm := turingMachine.NewTuringMachine(states, void, initial, finals,
                                         transitions, input)
    // Runs the Turing machine.
    var accepted bool
    fmt.Println(tm.GetConfiguration())

    for !tm.HasFinished() {
        accepted = tm.Execute()
        fmt.Println(tm.GetConfiguration())
    }
    if accepted {
        fmt.Println("La cadena fue aceptada")
    }
}
