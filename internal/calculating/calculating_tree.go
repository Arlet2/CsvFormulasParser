package calculating

import (
	"errors"
	"strconv"
	"test_task/internal/format"
	"test_task/internal/operations"
	"test_task/internal/parsing"
)

type TreeCreatingError error
type calculatingTree map[string][]string

func CreateTree(csv format.Csv) (calculatingTree, error) {
	tree := make(map[string][]string, 0)
	for index, line := range csv.Data {
		for jndex, cell := range line {
			if parsing.IsFormula(cell) {
				link, err := csv.GetLinkByIndexes(jndex, index)
				if err != nil {
					panic(err)
				}
				formula := parsing.ParseFormula(cell)

				// если формулу можно посчитать сразу, то вычисляем на месте
				if !formula.FirstOperand.IsLink() && !formula.SecondOperand.IsLink() {
					value, err := formula.Action(formula.FirstOperand.GetConstant(),
						formula.SecondOperand.GetConstant())

					if err != nil {
						return calculatingTree{}, err.(operations.CalculatingError)
					}
					csv.Data[index][jndex] = strconv.FormatInt(int64(value), 10)
					continue
				}

				if _, ok := tree[link]; !ok {
					tree[link] = make([]string, 0)
				}

				if formula.FirstOperand.IsLink() {
					if !csv.IsLinkExist(formula.FirstOperand.GetLink()) {
						return calculatingTree{},
							errors.New("ячейки " + formula.FirstOperand.GetLink() + " не существует").(TreeCreatingError)
					}

					tree[link] = append(tree[link], formula.FirstOperand.GetLink())
				}

				if formula.SecondOperand.IsLink() {
					if !csv.IsLinkExist(formula.SecondOperand.GetLink()) {
						return calculatingTree{},
							errors.New("ячейки " + formula.SecondOperand.GetLink() + " не существует").(TreeCreatingError)
					}

					tree[link] = append(tree[link], formula.SecondOperand.GetLink())
				}
			}
		}
	}

	return tree, nil
}
