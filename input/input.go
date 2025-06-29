package input

import (
	"fmt"
	"os"
	"strings"
)

// InputCommand получает значение команды из командной строки
func InputCommand() (string, error) {
	var command string
	fmt.Println("Введите команду (Add, Get, Delete, List):")
	fmt.Fscan(os.Stdin, &command)

	commandToLower := strings.ToLower(command)

	return commandToLower, nil
}

// InputWordsForAdd приинмает данные слова и значения из командной строки для комманды Add
func InputWordsforAdd() (string, string, error) {
	var word, description string
	for {
		fmt.Println("Введите слово")
		fmt.Fscan(os.Stdin, &word)
		fmt.Println("Введите определение")
		fmt.Fscan(os.Stdin, &description)
		if word != "" && description != "" {
			return word, description, nil
		}
		fmt.Println("Слово и определение не заполнены")
	}
}

// InputWordforDelete принимает данные слова для комманд Get и Delete
func InputWordForGetDelete() (string, error) {
	var word string
	for {
		fmt.Println("Введите слово")
		fmt.Fscan(os.Stdin, &word)
		if word != "" {
			return word, nil
		}
		fmt.Println("Слово не может быть пустым")

	}
}
