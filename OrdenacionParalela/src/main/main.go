package main

import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
	_ "github.com/mattn/go-sqlite3"
	_ "os"
	"daos"
	"reflect"
	//"runtime"
)

func main() {
	//NCPU := runtime.NumCPU()
	//runtime.GOMAXPROCS(NCPU)
	NCPU := 2

	//a:=[]int{2,54,7,98,56,51,0,23}
	//a := ordenacion.CrearArrayDescendente(100000)
	//a := ordenacion.CrearArrayAleatorio(10000)
	//a := ordenacion.CrearArrayAscendente(100000)
	
	//algoritmos de ordenación secuencial

	

	var os []ordenacion.OrdenacionSec = []ordenacion.OrdenacionSec{ordenacion.QuicksortSec1 {}, ordenacion.Mergesort1{}, ordenacion.ShellsortSec1{}, ordenacion.RadixSort1{}}
	//algoritmos de ordenación paralela
	quicksortParal:=  ordenacion.QuicksortParal1 { NCPU }
	shellsortParal:= ordenacion.ShellsortParal1 { NCPU }
	radixsortParal:= ordenacion.RadixSortParalelo { NCPU }
	bitonicMergesortParal:= ordenacion.BitonicMergesortParallell { NCPU }
	psbrs:= ordenacion.ParallellSRegularSampling { NCPU }
	histogramSort:= ordenacion.HistogramSort { NCPU }
	
	//var op []ordenacion.OrdenacionParal = []ordenacion.OrdenacionParal{ordenacion.QuicksortParal1 {}, ordenacion.ParallellQuicksort1 {}, ordenacion.BitonicMergesortParallell{}, 
	//ordenacion.ShellsortParal1{}, ordenacion.RadixSortParalelo{}, ordenacion.ParallellSRegularSampling{}}

	var op []ordenacion.OrdenacionParal = []ordenacion.OrdenacionParal { quicksortParal, shellsortParal, radixsortParal, bitonicMergesortParal, psbrs, histogramSort }

	var tamañosEntrada []int = []int{ 1000, 10000, 100000}
	var nPruebas int = 5
	var test daos.Test
	test.IdTest = 0
	test.NCores = 2
	test.Fecha = time.Now()
	test.Nombre = "Comparativa algoritmos Quicksort, Shellsort , burbuja, insercion, mergesort secuencial"
	var resultados []daos.ResultadoTest = []daos.ResultadoTest{}
	var resultado daos.ResultadoTest
	resultado.IdResultado = 0
	for _, v := range tamañosEntrada{
		resultado.TamanhoEntrada = v
		fmt.Printf("****************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CrearArrayAleatorio(v)
		
		b := make([]int, len(a))
		for _, v:= range os{
			resultado.Algoritmo = reflect.TypeOf(v).Name()
			var sumaT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Ordenar(b)
				t2 := time.Now()
				duracion := t2.Sub(t1)
				tiempo := int(duracion.Nanoseconds())
				sumaT =  sumaT + tiempo 
			}
			resultado.Tiempo = sumaT/(j+1)
			resultados = append(resultados, resultado)
			fmt.Println("Esta ordenado? ",ordenacion.EstaOrdenado(b))
		}
		for _, v:= range op{
			resultado.Algoritmo = reflect.TypeOf(v).Name()
			var sumaT int = 0
			var j int = 0
			for j = 0; j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Ordenar(b)
				t2 := time.Now()
				duracion := t2.Sub(t1)
				tiempo := int(duracion.Nanoseconds())
				sumaT =  sumaT + tiempo
			}
			resultado.Tiempo = sumaT/(j+1)
			resultados = append(resultados, resultado)
		} 	
	}
	daos.GuardarResultadosTest(&test, resultados) 
	
	
	//	//ordenacion.AbrirBasePruebas()

}
