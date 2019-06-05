package turingMachine

type TuringMachine struct {
    states       map[State]bool
    tape         Tape
    delta        Delta
    currentState State
    voidSymbol   rune
    hasFinished  bool
}

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
        turingMachine.tape.AddEnd(symbol)
    }
    turingMachine.tape.Right()

    turingMachine.delta = *NewDelta(transitions)
    return turingMachine
}

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

func (turingMachine *TuringMachine) HasFinished() bool {
    return turingMachine.hasFinished
}

func (turingMachine *TuringMachine) GetConfiguration() string {
    return turingMachine.tape.GetConfiguration(turingMachine.currentState.name)
}
