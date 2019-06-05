package turingMachine

type Cell struct {
    previous *Cell
    symbol   rune
    next     *Cell
}

type Tape struct {
    voidSymbol rune
    beginning  *Cell
    current    *Cell
    end        *Cell
}

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

func (tape *Tape) AddBeginning(symbol rune) {
    newCell := new(Cell)
    newCell.symbol = symbol
    newCell.next = tape.beginning
    tape.beginning.previous = newCell
    tape.beginning = newCell
}

func (tape *Tape) AddEnd(symbol rune) {
    newCell := new(Cell)
    newCell.symbol = symbol
    newCell.previous = tape.end
    tape.end.next = newCell
    tape.end = newCell
}

func (tape *Tape) addCell(symbol rune) {
    newCell := new(Cell)
    newCell.symbol = symbol
    newCell.previous = tape.end
    tape.end.next = newCell
    tape.end = newCell
}

func (tape *Tape) WriteSymbol(symbol rune) {
    tape.current.symbol = symbol
}

func (tape *Tape) GetSymbol() rune {
    return tape.current.symbol
}

func (tape *Tape) Right() {
    if tape.current.next != nil {
        tape.current = tape.current.next
    } else {
        tape.AddEnd(tape.voidSymbol)
        tape.current = tape.current.next
    }
}

func (tape *Tape) Left() {
    if tape.current.previous != nil {
        tape.current = tape.current.previous
    } else {
        tape.AddBeginning(tape.voidSymbol)
        tape.current = tape.current.previous
    }
}

func (tape *Tape) Neutral() {
    tape.current = tape.current
}

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
