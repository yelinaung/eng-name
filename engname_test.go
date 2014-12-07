package engname

import (
	"testing"
)

func TestAreMenNamesRamdonOrNot(t *testing.T) {
	firstManName := GetMenName()
	secondManName := GetMenName()
	if firstManName == secondManName {
		t.Error("two ramdom man names cannot be the same")
	}
}

func TestAreWomenNamesRamdomOrNot(t *testing.T) {
	firstWomanName := GetWomenName()
	secondWomanName := GetWomenName()
	if firstWomanName == secondWomanName {
		t.Error("two ramdom woman names cannot be the same")
	}
}
