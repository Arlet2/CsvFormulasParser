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

//go:embed test_csvs/test1.csv
var test1Csv string // с не существующей ссылкой

//go:embed test_csvs/test2.csv
var test2Csv string // с циклическими ссылками

//go:embed test_csvs/test3.csv
var test3Csv string // с разными видами формул

//go:embed test_csvs/test4.csv
var test4Csv string // с разными видами формул

func TestMain(t *testing.T) {

	t.Log("Start check calculating nodes.")
	{
		testID := 0

		t.Logf("\tTest %d: check calculating example csv", testID)
		{
			csv, err := parsing.ParseCsv(strings.NewReader(exampleCsv))

			if err != nil {
				t.Logf("\tFail on test %d. Get error from parsing: "+err.Error(), testID)
				t.FailNow()
			}

			tree, err := calculating.CreateTree(csv)

			if err != nil {
				t.Logf("\tFail on test %d. Get error from tree generating: "+err.Error(), testID)
				t.FailNow()
			}

			sortedNodes, err := tree.SortTree()

			if err != nil {
				t.Logf("\tFail on test %d. Get error from tree sorting: "+err.Error(), testID)
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

		t.Logf("\tTest %d: not existing link in csv", testID)
		{
			csv, err := parsing.ParseCsv(strings.NewReader(test1Csv))

			if err != nil {
				t.Logf("\tFail on test %d. Get error from parsing: "+err.Error(), testID)
				t.FailNow()
			}

			_, err = calculating.CreateTree(csv)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.FailNow()
			}

		}
		testID++

		t.Logf("\tTest %d: csv with cycled links", testID)
		{
			csv, err := parsing.ParseCsv(strings.NewReader(test2Csv))

			if err != nil {
				t.Logf("\tFail on test %d. Get error from parsing: "+err.Error(), testID)
				t.FailNow()
			}

			tree, err := calculating.CreateTree(csv)

			if err != nil {
				t.Logf("\tFail on test %d. Get error from tree creating: "+err.Error(), testID)
				t.FailNow()
			}

			_, err = tree.SortTree()

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.FailNow()
			}

		}
		testID++

		t.Logf("\tTest %d: csv with some types of formulas", testID)
		{
			csv, err := parsing.ParseCsv(strings.NewReader(test3Csv))

			if err != nil {
				t.Logf("\tFail on test %d. Get error from parsing: "+err.Error(), testID)
				t.FailNow()
			}

			tree, err := calculating.CreateTree(csv)

			if err != nil {
				t.Logf("\tFail on test %d. Get error from tree generating: "+err.Error(), testID)
				t.FailNow()
			}

			sortedNodes, err := tree.SortTree()

			if err != nil {
				t.Logf("\tFail on test %d. Get error from tree sorting: "+err.Error(), testID)
				t.FailNow()
			}

			err = calculating.CalculateNodes(&csv, sortedNodes)

			if err != nil {
				t.Logf("\tFail on test %d. Get error from node calculating: "+err.Error(), testID)
				t.FailNow()
			}

			buf := bytes.NewBufferString("")

			csv.PrintWithWriter(buf)

			if buf.String() != ",A,B,Lol\n30,2,0,1\n2,15,3,0\n5,3,0,9\n" {
				t.Logf("\tFail on test %d. Expected another output but got "+buf.String(), testID)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: csv with extra cell", testID)
		{
			_, err := parsing.ParseCsv(strings.NewReader(test4Csv))

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.FailNow()
			}

		}
		testID++
	}
}
