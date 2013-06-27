package ordenacion

import (
	"testing"
	_"fmt"
	"reflect"
)

//var os []OrdenacionSec = []OrdenacionSec{QuicksortSec1 {}, Mergesort1{}, ShellsortSec1{}, RadixSort1{}}
	//algoritmos de ordenación secuencial
//var op []OrdenacionParal = []OrdenacionParal{QuicksortParal1 {}, ParallellQuicksort1 {}, MergesortParallel{}, ShellsortParal1{}, RadixSortParalelo{}, ParallellSRegularSampling{}}


var os []OrdenacionSec = []OrdenacionSec{QuicksortSec1 {},Mergesort1{}, ShellsortSec1{}, RadixSort1{}}
	//algoritmos de ordenación secuencial
var op []OrdenacionParal = []OrdenacionParal{QuicksortParal1 {}, BitonicMergesortParallell{}, ShellsortParal1{}, RadixSortParalelo{}, ParallellSRegularSampling{}}

func TestOrdenacion(t *testing.T){
	var a []int
	var ordenado bool
	//fmt.Println(a)
	for _,o := range os {

		t.Log("Probando", reflect.TypeOf(o).Name())
		a = [] int{}
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		if !ordenado {
			t.Error("El array no está ordenado con tamaño 0")
		}
		a = [] int{1}
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		if !ordenado {
			t.Error("El array no está ordenado con tamaño 1")
		}
		a = CrearArrayAleatorio(101)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		if !ordenado {
			t.Error("El array no está ordenado con tamaño primo")
		}

		a = CrearArrayAleatorio(128)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		if !ordenado {
			t.Error("El array no está ordenado con tamaño potencia de 2")
		}

		a = CrearArrayAleatorio(1000)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		//fmt.Println(a)
		if !ordenado{
			t.Error("El array no esta ordenado con tamaño par") 
		}
		a = CrearArrayAleatorio(1001)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		//fmt.Println(a)
		if !ordenado{
			t.Error("El array no esta ordenado con tamaño impar") 
		}
	}
	for _,o := range op {
		a = [] int{}
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		t.Log("Probando", reflect.TypeOf(o).Name())
		if !ordenado {
			t.Error("El array no está ordenado con tamaño 0")
		}
		a = [] int{1}
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		if !ordenado {
			t.Error("El array no está ordenado con tamaño 1")
		}
		a = CrearArrayAleatorio(101)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		if !ordenado {
			t.Error("El array no está ordenado con tamaño primo")
		}

		a = CrearArrayAleatorio(128)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		if !ordenado {
			t.Error("El array no está ordenado con tamaño potencia de 2")
		}

		a = CrearArrayAleatorio(1000)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		//fmt.Println(a)
		if !ordenado{
			t.Error("El array no esta ordenado con tamaño par") 
		}
		a = CrearArrayAleatorio(1001)
		o.Ordenar(a)
		ordenado = EstaOrdenado(a)
		//fmt.Println(a)
		if !ordenado{
			t.Error("El array no esta ordenado con tamaño impar") 
		}

	}
}

func TestNumCPUS(t *testing.T) {
	for i := 0; i < 16; i++ {
		quicksortParal:=  QuicksortParal1 { i }
		shellsortParal:= ShellsortParal1 { i }
		radixsortParal:= RadixSortParalelo { i }
		bitonicMergesortParal:= BitonicMergesortParallell { i }
		psbrs:= ParallellSRegularSampling { i }
		histogramSort:= HistogramSort { i }
		var op []OrdenacionParal = []OrdenacionParal { quicksortParal, shellsortParal, radixsortParal, bitonicMergesortParal, psbrs, histogramSort }
		for _,o := range op {
			t.Logf("Algoritmo %s con %d CPUs", reflect.TypeOf(o).Name(), i)
			a := CrearArrayAleatorio(10007)
			o.Ordenar(a)
			ordenado := EstaOrdenado(a)
			if !ordenado{
				t.Errorf("El array no esta ordenado con algoritmo %s y numero CPUS %d", reflect.TypeOf(o).Name(), i) 
			}
		}
	}
}

// func TestMergesortParalelo(t *testing.T){
// 	var a []int
// 	a = CrearArrayAleatorio(1000)
// 	//a = []int {3, 10, 1, 7, 5, 6, 8, 4, 9 ,2}
// 	ordenaMergesortParalelo(a)
// 	ordenado := EstaOrdenado(a)
// 	fmt.Println(a)
// 	if !ordenado{
// 		t.Error("El array no esta ordenado") 
// 	}
// }

// func TestShellsortParalelo(t *testing.T){
// 	var a []int
// 	a = CrearArrayAleatorio(1000)
// 	//a = []int {3, 10, 1, 7, 5, 6, 8, 4, 9 ,2}
// 	ordenaShellsortParalelo1(a)
// 	ordenado := EstaOrdenado(a)
// 	fmt.Println(a)
// 	if !ordenado{
// 		t.Error("El array no esta ordenado") 
// 	}

// }

