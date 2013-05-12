package ordenacion

import (
	"testing"
)

func TestEstaOrdenado(t *testing.T){

	arrayVacio :=[]int{}
	if !EstaOrdenado(arrayVacio){
		t.Errorf("El array vacío no se considera ordenado")
	}

}

