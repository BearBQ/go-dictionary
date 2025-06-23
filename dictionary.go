package main

import "fmt"

type dictionary struct {
	words (map[string]string)
}

func (d *dictionary) add(word, description string) {
	d.words[word] = description
}
func (d *dictionary) get(word string) {
	_, ok := d.words[word]
	if !ok {
		fmt.Printf("Слово %s отсутствует в словаре", word)
	} else {
		fmt.Printf("Значение слова %s - %s", word, d.words[word])
	}
}
func (d *dictionary) delete(word string) {
	_, ok := d.words[word]
	if !ok {
		fmt.Printf("Слово %s отсутствует в словаре", word)
	} else {
		delete(d.words, word)
	}
}
func (d *dictionary) list() {
	fmt.Println(d.words)
}
