package operations

import "errors"

type CalculatingError error

var AllowedOperations = map[string]Operation{
	"+": add, "-": sub, "*": mul, "/": div,
}

type Operation func(int, int) (int, error)

func add (a int, b int) (int, error) {
	return a+b, nil
}

func sub (a int, b int) (int, error) {
	return a-b, nil
}

func mul (a int, b int) (int, error) {
	return a*b, nil
}

func div (a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a/b, nil
}