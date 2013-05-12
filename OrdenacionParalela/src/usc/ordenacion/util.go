package ordenacion

import (
	"math/rand"
)

func EstaOrdenado(a []int) bool{
	i:=0
	for i=0; i<len(a)-1; i++{
		if a[i]>a[i+1]{
		
			return false;
		}
	}
	return true

}


func CrearArrayAleatorio(n int) []int{
	a := make([]int, n)
	for i:= 0; i < n; i++ {
		a[i]=rand.Int()
	}
	return a
}