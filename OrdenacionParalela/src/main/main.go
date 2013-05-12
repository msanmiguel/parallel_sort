package main

import (
	"usc/ordenacion"
	"fmt"
	"time"
	
)



func main(){

	//a:=[]int{2,54,7,98,56,51,0,23}
	//a := ordenacion.CrearArrayDescendente(100000)
	a := ordenacion.CrearArrayAleatorio(1000000)
	//
	//a := ordenacion.CrearArrayAscendente(100000)
	
	//algoritmos de ordenación secuencial
	//fmt.Println(a)
	
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println("Ordenacion Burbuja: Tamaño de la entrada", len(a))
	t1 := time.Now()
	//ordenacion.OrdenaBurbuja1(b)
	t2 := time.Now()
	duracion := t2.Sub(t1)
	tiempo := duracion.Nanoseconds()
	fmt.Println(ordenacion.EstaOrdenado(b))
	fmt.Println("tiempo", tiempo)
	copy(b, a)
	fmt.Println("Ordenacion Inserción: Tamaño de la entrada", len(a))
	t1 = time.Now()
	//ordenacion.OrdenaInsercion1(b)
	t2 = time.Now()
	duracion =t2.Sub(t1)
	tiempo = duracion.Nanoseconds()
	fmt.Println(ordenacion.EstaOrdenado(b))
	fmt.Println("tiempo", tiempo)
	copy(b, a)
	fmt.Println("Ordenacion Mergesort: Tamaño de la entrada", len(a))
	t1 = time.Now()
	ordenacion.OrdenaMergesort1(b)
	t2 = time.Now()
	duracion = t2.Sub(t1)
	tiempo = duracion.Nanoseconds()
	fmt.Println(ordenacion.EstaOrdenado(b))
	fmt.Println("tiempo", tiempo)
	copy(b, a)
	fmt.Println("Ordenacion Quicksort: Tamaño de la entrada", len(a))
	t1 = time.Now()
	ordenacion.OrdenaQuicksort1(b)
	t2 = time.Now()
	duracion = t2.Sub(t1)
	tiempo = duracion.Nanoseconds()
	fmt.Println(ordenacion.EstaOrdenado(b))
	fmt.Println("tiempo", tiempo)
	copy(b, a)
	fmt.Println("Ordenacion Shellsort: Tamaño de la entrada", len(a))
	t1 = time.Now()
	ordenacion.OrdenaShellsort1(b)
	t2 = time.Now()
	duracion = t2.Sub(t1)
	tiempo = duracion.Nanoseconds()
	fmt.Println(ordenacion.EstaOrdenado(b))
	fmt.Println("tiempo", tiempo)
	
	//algoritmos de ordenacion paralela
	
	copy(b, a)
	fmt.Println("Ordenacion Quicksort paralelo: Tamaño de la entrada", len(a))
	t1 = time.Now()
	ordenacion.OrdenaQuicksortParalelo(b)
	t2 = time.Now()
	duracion = t2.Sub(t1)
	tiempo = duracion.Nanoseconds()
	fmt.Println(ordenacion.EstaOrdenado(b))
	fmt.Println("tiempo", tiempo)
	copy(b, a)
	fmt.Println("Ordenacion Shellsort paralelo: Tamaño de la entrada", len(a))
	t1 = time.Now()
	ordenacion.OrdenaShellsortParalelo(b)
	t2 = time.Now()
	duracion = t2.Sub(t1)
	tiempo = duracion.Nanoseconds()
	fmt.Println(ordenacion.EstaOrdenado(b))
	fmt.Println("tiempo", tiempo)
	
	//ordenacion.AbrirBasePruebas()
	
}
