package main

import (
	"usc/ordenacion"
	"fmt"
)

func main() {
	a := ordenacion.CrearArrayAleatorio(32000)
	//fmt.Println(a)
	//ordenacion.ParallellQuicksort(a)
	//ordenacion.OrdenaRadixSort(a)
	ordenacion.OrdenaRadixSortParalelo(a)
	fmt.Println(ordenacion.EstaOrdenado(a))

}
