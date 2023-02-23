package main

import (
	"fmt"
	"os"
	"test_task/internal/parsing"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Введите название .csv файла, который нужно обработать")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Слишком много аргументов. Укажите ровно один .csv файл")
		return
	}
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Не удалось открыть файл. Ошибка: "+err.Error())
		return
	}
	defer file.Close()

	csv, err := parsing.ParseCsv(file)

	if err != nil {
		fmt.Println("Ошибка при парсинге CSV: "+err.Error())
		return
	}
	csv.Print()
	
}