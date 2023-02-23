package format

import (
	"fmt"
	"io"
	"os"
)

type Csv struct {
	ColHeaders map[string]int
	RowHeaders map[string]int
	Data       [][]string
}

func (csv Csv) PrintWithWriter(writer io.Writer) {
	// порядок ключей в map не детерминирован
	orderedColHeaders := make([]string, len(csv.ColHeaders))
	for key, value := range csv.ColHeaders {
		orderedColHeaders[value] = key
	}

	fmt.Fprint(writer, ",")
	for index, value := range orderedColHeaders {
		fmt.Fprint(writer, value)
		if index != len(csv.ColHeaders)-1 {
			fmt.Fprint(writer, ",")
		}
	}
	fmt.Fprintln(writer)

	orderedRowHeaders := make([]string, len(csv.RowHeaders))

	for key, value := range csv.RowHeaders {
		orderedRowHeaders[value] = key
	}

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
