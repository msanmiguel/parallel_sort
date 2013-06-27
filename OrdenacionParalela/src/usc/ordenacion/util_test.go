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

func TestCrearArrayAleatorio(t *testing.T) {
	array := CrearArrayAleatorio(100)
	if len(array) != 100 {
		t.Errorf("La longitud del array creado no es correcta");
	}
}

