package parsing

import "testing"

func TestCellParsing(t *testing.T) {
	t.Log("Start check parsing cells.")
	{
		testID := 0
		
		t.Logf("\tTest %d: check parsing only numbers formula", testID)
		{
			formula, err := ParseCell("=1+2")

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.FailNow()
			}

			if formula.FirstOperand.isLink() {
				t.Logf("\tFail on test %d. 1 was parsed like link", testID)
				t.FailNow()
			}

			if formula.SecondOperand.isLink() {
				t.Logf("\tFail on test %d. 2 was parsed like link", testID)
				t.FailNow()
			}

			if formula.FirstOperand.getConstant() != 1 {
				t.Logf("\tFail on test %d. Expected 1 but got %d", testID, formula.FirstOperand.getConstant())
				t.FailNow()
			}

			if formula.SecondOperand.getConstant() != 2 {
				t.Logf("\tFail on test %d. Expected 2 but got %d", testID, formula.SecondOperand.getConstant())
				t.FailNow()
			}

			value, err := formula.Action(1, 2)

			if err != nil {
				t.Logf("\tFail on test %d. Found error when use operation: "+err.Error(), testID)
				t.FailNow()
			}

			if value != 3 {
				t.Logf("\tFail on test %d. Expected 3 but got %d", testID, value)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check missing = symbol on start", testID)
		{
			_, err := ParseCell("Cell30+Cell20")

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check parsing formula only with links", testID)
		{
			formula, err := ParseCell("=Cell30+Cell20")

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.FailNow()
			}

			if !formula.FirstOperand.isLink() {
				t.Logf("\tFail on test %d. Cell30 was parsed like constant", testID)
				t.FailNow()
			}

			if !formula.SecondOperand.isLink() {
				t.Logf("\tFail on test %d. Cell20 was parsed like constant", testID)
				t.FailNow()
			}

			if formula.FirstOperand.getLink() != "Cell30" {
				t.Logf("\tFail on test %d. Expected Cell30 but got "+formula.FirstOperand.getLink(), testID)
				t.FailNow()
			}

			if formula.SecondOperand.getLink() != "Cell20" {
				t.Logf("\tFail on test %d. Expected Cell20 but got "+formula.SecondOperand.getLink(), testID)
				t.FailNow()
			}

			value, err := formula.Action(1, 1)

			if err != nil {
				t.Logf("\tFail on test %d. Found error after action: "+err.Error(), testID)
				t.FailNow()
			}

			if value != 2 {
				t.Logf("\tFail on test %d. Action + was parsed wrongly. Expected 2 but got %d", testID, value)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check not allowed operation", testID)
		{
			_, err := ParseCell("=2^2")

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check link and constant formula", testID)
		{
			formula, err := ParseCell("=2+Cell30")

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.FailNow()
			}

			if formula.FirstOperand.isLink() {
				t.Logf("\tFail on test %d. 2 was parsed like link", testID)
				t.FailNow()
			}

			if !formula.SecondOperand.isLink() {
				t.Logf("\tFail on test %d. Cell30 was parsed like constant", testID)
				t.FailNow()
			}

			if formula.FirstOperand.getConstant() != 2 {
				t.Logf("\tFail on test %d. Expected 2 but got %d", testID, formula.FirstOperand.getConstant())
				t.FailNow()
			}

			if formula.SecondOperand.getLink() != "Cell30" {
				t.Logf("\tFail on test %d. Expected Cell30 but got "+formula.SecondOperand.getLink(), testID)
				t.FailNow()
			}

			value, err := formula.Action(1, 2)

			if err != nil {
				t.Logf("\tFail on test %d. Found error from action: "+err.Error(), testID)
				t.FailNow()
			}

			if value != 3 {
				t.Logf("\tFail on test %d. Action works wrongly. Expected 3 but got %d", testID, value)
				t.FailNow()
			}

		}
		testID++
	}
}