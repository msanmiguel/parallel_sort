
package reflection

import (
	"testing"
	"reflect"
)

type Prueba struct {
	id int
	nombre string
}

func TestSwap(t *testing.T) {
	p1 := []int { 4, 9, 10 }
	v1 := reflect.ValueOf(p1).Index(0)
	v2 := reflect.ValueOf(p1).Index(2)
	GenericSwap(v1, v2)
	if p1[0] != 10 || p1[2] != 4 {
		t.Error("El intercambio no fue correcto")
	}

	p2 := []string { "hola", "adios", "maria" }
	v3 := reflect.ValueOf(p2).Index(0)
	v4 := reflect.ValueOf(p2).Index(2)
	GenericSwap(v3, v4)
	//t.Log(p2)
	if p2[0] != "maria" || p2[2] != "hola" {
		t.Error("El intercambio no fue correcto")
	}

	p3 := []Prueba { Prueba{0, "Antonio"}, Prueba{2, "Marta"}, Prueba{30, "Faustino"} }
	v5 := reflect.ValueOf(p3).Index(0)
	v6 := reflect.ValueOf(p3).Index(2)
	GenericSwap(v5, v6)
	//t.Log(p3)
	if p3[0].id != 30 || p3[2].id != 0 {
		t.Error("El intercambio no fue correcto")
	}
}
