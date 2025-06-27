package dictionary

import (
	"errors"
	"fmt"
)

// Dictionary Структура основного словаря, типа map[string]string
type Dictionary struct {
	words map[string]string
}

// NewDictionary создает словарь с уже заданными значениями
func NewDictionary() Dictionary {
	var newDictionary Dictionary
	newDictionary.words = map[string]string{
		"голова": "часть тела",
		"Рука":   "конечность",
		"Нога":   "нижняя конечность",
	}
	return newDictionary
}

func (d *Dictionary) Add(incomingDictionary *Dictionary, word string, description string) (*Dictionary, error) {
	if incomingDictory == nil {
		return nil, errors.New("словаря не существует")
	}
	if word == "" || description == "" {
		return incomingDictory, errors.New("слово не может быть пустым")
	}
	_, ok := incomingDictory.words[word]
	if ok {
		return incomingDictory, fmt.Errorf("слово %s присутствует в словаре", word)
	}

	incomingDictory.words[word] = description
	return incomingDictory, nil
}

func (d *Dictionary) Get(incword string) (string, string, error) {

	_, ok := d.words[word]
	if !ok {
		fmt.Printf("Слово %s отсутствует в словаре\n", word)
	} else {
		fmt.Printf("Значение слова %s - %s\n", word, d.words[word])
	}
}
func (d *Dictionary) Delete(word string) {
	_, ok := d.words[word]
	if !ok {
		fmt.Printf("Слово %s отсутствует в словаре\n", word)
	} else {
		delete(d.words, word)
		fmt.Printf("Значение слова %s удалено из словаря\n", word)
	}
}
func (d *Dictionary) List() {
	for key, dict := range d.words {
		fmt.Printf("%s - %s\n", key, dict)
	}
}
