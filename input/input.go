package input

import (
	"fmt"
	"os"
)

// InputCommand получает значение команды из командной строки
func InputCommand() (string, error) {
	var command string
	fmt.Println("Введите команду (Add, Get, Delete, List):")
	fmt.Fscan(os.Stdin, &command)
	return command, nil
}

// InputWordsForAdd приинмает данные слова и значения из командной строки для комманды Add
func InputWordsforAdd() (string, string, error) {
	var word, description string
	for {
		fmt.Println("Введите слово")
		fmt.Fscan(os.Stdin, &word)
		if word == "" {
			fmt.Println("Слово не может быть пустым")
		} else {
			fmt.Println("Введите слово")
			fmt.Fscan(os.Stdin, description)
			if word == "" {
				fmt.Println("Слово не может быть пустым")
			} else {
				return word, description, nil
			}
		}
	}
}

// InputWordforDelete принимает данные слова для комманд Get и Delete
func InputWordForGetDelete() (string, error) {
	var word string
	for {
		fmt.Println("Введите слово")

		fmt.Fscan(os.Stdin, &word)
		if word == "" {
			fmt.Println("Слово не может быть пустым")
		} else {
			return word, nil
		}
	}
}
