package main

import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
	"usc/ordenacion/generic"
	"usc/ordenacion/integers"
	"usc/ordenacion/reflection"
	_ "github.com/mattn/go-sqlite3"
	_ "os"
	"daos"
	"reflect"
	//"runtime"
)


type IntegerComparator struct {}

func (i IntegerComparator) Compare(i1, i2 interface{}) int {
	v1 := i1.(int)
	v2 := i2.(int)
	return v1-v2
}

type IntSlice struct {
	a []int
}

func (is IntSlice) Length() int {
	return len(is.a)
}

func (is IntSlice) Swap(i, j int) {
	is.a[i], is.a[j] = is.a[j], is.a[i]
}

func (is IntSlice) Compare(i, j int) int {
	return is.a[i]-is.a[j]
}

func main() {
	//NCPU := runtime.NumCPU()
	//runtime.GOMAXPROCS(NCPU)
	NCPU := 2

	//a:=[]int{2,54,7,98,56,51,0,23}
	//a := ordenacion.CrearArrayDescendente(100000)
	//a := ordenacion.CrearArrayAleatorio(10000)integers
	//a := ordenacion.CrearArrayAscendente(100000)
	
	//algoritmos de ordenación secuencial enteros
	var SSIntegers []integers.SequentialSort = []integers.SequentialSort{integers.GolangSort{}, 
																		integers.QuickSortSequential {},
																		integers.InsertionSortSequential{},
																		integers.BubbleSortSequential{}, 
																		integers.ShellSortSequential{},
																		integers.RadixSortSequential{},
																		integers.BitonicMergeSortSequential{}}
	
	//algoritmos de ordenación paralela enteros
 	var PSIntegers []integers.ParallelSort = []integers.ParallelSort { integers.QuickSortParallelized {NCPU}, 
 																	integers.ShellSortParallelized {NCPU},
 																	integers.BitonicMergeSortParallelized {NCPU}, 
 																	integers.RadixSortParallelized {NCPU},
 																	integers.ParallellSortRegularSampling {NCPU},
 																	integers.HistogramSort{NCPU}}


	// Ordenacion secuencial Para slices genericos usando la api de reflexion
	var SSReflect []reflection.SequentialSort = []reflection.SequentialSort{ reflection.BubbleSortSequential {},
																	reflection.InsertionSortSequential {}, 
																	reflection.BitonicMergesortSequential {},
																	reflection.QuickSortSequential {}, 
																	reflection.ShellSortSequential {} }
 
	// Ordenacion Paralela para slices genéricos usando la api de reflexión
	var PSReflect []reflection.ParallelSort = []reflection.ParallelSort{reflection.QuickSortParallelized {NCPU}, 
																		reflection.ShellSortParallelized {NCPU},
																		reflection.BitonicMergeSortParallelized {NCPU}, 
																		reflection.ParallellSortRegularSampling {NCPU}}
	
	// algoritmos de ordenacion secuencial generalizados

	var SSgeneric []generic.SequentialSort = []generic.SequentialSort { generic.QuickSortSequential {}, 
																		generic.ShellSortSequential {}, 
																		generic.BubbleSortSequential {}, 
																		generic.InsertionSortSequential {}, 
																		generic.BitonicMergesortSequential {}}
	

 	var PSgeneric []generic.ParallelSort = []generic.ParallelSort { generic.BitonicMergeSortParallelized { NCPU }, 
 																	generic.QuickSortParallelized { NCPU },
 																	generic.ShellSortParallelized { NCPU } }



	var tamañosEntrada []int = []int{ 100 , 1000 }
	var nPruebas int = 5
	var test daos.Test
	test.IdTest = 0
	test.NCores = 2
	test.Fecha = time.Now()
	test.Nombre = "Comparativa algoritmos Quicksort, Shellsort , burbuja, insercion, mergesort secuencial"
	var results []daos.ResultadoTest = []daos.ResultadoTest{}
	var result daos.ResultadoTest
	result.IdResultado = 0
	for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("****************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CreateRandomArray(v)
		
		b := make([]int, len(a))
		for _, v:= range SSIntegers{
			result.Algoritmo = reflect.TypeOf(v).Name()
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Sort(b)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		}
		for _, v:= range PSIntegers{
			result.Algoritmo = reflect.TypeOf(v).Name()
			var sumT int = 0
			var j int = 0
			for j = 0; j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Sort(b)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
		}
		for _, v:= range SSReflect{
			result.Algoritmo = reflect.TypeOf(v).Name()
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Sort(b, IntegerComparator{})
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		}
		for _, v:= range PSReflect{
			result.Algoritmo = reflect.TypeOf(v).Name()
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Sort(b, IntegerComparator{})
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		}
		for _, v:= range SSgeneric {
			result.Algoritmo = reflect.TypeOf(v).Name()
			var sumT int = 0
			var j int = 0
			is := IntSlice { b }
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(is.a, a)
				t1 := time.Now()
				v.Sort(is)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		} 
		for _, v:= range PSgeneric{
			result.Algoritmo = reflect.TypeOf(v).Name()
			var sumT int = 0
			var j int = 0
			is := IntSlice { b }
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(is.a, a)
				t1 := time.Now()
				v.Sort(is)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		}	
	}
	daos.GuardarResultadosTest(&test, results) 
	
	
	//	//ordenacion.AbrirBasePruebas()

}
