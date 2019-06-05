package turingMachine

// Cell type for the Turing machine tape.
type Cell struct {
    previous *Cell
    symbol   rune
    next     *Cell
}

// Tape type for the Turing machine.
type Tape struct {
    voidSymbol rune
    beginning  *Cell
    current    *Cell
    end        *Cell
}

// Tape constructor. Receives the void symbol for the tape.
func NewTape(voidSymbol rune) *Tape {
    tape := new(Tape)
    cell := new(Cell)
    cell.symbol = voidSymbol
    tape.voidSymbol = voidSymbol
    tape.beginning = cell
    tape.current = cell
    tape.end = cell
    return tape
}

// Function that adds a new cell at the beginning of the tape.
func (tape *Tape) addBeginning(symbol rune) {
    newCell := new(Cell)
    newCell.symbol = symbol
    newCell.next = tape.beginning
    tape.beginning.previous = newCell
    tape.beginning = newCell
}

// Function that adds a new cell at the end of the tape.
func (tape *Tape) addEnd(symbol rune) {
    newCell := new(Cell)
    newCell.symbol = symbol
    newCell.previous = tape.end
    tape.end.next = newCell
    tape.end = newCell
}

// Function that adds a new cell to the tape.
func (tape *Tape) addCell(symbol rune) {
    newCell := new(Cell)
    newCell.symbol = symbol
    newCell.previous = tape.end
    tape.end.next = newCell
    tape.end = newCell
}

// Function that writes a symbol on the current cell of the tape.
func (tape *Tape) WriteSymbol(symbol rune) {
    tape.current.symbol = symbol
}

// Function that reads a symbol from the current cell of the tape.
func (tape *Tape) GetSymbol() rune {
    return tape.current.symbol
}

// Function that changes the current cell of the tape for the one that is
// to the right of it.
func (tape *Tape) Right() {
    if tape.current.next != nil {
        tape.current = tape.current.next
    } else {
        tape.addEnd(tape.voidSymbol)
        tape.current = tape.current.next
    }
}

// Function that changes the current cell of the tape for the one that is
// to the left of it.
func (tape *Tape) Left() {
    if tape.current.previous != nil {
        tape.current = tape.current.previous
    } else {
        tape.addBeginning(tape.voidSymbol)
        tape.current = tape.current.previous
    }
}

// Function that defines a neutral movement (does nothing).
func (tape *Tape) Neutral() {
    tape.current = tape.current
}

// Function that gets the current tape configuration as a string.
func (tape *Tape) GetConfiguration(currentState string) string {
    str := ""
    it := tape.beginning
    for it != nil {
        if it == tape.current {
            str += "|" + currentState + "|"
        }
        if !((it.symbol == tape.voidSymbol) && (it == tape.beginning)) {
            str += string(it.symbol)
        }
        it = it.next
    }
    return str
}
