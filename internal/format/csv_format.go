package format

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Csv struct {
	ColHeaders map[string]int
	RowHeaders map[string]int
	Data       [][]string
	// для быстрого поиска
	orderedColHeaders []string
	orderedRowHeaders []string
}

func (csv Csv) GetLinkByIndexes(col int, row int) (string, error) {
	if csv.orderedColHeaders == nil {
		csv.orderedColHeaders = getOrderedHeaders(csv.ColHeaders)
	}
	if csv.orderedRowHeaders == nil {
		csv.orderedRowHeaders = getOrderedHeaders(csv.RowHeaders)
	}

	if col >= len(csv.ColHeaders) {
		return "", errors.New("некорректный номер столбца")
	}
	if row >= len(csv.RowHeaders) {
		return "", errors.New("некорректный номер строки")
	}

	return csv.orderedColHeaders[col] + csv.orderedRowHeaders[row], nil
}

func getOrderedHeaders(headers map[string]int) []string {
	orderedHeaders := make([]string, len(headers))

	for key, value := range headers {
		orderedHeaders[value] = key
	}
	return orderedHeaders
}

func (csv Csv) PrintWithWriter(writer io.Writer) {
	fmt.Fprint(writer, ",") // пустая строка
	// порядок ключей в map не детерминирован
	for index, value := range getOrderedHeaders(csv.ColHeaders) {
		fmt.Fprint(writer, value)
		if index != len(csv.ColHeaders)-1 {
			fmt.Fprint(writer, ",")
		}
	}
	fmt.Fprintln(writer)

	orderedRowHeaders := getOrderedHeaders(csv.RowHeaders)
	for index, line := range csv.Data {
		fmt.Fprint(writer, orderedRowHeaders[index]+",")

		for jndex, element := range line {
			fmt.Fprint(writer, element)
			if jndex != len(csv.Data[index])-1 {
				fmt.Fprint(writer, ",")
			}
		}
		fmt.Fprintln(writer)
	}

}

func (csv Csv) Print() {
	csv.PrintWithWriter(os.Stdout)
}
