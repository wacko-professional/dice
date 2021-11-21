package diceprinter

import (
	"fmt"
	"github.com/wacko-professional/dice"
)

func PrintRoll(sides int, string comment) {
	fmt.Printf("%s: %dn", comment, dice.Roll(sides))
}
