package parsing

import (
	"testing"
)

func TestLinkParsing(t *testing.T) {
	t.Log("Start check parsing link.")
	{
		testID := 0

		t.Logf("\tTest %d: check parsing simple link", testID)
		{
			col, row := ParseLink(link{"A10"})

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
