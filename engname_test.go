package engname

import (
	"testing"
)

func TestAreMenNamesRamdonOrNot(t *testing.T) {
	menNames := make([]string, 100)

	for i, _ := range menNames {
		menNames[i] = GetMenName()
	}

	for i, name1 := range menNames {
		for j, name2 := range menNames {
			if i != j && name1 == name2 {
				t.Fatalf("not unique: %v : %v and %v :%v", i, j, name1, name2)
			}
		}
	}
}

func TestAreWomenNamesRamdomOrNot(t *testing.T) {
	Names := make([]string, 100)

	for i, _ := range Names {
		Names[i] = GetWomenName()
	}

	for i, name1 := range Names {
		for j, name2 := range Names {
			if i != j && name1 == name2 {
				t.Fatalf("not unique: %v : %v and %v :%v", i, j, name1, name2)
			}
		}
	}
}
