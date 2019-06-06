package turingMachine

// Struct that defines a Turing machine, it has the following members:
// - states : a hash table, whose keys are the states of the Turing machine,
//            and whose values are booleans that indicate if the state is a
//            final one.
// - tape : the tape for the Turing machine.
// - delta : the transition function for the Turing machine.
// - currentState : the current state of the Turing machine.
// - voidSymbol : the void symbol for the tape of the Turing machine.
// - hasFinished : boolean that indicate if the Turing machine has stopped
//                 (either accepting or rejecting a given input).
type TuringMachine struct {
    states       map[State]bool
    tape         Tape
    delta        Delta
    currentState State
    voidSymbol   rune
    hasFinished  bool
}

// Turing machine constructor. Receives:
// - inputStates : an slice of strings, each string represents a different
//                 state.
// - voidSymbol : the void symbol for the tape
// - initialState : a string that represents the initial state.
// - finalStates : a slice of strings, each string represents a different
//                 final state.
// - transitions : a slice of string slices, each intern slice represents
//                 a transition (so it must have length equal to 5).
// - input : the string that the machine receives as an input
func NewTuringMachine(inputStates []string,
                      voidSymbol rune,
                      initialState string,
                      finalStates []string,
                      transitions [][]string,
                      input string) *TuringMachine {
    turingMachine := new(TuringMachine)
    turingMachine.voidSymbol = voidSymbol
    turingMachine.tape = *NewTape(voidSymbol)
    turingMachine.states = make(map[State]bool)
    for _, s := range inputStates {
        state := *NewState(s)
        turingMachine.states[state] = false
    }
    initial := *NewState(initialState)
    turingMachine.currentState = initial
    for _, s := range finalStates {
        state := *NewState(s)
        turingMachine.states[state] = true
    }
    turingMachine.hasFinished = false
    runeInput := []rune(input)
    for _, symbol := range runeInput {
        turingMachine.tape.Add(symbol)
    }
    turingMachine.tape.Right()
    turingMachine.delta = *NewDelta(transitions)
    return turingMachine
}

// Turing machine function, that performs an execution (it reads a symbol
// from the tape, and alongside the current state, fetches the corresponding
// output from the transition function, then, it writes a symbol in the current
// cell of the tape, moves to a new state and performs a movement over the tape)
func (turingMachine *TuringMachine) Execute() bool {
    symbol := turingMachine.tape.GetSymbol()
    state  := turingMachine.currentState
    argument := Argument{state, symbol}
    result, isDefined := turingMachine.delta.GetResult(argument)
    if !isDefined {
        turingMachine.hasFinished = true
        return false
    }
    turingMachine.currentState = result.state
    turingMachine.tape.WriteSymbol(result.symbol)
    switch result.movement {
    case L:
        turingMachine.tape.Left()
    case N:
        turingMachine.tape.Neutral()
    case R:
        turingMachine.tape.Right()
    }
    isFinal, _ := turingMachine.states[turingMachine.currentState]
    if isFinal {
        turingMachine.hasFinished = true
        return true
    }
    return false
}

// Function that checks if the Turing machine has already stopped.
// Returns a true boolean value if it has stopped, false in other case.
func (turingMachine *TuringMachine) HasFinished() bool {
    return turingMachine.hasFinished
}

// Function that gets the current tape configuration of the Turing
// Machine as a string.
func (turingMachine *TuringMachine) GetConfiguration() string {
    return turingMachine.tape.GetConfiguration(turingMachine.currentState.name)
}
