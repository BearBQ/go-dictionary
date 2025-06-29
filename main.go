package main

import (
	"fmt"
	. "go-dictionary/dictionary"
)

func main() {

	var MainDictionary Dictionary = NewDictionary()

	for {
		cycleOfProgram((&MainDictionary))
	}
}

func cycleOfProgram(d *Dictionary) error {

	command, err := InputStartCommand(d)
	if err != nil {
		return fmt.Errorf("ошибка ввода: %v\n", err)

	}
	err = DictionaryOperations(d, command)
	if err != nil {
		return fmt.Errorf("ошибка ввода: %v\n", err)
	}
	return nil
}
