package parsing

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {

	t.Log("Start check parsing of CSV files.")
	{
		testID := 0

		t.Logf("\tTest %d: check parsing example file", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n2,2,=A1+Cell30,0\n30,0,=B1+A1,5")

			csv, err := ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Get error from parse: "+err.Error(), testID)
				t.Fail()
			}

			if !reflect.DeepEqual(csv.ColHeaders, map[string]int{"A": 0, "B": 1, "Cell": 2}) {
				t.Logf("\tFail on test %d. Headers were parsed wrong: %q", testID, csv.ColHeaders)
				t.Fail()
			}

			if !reflect.DeepEqual(csv.RowHeaders, map[string]int{"1": 0, "2": 1, "30": 2}) {
				t.Logf("\tFail on test %d. Row headers were parsed wrong: %q", testID, csv.RowHeaders)
				t.Fail()
			}

			if !reflect.DeepEqual(csv.Data[0], []string{"1", "0", "1"}) ||
				!reflect.DeepEqual(csv.Data[1], []string{"2", "=A1+Cell30", "0"}) ||
				!reflect.DeepEqual(csv.Data[2], []string{"0", "=B1+A1", "5"}) {
				t.Logf("\tFail on test %d. Data were parsed wrong: %q", testID, csv.Data)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check parsing not empty first cell", testID)
		{
			reader := strings.NewReader("0,A,B,Cell\n1,1,0,1\n2,2,=A1+Cell30,0\n30,0,=B1+A1,5")

			_, err := ParseCsv(reader)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.Fail()
			}
			if _, ok := err.(CsvParseError); !ok {
				t.Logf("\tFail on test %d. Expected CsvParseError but found another: "+err.Error(), testID)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check parsing invalid headers", testID)
		{
			reader := strings.NewReader(",A,2,Cell\n1,1,0,1\n2,2,=A1+Cell30,0\n30,0,=B1+A1,5")

			_, err := ParseCsv(reader)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.Fail()
			}
			if _, ok := err.(CsvParseError); !ok {
				t.Logf("\tFail on test %d. Expected CsvParseError but found another: "+err.Error(), testID)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check parsing invalid row header", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\nas,2,=A1+Cell30,0\n30,0,=B1+A1,5")

			_, err := ParseCsv(reader)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.Fail()
			}
			if _, ok := err.(CsvParseError); !ok {
				t.Logf("\tFail on test %d. Expected CsvParseError but found another: "+err.Error(), testID)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check parsing with newline in the end", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n")

			csv, err := ParseCsv(reader)

			if err != nil {
				t.Logf("Fail on test %d. Found error: "+err.Error(), testID)
				t.Fail()
			}

			if !reflect.DeepEqual(csv.ColHeaders, map[string]int{"A": 0, "B": 1, "Cell": 2}) ||
				!reflect.DeepEqual(csv.RowHeaders, map[string]int{"1": 0}) ||
				!reflect.DeepEqual(csv.Data[0], []string{"1", "0", "1"}) {
				t.Logf("Fail on test %d. Data were parsed wrong %q", testID, csv)
				t.Fail()
			}

		}
		testID++
	}
}

func TestPrinting(t *testing.T) {
	t.Log("Start check printing of CSV files.")
	{
		testID := 0

		t.Logf("\tTest %d: check printing example file", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n2,2,=A1+Cell30,0\n30,0,=B1+A1,5")

			csv, err := ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.Fail()
			}

			buf := bytes.NewBufferString("")

			csv.PrintWithWriter(buf)

			if buf.String() != ",A,B,Cell\n1,1,0,1\n2,2,=A1+Cell30,0\n30,0,=B1+A1,5\n" {
				t.Logf("\tFail on test %d. Expected another output but got "+buf.String(), testID)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check printing file only with headers", testID)
		{
			reader := strings.NewReader(",A,B,Cell")

			csv, err := ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.Fail()
			}

			buf := bytes.NewBufferString("")

			csv.PrintWithWriter(buf)

			if buf.String() != ",A,B,Cell\n" {
				t.Logf("\tFail on test %d. Expected another output but got "+buf.String(), testID)
				t.Fail()
			}

		}
		testID++
	}
}
