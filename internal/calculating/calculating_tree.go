package calculating

import (
	"errors"
	"strconv"
	"test_task/internal/format"
	"test_task/internal/operations"
	"test_task/internal/parsing"
)

type TreeCreatingError error
type FormulaCycleError error
type calculatingTree struct {
	nodes map[string][]string
}

func CreateTree(csv format.Csv) (calculatingTree, error) {
	nodes := make(map[string][]string, 0)
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

				// если вершины не существует, то создаем запись о ней
				if _, ok := nodes[link]; !ok {
					nodes[link] = make([]string, 0)
				}

				if formula.FirstOperand.IsLink() {
					if !csv.IsLinkExist(formula.FirstOperand.GetLink()) {
						return calculatingTree{},
							errors.New("ячейки " + formula.FirstOperand.GetLink() + " не существует").(TreeCreatingError)
					}

					nodes[link] = append(nodes[link], formula.FirstOperand.GetLink())
				}

				if formula.SecondOperand.IsLink() {
					if !csv.IsLinkExist(formula.SecondOperand.GetLink()) {
						return calculatingTree{},
							errors.New("ячейки " + formula.SecondOperand.GetLink() + " не существует").(TreeCreatingError)
					}

					nodes[link] = append(nodes[link], formula.SecondOperand.GetLink())
				}
			}
		}
	}

	return calculatingTree{nodes: nodes}, nil
}

// (!) сортировка не детерминирована из-за недетерминированности порядка ключей в tree.nodes
func (tree calculatingTree) SortTree() ([]string, error) {

	nodesState := make(map[string]int, 0)
	sortedNodes := make([]string, 0)

	for key := range tree.nodes {
		nodesState[key] = 0
	}

	for key := range tree.nodes {
		err := tree.dfc(key, &nodesState, &sortedNodes)
		if err != nil {
			return nil, err
		}
	}

	// сортировка представлена в обратном порядке, так как листья уже известны и их считать не надо

	return sortedNodes, nil
}

func (tree calculatingTree) dfc(currentNode string, nodesState *map[string]int, sortedNodes *[]string) error {

	if (*nodesState)[currentNode] == 1 {
		return errors.New("обнаружена циклическая зависимость у формул").(FormulaCycleError)
	}
	if (*nodesState)[currentNode] == 2 {
		return nil
	}

	(*nodesState)[currentNode] = 1

	for _, value := range tree.nodes[currentNode] {
		err := tree.dfc(value, nodesState, sortedNodes)
		if err != nil {
			return err
		}
	}

	(*nodesState)[currentNode] = 2
	*sortedNodes = append(*sortedNodes, currentNode)

	return nil
}

func CalculateNodes(csv *format.Csv, sortedNodes []string) error {
	var col, row, cell string
	for _, value := range sortedNodes {
		col, row = format.ParseLink(value)
		cell = csv.Data[csv.RowHeaders[row]][csv.ColHeaders[col]]
		if parsing.IsFormula(cell) {
			formula := parsing.ParseFormula(cell)
			err := processFormula(value, csv, formula)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func processFormula(currentCellLink string, csv *format.Csv, formula parsing.Formula) error {
	var col, row, cellValue string
	var parsedFirstOperand, parsedSecondOperand int64
	if formula.FirstOperand.IsLink() {
		col, row = format.ParseLink(formula.FirstOperand.GetLink())
		cellValue = csv.Data[csv.RowHeaders[row]][csv.ColHeaders[col]]

		//if parsing.IsFormula(cellValue) {
		//	processFormula(col+row, csv, parsing.ParseFormula(cellValue))
		//}

		// игнорируем ошибку, так как до этого были проверки на число в ячейке
		parsedFirstOperand, _ = strconv.ParseInt(cellValue, 10, 32)
	} else {
		parsedFirstOperand = int64(formula.FirstOperand.GetConstant())
	}

	if formula.SecondOperand.IsLink() {
		col, row = format.ParseLink(formula.SecondOperand.GetLink())
		cellValue = csv.Data[csv.RowHeaders[row]][csv.ColHeaders[col]]

		//if parsing.IsFormula(cellValue) {
		//	processFormula(col+row, csv, parsing.ParseFormula(cellValue))
		//}

		// игнорируем ошибку, так как до этого были проверки на число в ячейке
		parsedSecondOperand, _ = strconv.ParseInt(cellValue, 10, 32)
	} else {
		parsedSecondOperand = int64(formula.SecondOperand.GetConstant())
	}

	calculatedValue, err := formula.Action(int(parsedFirstOperand), int(parsedSecondOperand))

	if err != nil {
		return err
	}

	col, row = format.ParseLink(currentCellLink)
	csv.Data[csv.RowHeaders[row]][csv.ColHeaders[col]] = strconv.FormatInt(int64(calculatedValue), 10)

	return nil
}
