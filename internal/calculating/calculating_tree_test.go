package calculating

import (
	"bytes"
	"reflect"
	"strings"
	"test_task/internal/parsing"
	"testing"
)

func TestCalculatingTreeCreating(t *testing.T) {
	t.Log("Start check creating calculating graph.")
	{
		testID := 0

		t.Logf("\tTest %d: check calculating formulas without links", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n2,2,=2*2,0\n30,0,=4/2,5")

			csv, err := parsing.ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after parsing: "+err.Error(), testID)
				t.Fail()
			}

			_, err = CreateTree(csv)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after calculation: "+err.Error(), testID)
				t.Fail()
			}

			buf := bytes.NewBufferString("")

			csv.PrintWithWriter(buf)

			if buf.String() != ",A,B,Cell\n1,1,0,1\n2,2,4,0\n30,0,2,5\n" {
				t.Logf("\tFail on test %d. Expected another output but got "+buf.String(), testID)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: check creating graph from example csv", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n2,2,=A1+Cell30,0\n30,0,=B1+A1,5")

			csv, err := parsing.ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after parsing: "+err.Error(), testID)
				t.Fail()
			}

			tree, err := CreateTree(csv)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after tree creating: "+err.Error(), testID)
				t.Fail()
			}

			if reflect.DeepEqual(tree, map[string][]string{"B2": {"A1", "Cell30"}, "B30": {"B1", "A1"}}) {
				t.Logf("\tFail on test %d. Expected another tree but found: %q", testID, tree)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: check creating graph with not existing row", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n2,2,=A1+Cell40,0\n30,0,=B1+A1,5")

			csv, err := parsing.ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after parsing: "+err.Error(), testID)
				t.Fail()
			}

			_, err = CreateTree(csv)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: check creating graph with not existing col", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n2,2,=A1+Cells30,0\n30,0,=B1+A1,5")

			csv, err := parsing.ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after parsing: "+err.Error(), testID)
				t.Fail()
			}

			_, err = CreateTree(csv)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: check creating graph with zero division", testID)
		{
			reader := strings.NewReader(",A,B,Cell\n1,1,0,1\n2,2,=5/0,0\n30,0,=B1+A1,5")

			csv, err := parsing.ParseCsv(reader)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after parsing: "+err.Error(), testID)
				t.Fail()
			}

			_, err = CreateTree(csv)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.Fail()
			}
		}
		testID++

	}
}

func TestSortingTree(t *testing.T) {
	t.Log("Start check sorting tree.")
	{
		testID := 0

		t.Logf("\tTest %d: check sorting tree from example csv", testID)
		{
			tree := calculatingTree{nodes: map[string][]string{"B30": {"B1", "A1"}, "B2": {"A1", "Cell30"}}}

			sortedNodes, err := tree.SortTree()

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.Fail()
			}

			if !reflect.DeepEqual(sortedNodes, []string{"B2", "Cell30", "B30", "A1", "B1"}) &&
				!reflect.DeepEqual(sortedNodes, []string{"B30", "B1", "B2", "Cell30", "A1"}) {
				t.Logf("\tFail on test %d. Found another (maybe correct) sorting: %q", testID, sortedNodes)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: check sorting tree with cycle", testID)
		{
			tree := calculatingTree{nodes: map[string][]string{
				"A1": {"B2", "C6"}, "B2": {"C6", "B3"}, "B3": {"C7", "B4"}, "B4": {"A1"}}}

			_, err := tree.SortTree()

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.Fail()
			}
		}
		testID++
	}
}
