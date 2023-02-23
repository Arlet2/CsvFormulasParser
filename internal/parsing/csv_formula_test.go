package parsing

import "testing"

func TestCellParsing(t *testing.T) {
	t.Log("Start check parsing cells.")
	{
		testID := 0
		
		t.Logf("\tTest %d: check parsing only numbers formula", testID)
		{
			formula, err := parseCell("=1+2")

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.FailNow()
			}

			if formula.firstOperand.isLink() {
				t.Logf("\tFail on test %d. 1 was parsed like link", testID)
				t.FailNow()
			}

			if formula.secondOperand.isLink() {
				t.Logf("\tFail on test %d. 2 was parsed like link", testID)
				t.FailNow()
			}

			if formula.firstOperand.getConstant() != 1 {
				t.Logf("\tFail on test %d. Expected 1 but got %d", testID, formula.firstOperand.getConstant())
				t.FailNow()
			}

			if formula.secondOperand.getConstant() != 2 {
				t.Logf("\tFail on test %d. Expected 2 but got %d", testID, formula.secondOperand.getConstant())
				t.FailNow()
			}

			value, err := formula.action(1, 2)

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
			_, err := parseCell("Cell30+Cell20")

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check parsing formula only with links", testID)
		{
			formula, err := parseCell("=Cell30+Cell20")

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.FailNow()
			}

			if !formula.firstOperand.isLink() {
				t.Logf("\tFail on test %d. Cell30 was parsed like constant", testID)
				t.FailNow()
			}

			if !formula.secondOperand.isLink() {
				t.Logf("\tFail on test %d. Cell20 was parsed like constant", testID)
				t.FailNow()
			}

			if formula.firstOperand.getLink() != "Cell30" {
				t.Logf("\tFail on test %d. Expected Cell30 but got "+formula.firstOperand.getLink(), testID)
				t.FailNow()
			}

			if formula.secondOperand.getLink() != "Cell20" {
				t.Logf("\tFail on test %d. Expected Cell20 but got "+formula.secondOperand.getLink(), testID)
				t.FailNow()
			}

			value, err := formula.action(1, 1)

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
			_, err := parseCell("=2^2")

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing got", testID)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check link and constant formula", testID)
		{
			formula, err := parseCell("=2+Cell30")

			if err != nil {
				t.Logf("\tFail on test %d. Found error: "+err.Error(), testID)
				t.FailNow()
			}

			if formula.firstOperand.isLink() {
				t.Logf("\tFail on test %d. 2 was parsed like link", testID)
				t.FailNow()
			}

			if !formula.secondOperand.isLink() {
				t.Logf("\tFail on test %d. Cell30 was parsed like constant", testID)
				t.FailNow()
			}

			if formula.firstOperand.getConstant() != 2 {
				t.Logf("\tFail on test %d. Expected 2 but got %d", testID, formula.firstOperand.getConstant())
				t.FailNow()
			}

			if formula.secondOperand.getLink() != "Cell30" {
				t.Logf("\tFail on test %d. Expected Cell30 but got "+formula.secondOperand.getLink(), testID)
				t.FailNow()
			}

			value, err := formula.action(1, 2)

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