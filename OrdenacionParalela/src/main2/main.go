package main

import (
	"usc/ordenacion"
	"fmt"
)

func main() {
	a := ordenacion.CrearArrayAleatorio(10007)
	//a := []int{2,34,56,12,98,23,66,91,20,45,676,221,434,17,33,234,100, 5}
	//a := []int{2,34,5135126,12,98,23,124166,91,20,4512335,676,221,434}
	//fmt.Println(a)
	//ordenacion.ParallellQuicksort(a)

	ordenacion.OrdenaRadixSortParalelo(a, 3)
	//ordenacion.OrdenaPSRS(a)
	//fmt.Println(a)
	fmt.Println(ordenacion.EstaOrdenado(a))

}
