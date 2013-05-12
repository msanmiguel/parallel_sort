package ordenacion

import (

)

func OrdenaQuicksortParalelo(a []int) {
	c := make(chan int)
	go ordenaQuicksortParalelo(a, c)
	<-c
} 

func ordenaQuicksortParalelo(a []int, c chan int){
	if len(a) > 1 {
		c2 := make(chan int)
		pos_pivote := recolocar(a) 
		go ordenaQuicksortParalelo(a[:pos_pivote], c2) // recoloco la lista de los menores
		go ordenaQuicksortParalelo(a[(pos_pivote+1):], c2) // recoloco la lista de los mayores
		<- c2
		<- c2
	}
	c <- 0
} 

func funcion(a []int, salto int, c chan int, k int){
	for i:=k+salto; i<len(a);i+=salto {
		p := a[i]
		for j:=i-salto; j>=0 && a[j]>p; j-=salto{
			if a[j] > p { 
				a[j+salto]=a[j]
			} else {
				a[j] = p
				break
			}
		}
	}
	c <-0
}

func OrdenaShellsortParalelo(a []int) []int{
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
	return a
}


