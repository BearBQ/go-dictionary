package dictionary

import (
	"errors"
	"fmt"
	. "go-dictionary/input"
	. "go-dictionary/output"
	"strings"
)

//Commands - список команд

// Dictionary Структура основного словаря, типа map[string]string
type Dictionary struct {
	Words map[string]string
}

var CommandsType = map[string]string{
	"add":    "добавить в словарь",
	"get":    "Найти определение слова",
	"delete": "Удалить слово",
	"list":   "Вывести словарь",
}

// NewDictionary создает словарь с уже заданными значениями
func NewDictionary() Dictionary {
	var newDictionary Dictionary
	newDictionary.Words = map[string]string{
		"Голова": "часть тела",
		"Рука":   "конечность",
		"Нога":   "нижняя конечность",
	}
	return newDictionary
}

// ADD добавляет слово и значение в словарь
func Add(incomingDictionary *Dictionary, word string, description string) error {
	if incomingDictionary == nil {
		return errors.New("словаря не существует")
	}
	if word == "" || description == "" {
		return errors.New("слово не может быть пустым")
	}
	_, ok := incomingDictionary.Words[word]
	if ok {
		return fmt.Errorf("слово %s присутствует в словаре", word)
	}

	incomingDictionary.Words[word] = description
	err := PrintStatus("Add", word, "")
	if err != nil {
		fmt.Errorf("Печать не удалась")
	}
	return nil
}

// Get проверяет значение слова
func Get(d *Dictionary, word string) (string, string, error) {

	value, ok := d.Words[word]
	if !ok {
		return "", "", fmt.Errorf("cлово %s отсутствует в словаре", word)
	}

	err := PrintStatus("Get", word, value)
	if err != nil {
		return "", "", fmt.Errorf("ошибка при выводе статуса команды %v", err)
	}
	return word, value, nil
}

// Delete - удаляет значение слова
func Delete(d *Dictionary, word string) error {
	_, ok := d.Words[word]
	if !ok {
		return fmt.Errorf("cлово %s отсутствует в словаре", word)
	} else {
		delete(d.Words, word)
		err := PrintStatus("Delete", word, "")
		if err != nil {
			return fmt.Errorf("ошибка вывода статуса")
		}
		return nil
	}

}

// List выводит в консоль сесь словарь
func List(d *Dictionary) (string, error) {
	if d == nil || len(d.Words) == 0 {
		return "", errors.New("в словаре отсутствуют слова")
	}
	result := ""
	for key, dict := range d.Words {
		result += key + " - " + dict + "\n"
	}
	return result, nil
}

// checkCommand проверяет корректность ввода команды
func CheckCommand(s string) error {
	_, ok := CommandsType[s]
	if !ok {
		return fmt.Errorf("комманды  %s не обнаружено. Пробуй еще раз", s)
	}
	return nil
}

// InputStartCommand принимает ввод команды из терминала и проверяет его
func InputStartCommand(d *Dictionary) (string, error) {
	for {
		command, err := InputCommand()
		if err != nil {
			return "", errors.New("ошибка приема команды")
		}
		err = CheckCommand(command)
		if err == nil {
			return command, nil
		}

		fmt.Println(err)
	}

}

// DitctionaryOperations - логика всей программы. Принимает Dictionary и комманду
func DictionaryOperations(d *Dictionary, command string) error {
	if d == nil {
		return errors.New("словарь не инициализирован")
	}
	commandToLower := strings.ToLower(command)
	switch commandToLower {
	case "add":
		word, discription, err := InputWordsforAdd()
		if err != nil {
			return fmt.Errorf("ошибка : %v", err)
		}
		err = Add(d, word, discription)
		if err != nil {

			return fmt.Errorf("ошибка : %v", err)
		}
	case "get":
		word, err := InputWordForGetDelete()
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ошибка : %v", err)
		}
		key, value, err := Get(d, word)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ошибка : %v", err)
		}
		err = PrintStatus("Get", key, value)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ошибка : %v", err)
		}

	case "delete":
		word, err := InputWordForGetDelete()
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ошибка : %v", err)
		}
		err = Delete(d, word)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ошибка : %v", err)
		}
	case "list":
		resultList, err := List(d)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("ошибка : %v", err)
		}
		fmt.Println(resultList)
	default:

		return fmt.Errorf("неизвестная команда: %s", command)
	}
	return nil
}
