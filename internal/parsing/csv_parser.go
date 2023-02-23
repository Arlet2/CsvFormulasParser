package parsing

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Csv struct {
	colHeaders map[string]int
	rowHeaders map[string]int
	data [][]string
}

func (csv Csv) PrintWithWriter(writer io.Writer) {
	orderedColHeaders := make([]string, len(csv.colHeaders))
	for key, value := range csv.colHeaders {
		orderedColHeaders[value] = key
	}

	fmt.Fprint(writer, ",")
	for index, value := range orderedColHeaders {
		fmt.Fprint(writer, value)
		if index != len(csv.colHeaders)-1 {
			fmt.Fprint(writer, ",")
		}
	}
	fmt.Fprintln(writer)

	for key, value := range csv.rowHeaders {
		fmt.Fprint(writer, key+",")
		for index, element := range csv.data[value] {
			fmt.Fprint(writer, element)
			if index != len(csv.data[value])-1 {
				fmt.Fprint(writer, ",")
			}
		}
		fmt.Fprintln(writer)
	}
}

func (csv Csv) Print() {
	csv.PrintWithWriter(os.Stdout)
}

type CsvParseError error

func ParseCsv(file io.Reader) (Csv, error) {

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return Csv{}, errors.New("файл пустой").(CsvParseError)
	}

	headers := strings.Split(scanner.Text(), ",")

	if headers[0] != "" {
		return Csv{}, errors.New("первая ячейка должна быть пустой").(CsvParseError)
	}

	colHeaders := make(map[string]int)

	for i := 1; i < len(headers); i++ {
		// проверяем на целые числа
		_, err := strconv.ParseInt(headers[i], 10, 64)
		if err == nil {
			return Csv{}, errors.New("названия столбцов не должны быть числами").(CsvParseError)
		}
		// проверяем на числа с плавающей точкой
		_, err = strconv.ParseFloat(headers[i], 64)
		if err == nil {
			return Csv{}, errors.New("названия столбцов не должны быть числами").(CsvParseError)
		}
		colHeaders[headers[i]] = i-1
	}

	rowHeaders := make(map[string]int, 0)
	data := make([][]string, 0)

	var values []string
	for scanner.Scan() {
		values = strings.Split(scanner.Text(), ",")
		rowIndex, err := strconv.ParseInt(values[0], 10, 64)

		if err != nil {
			return Csv{}, errors.New("номер строки должен быть числом").(CsvParseError)
		}

		if rowIndex <= 0 {
			return Csv{}, errors.New("номер строки должен быть положительным числом").(CsvParseError)
		}

		rowHeaders[values[0]] = len(data)

		data = append(data, values[1:])
	}

	if scanner.Err() != nil {
		return Csv{}, errors.Join(scanner.Err())
	}

	return Csv{colHeaders: colHeaders, rowHeaders: rowHeaders, data: data}, nil
}