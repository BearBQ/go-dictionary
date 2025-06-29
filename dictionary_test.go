package main

import (
	. "go-dictionary/dictionary"
	"testing"
)

func TestAdd(t *testing.T) {
	d := Dictionary{
		Words: make(map[string]string),
	}
	word, meaning := "Рука", "Одна из конечностей"
	err := Add(&d, word, meaning)
	if err != nil {
		t.Errorf("Ошибка %v", err)
		_, availability := d.Words[word]
		if !availability {
			t.Errorf("Add(%s) не добавил слово в словарь", word)
		}
	}
}

func TestDelete(t *testing.T) {
	word, meaning := "Рука", "Одна из конечностей"
	d := Dictionary{
		Words: map[string]string{word: meaning},
	}
	err := Delete(&d, word)
	// d.Delete(word)
	if err != nil {
		t.Errorf("Ошибка %v", err)
	}
	_, availability := d.Words[word]
	if availability == true {
		t.Errorf("Deleate(%s) не удалило слово %s", word, word)
	}
}

func TestList(t *testing.T) {
	word, meaning := "Рука", "Одна из конечностей"
	d := Dictionary{
		Words: map[string]string{word: meaning},
	}
	expected := ("Рука - Одна из конечностей\n")
	result, _ := List(&d)
	if expected != result {
		t.Errorf("некорректно работает функция list")
	}

}

func TestGet(t *testing.T) {
	word, meaning := "Рука", "Одна из конечностей"
	d := Dictionary{
		Words: map[string]string{word: meaning},
	}
	resultKey, resultValue, _ := Get(&d, word)
	if resultKey != word || resultValue != meaning {
		t.Errorf("неверно работает Get")
	}

}
