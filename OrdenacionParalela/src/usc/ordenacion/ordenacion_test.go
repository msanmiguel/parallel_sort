package ordenacion

import (
	"testing"
	"fmt"
)


func TestMergesortSecuencial(t *testing.T){
	var a []int

	a = CrearArrayAleatorio(1000)
	
		//fmt.Println(a)
	ordenaMergesortSecuencial(a, true)
	ordenado := EstaOrdenado(a)
	fmt.Println(a)
	if !ordenado{
		t.Error("El array no esta ordenado") 
	}
}

func TestMergesortParalelo(t *testing.T){
	var a []int
	a = CrearArrayAleatorio(1000)
	//a = []int {3, 10, 1, 7, 5, 6, 8, 4, 9 ,2}
	ordenaMergesortParalelo(a)
	ordenado := EstaOrdenado(a)
	fmt.Println(a)
	if !ordenado{
		t.Error("El array no esta ordenado") 
	}
}

func TestShellsortParalelo(t *testing.T){
	var a []int
	a = CrearArrayAleatorio(1000)
	//a = []int {3, 10, 1, 7, 5, 6, 8, 4, 9 ,2}
	ordenaShellsortParalelo1(a)
	ordenado := EstaOrdenado(a)
	fmt.Println(a)
	if !ordenado{
		t.Error("El array no esta ordenado") 
	}

}