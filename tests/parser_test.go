package tests

import (
	"testing"
)

func Parser(t *testing.T) {

	p := NewParser()
	p.Read("ArChivo")

	got := p.GetNamePlural()
	want := "Archivos"

	if got != want {
		t.Errorf("Tengo %v | Quiero %v", got, want)
	} else {
		t.Logf("Tengo %v | Quiero %v", got, want)
	}

}
