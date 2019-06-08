package turingMachine

// Constants that define three possible movenets for a Turing machine tape.
type Movement string
const (
    L Movement = "Left"
    N Movement = "Neutral"
    R Movement = "Right"
)

// Argument type for the Turing machine transition function.
// Contains the current state and the current symbol that is read by the
// Turing machine head.
type Argument struct {
    state  State
    symbol rune
}

// Result type for the Turing machine transition function.
// Contains the new state, the symbol that will be written in the tape and the
// movement that the Turing machine will perform.
type Result struct{
    state    State
    symbol   rune
    movement Movement
}

// Delta type that defines the Turing machine transition function.
// Contains a hash table that assigns to an specific argument (key),
// an specific output (value).
type Delta struct {
    transition map[Argument]Result
}

// Delta type constructor. Receives a slice of string slices, which contains
// the necessary transitions for the transition function.
func NewDelta(transitions [][]string) *Delta {
    delta := new(Delta)
    delta.transition = make(map[Argument]Result)
    delta.build(transitions)
    return delta
}

// Adds a new transition to the transition function.
func (delta *Delta) addTransition(argument Argument, result Result){
    delta.transition[argument] = result
}

// Gets the output of the transition function, given an specific argument.
func (delta *Delta) GetResult(argument Argument) (Result, bool) {
    result, defined := delta.transition[argument]
    if defined {
        return result, true
    } else {
        return result, false
    }
}

// Builds the transition function, given a slice of string slices.
func (delta *Delta) build(transitions [][]string ) {
    for _, transition := range transitions {
        if (len(transition) % 5 == 0) {
            argumentState := *NewState(transition[0])
            runeArgumentSlice := []rune(transition[1])
            argumentSymbol := runeArgumentSlice[0]
            argument := Argument{argumentState, argumentSymbol}
            resultState := *NewState(transition[2])
            runeResultSlice := []rune(transition[3])
            resultSymbol := runeResultSlice[0]
            var result Result
            switch transition[4] {
            case "L":
                result = Result{resultState, resultSymbol, L}
            case "N":
                result = Result{resultState, resultSymbol, N}
            case "R":
                result = Result{resultState, resultSymbol, R}
            default:
                result = Result{resultState, resultSymbol, N}
            }
            delta.addTransition(argument, result)
        }
    }
}
