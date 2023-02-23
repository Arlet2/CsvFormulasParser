package operations
import (
	"testing"
)

func TestOpeartions(t *testing.T) {
	t.Log("Start check operations.")
	{
		testID := 0

		t.Logf("\tTest %d: check division by zero", testID)
		{
			_, err := AllowedOperations["/"](2, 0)

			if err == nil {
				t.Logf("\tFail on test %d. Expected error but nothing was got", testID)
				t.Fail()
			}

			if _, ok := err.(CalculatingError); !ok {
				t.Logf("\tFail on test %d. Expected another error but found "+err.Error(), testID)
				t.Fail()
			}
		}
		testID++

		// другие тесты для будущих операций...
	}
}