package naturalsort

import (
	"regexp"
	"strconv"
	"strings"
)

type NaturalSort []string

var r = regexp.MustCompilePOSIX(`[^0-9]+|[0-9]+`)

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
			// Try to convert the left and right portions to integer
			inti, erri := strconv.Atoi(spliti[index])
			intj, errj := strconv.Atoi(splitj[index])
			// Handle numerical case
			if erri == nil && errj == nil { //if number
				if inti == intj {
					return len(spliti[index]) < len(splitj[index])
				}
				return inti < intj
			}
			// Handle non-numerical case
			return spliti[index] < splitj[index]
		}

	}
	return s[i] < s[j]
}
