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
	words map[string]string
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
	newDictionary.words = map[string]string{
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
	_, ok := incomingDictionary.words[word]
	if ok {
		return fmt.Errorf("слово %s присутствует в словаре", word)
	}

	incomingDictionary.words[word] = description
	err := PrintStatus("Add", word, "")
	if err != nil {
		fmt.Errorf("Печать не удалась")
	}
	return nil
}

// Get проверяет значение слова
func Get(d *Dictionary, word string) error {

	value, ok := d.words[word]
	if !ok {
		return fmt.Errorf("cлово %s отсутствует в словаре", word)
	}

	err := PrintStatus("Get", word, value)
	if err != nil {
		return fmt.Errorf("ошибка при выводе статуса команды %v", err)
	}
	return nil
}

func Delete(d *Dictionary, word string) error {
	_, ok := d.words[word]
	if !ok {
		return fmt.Errorf("cлово %s отсутствует в словаре", word)
	} else {
		delete(d.words, word)
		err := PrintStatus("Delete", word, "")
		if err != nil {
			return fmt.Errorf("ошибка вывода статуса")
		}
		return nil
	}

}

// List выводит в консоль сесь словарь
func List(d *Dictionary) error {
	if d == nil || len(d.words) == 0 {
		return errors.New("в словаре отсутствуют слова")
	}
	for key, dict := range d.words {
		fmt.Printf("%s - %s\n", key, dict)
	}
	return nil
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
			return fmt.Errorf("Ошибка", err)
		}
		err = Add(d, word, discription)
		if err != nil {

			return fmt.Errorf("Ошибка", err)
		}
	case "get":
		word, err := InputWordForGetDelete()
		if err != nil {
			return fmt.Errorf("Ошибка", err)
		}
		err = Get(d, word)
		if err != nil {
			return fmt.Errorf("Ошибка", err)
		}
	case "delete":
		word, err := InputWordForGetDelete()
		if err != nil {
			return fmt.Errorf("Ошибка", err)
		}
		err = Delete(d, word)
		if err != nil {
			return fmt.Errorf("Ошибка", err)
		}
	case "list":
		err := List(d)
		if err != nil {
			return fmt.Errorf("Ошибка", err)
		}
	default:
		return fmt.Errorf("неизвестная команда: %s", command)
	}
	return nil
}
