package main

import "fmt"

type Dictionary struct {
	words map[string]string
}

func (d *Dictionary) Add(word, description string) {
	d.words[word] = description
}
func (d *Dictionary) Get(word string) {
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
