package main

import (
	"fmt"
	"unicode"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (n *NumericInput) Add(i rune) {
	if unicode.IsDigit(i) == true {
		n.input = n.input + string(i)
	}
}

func (n *NumericInput) GetValue() string {
	return n.input
}

func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
