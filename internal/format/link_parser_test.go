package format

import (
	"testing"
)

func TestLinkParsing(t *testing.T) {
	t.Log("Start check parsing link.")
	{
		testID := 0

		t.Logf("\tTest %d: check parsing simple link", testID)
		{
			col, row := ParseLink("A10")

			if col != "A" {
				t.Logf("\tFail on test %d. Expected A but found "+col, testID)
				t.Fail()
			}

			if row != "10" {
				t.Logf("\tFail on test %d. Expected 10 but found "+row, testID)
				t.Fail()
			}
		}
		testID++
	}
}

func TestLinkExisting(t *testing.T) {
	t.Log("Start check existing links in CSV.")
	{
		testID := 0

		t.Logf("\tTest %d: check existing link", testID)
		{
			csv := Csv{ColHeaders: map[string]int{"A": 0, "B": 1}, 
					RowHeaders: map[string]int{"30": 0, "20": 1}, Data: [][]string{}}

			if !csv.IsLinkExist("A30") {
				t.Logf("\tFail on test %d. Expected true but found false", testID)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check missing link", testID)
		{
			csv := Csv{ColHeaders: map[string]int{"A": 0, "B": 1}, 
					RowHeaders: map[string]int{"30": 0, "20": 1}, Data: [][]string{}}

			if csv.IsLinkExist("Cell30") {
				t.Logf("\tFail on test %d. Expected false but found true", testID)
				t.Fail()
			}

		}
		testID++
	}
}
