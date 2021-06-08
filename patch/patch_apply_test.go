package patch

import (
	"fmt"
	doc "github.com/frolFomich/abstract-document"
	"testing"
)

func Test_Patch(t *testing.T) {
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
	dn2.Put("G", true)

	dn3 := doc.New()

	dn3.Put("O", "sss")
	dn3.Put("P", 99)

	d2.Put("A", 101)
	d2.Put("D", dn2)
	d2.Put("Z", "zzz")
	d2.Put("C", nil)
	d2.Put("B", dn3)

	bytes, err := d1.MarshalJson()
	if err != nil {
		t.Errorf("Error: Patch() = %v", err)
		return
	}
	fmt.Printf("Initial doc:\n %s\n", string(bytes))

	diffs := Diff(d2, d1)
	fmt.Println("Patches:")
	for _, p := range diffs {
		bytes, err := p.MarshalJson()
		if err != nil {
			t.Errorf("Error: Patch() = %v", err)
			return
		}
		fmt.Println(string(bytes))
	}

	d3 := doc.FromOther(d1)

	patched, err := Patch(d3, diffs...)
	if err != nil {
		t.Errorf("Error: Patch() = %v", err)
	}
	bytes, err = patched.MarshalJson()
	if err != nil {
		t.Errorf("Error: MarshalJson() = %v", err)
		return
	}
	fmt.Println("Result doc:")
	fmt.Println(string(bytes))
}
