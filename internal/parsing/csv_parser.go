package parsing

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"test_task/internal/format"
)

type CsvParseError error

func ParseCsv(file io.Reader) (format.Csv, error) {

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return format.Csv{}, errors.New("файл пустой").(CsvParseError)
	}

	headers := strings.Split(scanner.Text(), ",")

	if headers[0] != "" {
		return format.Csv{}, errors.New("первая ячейка должна быть пустой").(CsvParseError)
	}

	colHeaders := make(map[string]int)

	for i := 1; i < len(headers); i++ {
		// проверяем на целые числа
		_, err := strconv.ParseInt(headers[i], 10, 64)
		if err == nil {
			return format.Csv{}, errors.New("названия столбцов не должны быть числами").(CsvParseError)
		}
		// проверяем на числа с плавающей точкой
		_, err = strconv.ParseFloat(headers[i], 64)
		if err == nil {
			return format.Csv{}, errors.New("названия столбцов не должны быть числами").(CsvParseError)
		}
		colHeaders[headers[i]] = i - 1
	}

	rowHeaders := make(map[string]int, 0)
	data := make([][]string, 0)

	var values []string
	for scanner.Scan() {
		values = strings.Split(scanner.Text(), ",")

		if len(values) != len(colHeaders)+1 {
			return format.Csv{}, errors.New("количество ячеек в строке должно быть равно количеству именованных заголовков").(CsvParseError)
		}

		rowIndex, err := strconv.ParseInt(values[0], 10, 64)

		if err != nil {
			return format.Csv{}, errors.New("номер строки должен быть числом").(CsvParseError)
		}

		if rowIndex <= 0 {
			return format.Csv{}, errors.New("номер строки должен быть положительным числом").(CsvParseError)
		}

		rowHeaders[values[0]] = len(data)

		for _, value := range values[1:] {
			_, err := strconv.ParseInt(value, 10, 32)
			// если это не формула и не целое число, то это недопустимое значение
			if !IsFormula(value) && err != nil {
				return format.Csv{}, errors.New("значения ячеек должны быть формулами или целыми числами").(CsvParseError)
			}
		}

		data = append(data, values[1:])
	}

	if scanner.Err() != nil {
		return format.Csv{}, errors.Join(scanner.Err())
	}

	return format.Csv{ColHeaders: colHeaders, RowHeaders: rowHeaders, Data: data}, nil
}
