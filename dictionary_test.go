package main

import "testing"

func TestAdd(t *testing.T) {
	d := Dictionary{
		words: make(map[string]string),
	}
	word, meaning := "Рука", "Одна из конечностей"
	d.Add(word, meaning)
	meanResult, availability := d.words[word]
	if !availability {
		t.Errorf("Add(%s, %s) не добавил слово в словарь", word, meanResult)
	} else if meanResult != meaning {
		t.Errorf("Add(%s, %s) не корректно добавил слово в словарь", word, meanResult)
	}
}

func TestDelete(t *testing.T) {
	word, meaning := "Рука", "Одна из конечностей"
	d := Dictionary{
		words: map[string]string{word: meaning},
	}
	d.Delete(word)
	_, availability := d.words[word]
	if availability == true {
		t.Errorf("Deleate(%s) не удалило слово %s", word, word)
	}
}
