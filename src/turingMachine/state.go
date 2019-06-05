package turingMachine

// State type for the Turing machine.
type State struct {
    name string
}

// State constructor.
func NewState(name string) *State {
    state := new(State)
    state.name = name
    return state
}

// Checks if two states are equal.
func (this *State) Equals(state *State) bool {
    return this.name == state.name
}
