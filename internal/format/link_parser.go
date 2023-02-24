package format

import "regexp"

func ParseLink(link string) (string, string) {
	// Ссылка корректна, так как была проверена при парсинге формул

	var col, row string
	regex := regexp.MustCompile(`[A-z]+`)

	col = regex.FindString(link)

	regex = regexp.MustCompile(`[0-9]+`)
	row = regex.FindString(link)
	return col, row
}

func (csv Csv) IsLinkExist(link string) bool {
	col, row := ParseLink(link)

	if _, ok := csv.ColHeaders[col]; !ok {
		return false
	}

	if _, ok := csv.RowHeaders[row]; !ok {
		return false
	}

	return true
}
