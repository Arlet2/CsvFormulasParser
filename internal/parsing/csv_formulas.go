package parsing

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"test_task/internal/operations"
)

type operand interface {
	IsLink() bool
	GetConstant() int
	GetLink() string
}

type link struct {
	value string
}

func (link) IsLink() bool {
	return true
}

func (link) GetConstant() int {
	return 0
}

func (link link) GetLink() string {
	return link.value
}

type constant struct {
	value int
}

func (constant) IsLink() bool {
	return false
}

func (constant constant) GetConstant() int {
	return constant.value
}

func (constant) GetLink() string {
	return ""
}

type formula struct {
	FirstOperand  operand
	SecondOperand operand
	Action        operations.Operation
}

type FormulaParseError error

func IsFormula(input string) bool {
	regex := regexp.MustCompile(`=[A-z0-9]+[` + GetRegexOperations() + `][A-z0-9]+`)

	return regex.FindAllString(input, 1) != nil
}

func GetRegexOperations() string {
	regexOperations := ""
	for key := range operations.AllowedOperations {
		// экранируем для регулярок
		regexOperations += "\\" + key
	}
	return regexOperations
}

func ParseCell(cell string) (formula, error) {
	if cell[0] != '=' {
		return formula{}, errors.New(cell + " не формула. Формулы начинаются с =").(FormulaParseError)
	}
	cell = strings.ReplaceAll(cell, " ", "")
	cell = strings.ReplaceAll(cell, "\t", "")

	regexOperations := GetRegexOperations()

	if !IsFormula(cell) {
		regexOperations = strings.ReplaceAll(regexOperations, "\\", "")
		return formula{},
			errors.New(cell + " неправильная формула. Формула должна быть в формате =OP1 [" + regexOperations + "] OP2")
	}

	// Дальнейшие регулярки точно найдут совпадения,
	// так как они являются лишь частями регулярки, которая проверялась до этого (функция IsFormula)

	regex := regexp.MustCompile(`[A-z0-9]+`)

	notParsedOperands := regex.FindAllString(cell, 2)

	notParsedOperand := notParsedOperands[0]
	numericValue, err := strconv.ParseInt(notParsedOperand, 10, 64)

	var firstOperand operand
	// если не удалось спарсить число == это ссылка
	if err != nil {
		firstOperand = link{notParsedOperand}
	} else { // удалось спарсить == это константа
		firstOperand = constant{int(numericValue)}
	}

	notParsedOperand = notParsedOperands[1]
	numericValue, err = strconv.ParseInt(notParsedOperand, 10, 64)

	var secondOperand operand
	// если не удалось спарсить число == это ссылка
	if err != nil {
		secondOperand = link{notParsedOperand}
	} else { // удалось спарсить == это константа
		secondOperand = constant{int(numericValue)}
	}

	regex = regexp.MustCompile(`[+` + regexOperations + `+]`)

	action := operations.AllowedOperations[regex.FindAllString(cell, 1)[0]]

	return formula{FirstOperand: firstOperand, SecondOperand: secondOperand, Action: action}, nil
}
