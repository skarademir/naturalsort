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

	splitshortest := len(spliti)
	if len(spliti) > len(splitj) {
		splitshortest = len(splitj)
	}
	for index := 0; index < splitshortest; index++ {
		if spliti[index] != splitj[index] {
			inti, ei := strconv.Atoi(spliti[index])
			intj, ej := strconv.Atoi(splitj[index])
			if ei == nil && ej == nil { //if number
				if inti == intj {
					return len(spliti[index]) < len(splitj[index])
				}
				return inti < intj
			}
			return spliti[index] < splitj[index]
		}

	}
	return s[i] < s[j]
}
