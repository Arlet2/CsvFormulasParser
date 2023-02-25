package main

import (
	"bytes"
	_ "embed"
	"strings"
	"test_task/internal/calculating"
	"test_task/internal/parsing"
	"testing"
)

//go:embed test_csvs/example.csv
var exampleCsv string

func TestMain(t *testing.T) {

	t.Log("Start check calculating nodes.")
	{
		testID := 0

		t.Logf("\tTest %d: check calculating example csv", testID)
		{
			csv, err := parsing.ParseCsv(strings.NewReader(exampleCsv))

			if err != nil {
				t.Logf("\tFail on test %d. Get error from parse: "+err.Error(), testID)
				t.FailNow()
			}

			tree, err := calculating.CreateTree(csv)

			if err != nil {
				t.Logf("\tFail on test %d. Get error from tree generating: "+err.Error(), testID)
				t.FailNow()
			}

			sortedNodes, err := tree.SortTree()

			if err != nil {
				t.Logf("\tFail on test %d. Get error from tree generating: "+err.Error(), testID)
				t.FailNow()
			}

			err = calculating.CalculateNodes(&csv, sortedNodes)

			if err != nil {
				t.Logf("\tFail on test %d. Get error from node calculating: "+err.Error(), testID)
				t.FailNow()
			}

			buf := bytes.NewBufferString("")

			csv.PrintWithWriter(buf)

			if buf.String() != ",A,B,Cell\n1,1,0,1\n2,2,6,0\n30,0,1,5\n" {
				t.Logf("\tFail on test %d. Expected another output but got "+buf.String(), testID)
				t.Fail()
			}

		}
		testID++
	}
}
