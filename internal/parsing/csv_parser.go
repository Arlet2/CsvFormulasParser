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

type CsvParseError error
type Csv struct {
	colHeaders []string
	rowHeaders []int
	data [][]string
}

func (csv Csv) PrintWithWriter(writer io.Writer) {
	for index, value := range csv.colHeaders {
		fmt.Fprint(writer, value)
		if index != len(csv.colHeaders)-1 {
			fmt.Fprint(writer, ",")
		}
	}
	fmt.Fprintln(writer)
	for index, line := range csv.data {
		fmt.Fprintf(writer, "%d,", csv.rowHeaders[index])
		for jndex, element := range line {
			fmt.Fprint(writer, element)
			if jndex != len(line)-1 {
				fmt.Fprint(writer, ",")
			}
		}
		fmt.Fprintln(writer)
	}
}

func (csv Csv) Print() {
	csv.PrintWithWriter(os.Stdout)
}

func ParseCsv(file io.Reader) (Csv, error) {

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return Csv{}, errors.New("файл пустой").(CsvParseError)
	}

	headers := strings.Split(scanner.Text(), ",")

	if headers[0] != "" {
		return Csv{}, errors.New("первая ячейка должна быть пустой").(CsvParseError)
	}

	for _, value := range headers {
		// проверяем на целые числа
		_, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return Csv{}, errors.New("названия столбцов не должны быть числами").(CsvParseError)
		}
		// проверяем на числа с плавающей точкой
		_, err = strconv.ParseFloat(value, 64)
		if err == nil {
			return Csv{}, errors.New("названия столбцов не должны быть числами").(CsvParseError)
		}
	}

	rowHeaders := make([]int, 0)
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

		rowHeaders = append(rowHeaders, int(rowIndex))

		data = append(data, values[1:])
	}

	if scanner.Err() != nil {
		return Csv{}, errors.Join(scanner.Err())
	}

	return Csv{colHeaders: headers, rowHeaders: rowHeaders, data: data}, nil
}