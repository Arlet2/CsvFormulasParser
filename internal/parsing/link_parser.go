package parsing

import "regexp"

func ParseLink(link link) (string, string) {
	// Ссылка корректна, так как была проверена при парсинге формул

	var col, row string
	regex := regexp.MustCompile(`[A-z]+`)
	
	col = regex.FindString(link.value)

	regex = regexp.MustCompile(`[0-9]+`)
	row = regex.FindString(link.value)
	return col, row
}