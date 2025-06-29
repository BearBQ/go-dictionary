package output

import (
	"errors"
	"fmt"
)

// PrintStatus Выводит на экран статус результат выполнения операций
// Принимает название команды, ключ и значение
func PrintStatus(s, word, meaning string) error {
	fmt.Println("Запуск PrintStatus")
	switch s {
	case "Add":
		fmt.Printf("Слово %s добавлено в словарь\n", word)
	case "Get":
		fmt.Printf("Значение слова %s - %s\n", word, meaning)
	case "Delete":
		fmt.Printf("Слово %s удалено из словара\n", word)
	default:
		return errors.New("передана неверная команда в модуль печати")
	}

	return nil
}
