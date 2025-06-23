package main

import "fmt"

func main() {
	var dictionary1 Dictionary
	dictionary1.words = make(map[string]string)
	dictionary1.words["Голова"] = "Орган для приема пищи"
	fmt.Println(dictionary1.words)

	dictionary1.Add("Глаза", "Два шара")

	dictionary1.Get("Голова")
	dictionary1.Get("Мозг")

	dictionary1.List()
	dictionary1.Delete("Глаза")
	dictionary1.List()
}
