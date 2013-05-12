package main

import (
	"usc/ordenacion"
	"fmt"
	
)



func main(){
	//var a []int
	a := ordenacion.CrearArrayAleatorio(20000)
	
	//a:=[]int{2,54,7,98,56,51,0,23}
	ordenacion.OrdenaShellsortParalelo(a)
	//fmt.Println("array ordenado =",a)
	fmt.Println(ordenacion.EstaOrdenado(a))
	
	ordenacion.AbrirBasePruebas()
}
