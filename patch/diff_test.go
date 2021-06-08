package patch

import (
	"fmt"
	doc "github.com/frolFomich/abstract-document"
	"testing"
)

func Test_Diff(t *testing.T) {
	d1 := doc.New()
	dn1 := doc.New()

	dn1.Put("E", 200)
	dn1.Put("F", []int{1, 2, 3})

	d1.Put("A", 100)
	d1.Put("B", "abba")
	d1.Put("C", true)
	d1.Put("D", dn1)

	d2 := doc.New()
	dn2 := doc.New()

	dn2.Put("E", 202)
	dn2.Put("F", []int{1, 3, 5})

	dn3 := doc.New()

	dn3.Put("O", "sss")
	dn3.Put("P", 99)

	d2.Put("A", 101)
	d2.Put("D", dn2)
	d2.Put("Z", "zzz")
	d2.Put("C", nil)
	d2.Put("B", dn3)

	diffs := Diff(d2, d1)
	for _, d := range diffs {
		json, err := d.MarshalJson()
		if err != nil {
			t.Errorf("Error: Diff() = %v", err)
			return
		}
		fmt.Println(string(json))
	}
}
