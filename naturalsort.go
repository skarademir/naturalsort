package naturalsort

import (
	//"fmt"
	"regexp"
	"strconv"
)
type NaturalSort []string
func (s NaturalSort) Len() int {
    return len(s)
}
func (s NaturalSort) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s NaturalSort) Less(i, j int) bool {
	//fmt.Println(s[i])
	spliti := regexp.MustCompilePOSIX(`^([^0-9]*)+|[0-9]+`).FindAllString(regexp.MustCompilePOSIX(` `).ReplaceAllString(s[i],""),-1)
	splitj := regexp.MustCompilePOSIX(`^([^0-9]*)+|[0-9]+`).FindAllString(regexp.MustCompilePOSIX(` `).ReplaceAllString(s[j],""),-1)
	//fmt.Println(spliti)
	for index := range spliti {
		if spliti[index] != splitj[index] {
			inti, ei := strconv.Atoi(spliti[index])
			intj, ej := strconv.Atoi(splitj[index])
			if ei==nil&&ej==nil { //if number
				return inti < intj
			}
			return spliti[index] < splitj[index]
		}
		
	}
    return s[i] < s[j]
}


//func main() {
//    fruits := []string{"z24", "z2", "z15", "z1",
//                       "z3", "z20", "z5", "z11",
//                       "z 21", "z22"}
//    sort.Sort(NaturalSort(fruits))
//    fmt.Println(fruits)
//}
