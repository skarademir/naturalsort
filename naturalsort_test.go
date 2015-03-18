package naturalsort

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortValid(t *testing.T) {
	data := []string{"A1BA1", "A11AA1", "A2AB0", "B1AA1", "A1AA1"}
	expected := []string{"A1AA1", "A1BA1", "A2AB0", "A11AA1", "B1AA1"}

	sort.Sort(NaturalSort(data))
	if !reflect.DeepEqual(data, expected) {
		t.Fatalf("Wrong order.\nExpected=%v\nGot=%v", expected, data)
	}
}

func BenchmarkSort(b *testing.B) {
	var data = [...]string{"A1BA1", "A11AA1", "A2AB0", "B1AA1", "A1AA1"}
	for ii := 0; ii < b.N; ii++ {
		d := NaturalSort(data[:])
		sort.Sort(d)
	}
}
