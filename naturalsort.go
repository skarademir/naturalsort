package naturalsort

import (
	"regexp"
	"strings"
)

type NaturalSort []string

var r = regexp.MustCompile(`[^0-9]+|[0-9]+`)

func (s NaturalSort) Len() int {
	return len(s)
}
func (s NaturalSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s NaturalSort) Less(i, j int) bool {

	spliti := r.FindAllString(strings.Replace(s[i], " ", "", -1), -1)
	splitj := r.FindAllString(strings.Replace(s[j], " ", "", -1), -1)

	for index := 0; index < len(spliti) && index < len(splitj); index++ {
		if spliti[index] != splitj[index] {
			// Handle numerical case
			if isNumber(spliti[index]) && isNumber(splitj[index]) { //if number
				// Remove Leading Zeroes
				stringi := strings.TrimLeft(spliti[index], "0")
				stringj := strings.TrimLeft(splitj[index], "0")
				if len(stringi) == len(stringj) {
					for indexchar := 0; indexchar < len(stringi); indexchar++ {
						if stringi[indexchar] != stringj[indexchar] {
							return stringi[indexchar] < stringj[indexchar]
						}
					}
					return len(spliti[index]) < len(splitj[index])
				}
				return len(stringi) < len(stringj)
			}
			// Handle non-numerical case
			return spliti[index] < splitj[index]
		}

	}
	return s[i] < s[j]
}
func isNumber(input string) bool {
	return input[0] >= '0' && input[0] <= '9'
}
