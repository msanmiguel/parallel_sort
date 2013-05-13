package ordenacion

import (

)

func ordenaQuicksortParalelo(a []int) {
	c := make(chan int)
	go ordenaQuicksortParalelo_rec(a, c)
	<-c
} 

func ordenaQuicksortParalelo_rec(a []int, c chan int){
	if len(a) > 1 {
		c2 := make(chan int)
		pos_pivote := recolocar(a) 
		go ordenaQuicksortParalelo_rec(a[:pos_pivote], c2) // recoloco la lista de los menores
		go ordenaQuicksortParalelo_rec(a[(pos_pivote+1):], c2) // recoloco la lista de los mayores
		<- c2
		<- c2
	}
	c <- 0
} 

func funcion(a []int, salto int, c chan int, k int){
	for i:=k+salto; i<len(a);i+=salto {
		p := a[i]
		j:=i-salto
		for ; j>=0 && a[j]>p; j-=salto{
				a[j+salto]=a[j]
			}
				a[j+salto] = p	
	}
	c <-0
}

func ordenaShellsortParalelo(a []int){
	salto:= len(a)/2
	c := make(chan int)
	for salto >= 1 {
		for k := 0; k < salto; k++ {
			go funcion(a,salto,c,k)
				
		}
		for k := 0; k < salto; k++ {
			<-c
		}
		salto=salto/2
	}
}

type OrdenacionParal interface{
	Ordenar(a []int)
	ObtenerNombreAlgoritmo()string
}

type QuicksortParal1 struct{}
type ShellsortParal1 struct{}

func (o QuicksortParal1) Ordenar(a []int){
	ordenaQuicksortParalelo(a)
}
func (o QuicksortParal1) ObtenerNombreAlgoritmo()string{
	return "QuicksortParal1"
}

func (o ShellsortParal1) Ordenar(a []int){
	ordenaShellsortParalelo(a)
}
func (o ShellsortParal1) ObtenerNombreAlgoritmo()string{
	return "ShellsortParal1"
}


