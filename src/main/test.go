package main

import (
    "fmt"
    "../turingMachine"
)

func main()  {
    states := []string{"q0", "q1", "qf"}
    void := '⊔'
    initial := "q0"
    finals := []string{"qf"}
    transitions := [][]string{{"q0", "0", "q1", "1", "r"},
                              {"q1", "1", "q0", "0", "r"},
                              {"q1", "⊔", "qf", "⊔", "r"}}

    tm := turingMachine.NewTuringMachine(states, void, initial, finals,
                                         transitions, "010")
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
