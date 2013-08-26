package main


import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
	"usc/ordenacion/integers"
	_ "github.com/mattn/go-sqlite3"
	_ "os"
	"daos"
	"reflect"
	//"runtime"
)


func main() {
	//algoritmos de ordenación paralela enteros
	var NCPU int
 	var PSIntegers []integers.ParallelSort = []integers.ParallelSort { integers.QuickSortParallelized {NCPU}, 
 																	integers.ShellSortParallelized {NCPU},
 																	integers.BitonicMergeSortParallelized {NCPU},
 																	integers.RadixSortParallelized {NCPU}, 
 																	integers.ParallellSortRegularSampling {NCPU},
 																	integers.HistogramSort{NCPU}}


	var tamañosEntrada []int = []int{ 20000, 40000, 80000 }
	var nPruebas int = 100
	var test daos.Test
	NCPU = 1
	var results []daos.ResultadoTest = []daos.ResultadoTest{}
	var result daos.ResultadoTest
	for _, v := range tamañosEntrada{
		
		test.NCores = NCPU
		result.TamanhoEntrada = v
		test.Fecha = time.Now()
		test.Nombre =  fmt.Sprintf("Test paralelo array aleatorio tamaño entrada: %d NumeroCores: %d", v, test.NCores)
		fmt.Printf("**************************** size array == %d  ***********************************\n", v)
		a := ordenacion.CreateRandomArray(v)
		results= results[:0]
		b := make([]int, len(a))
		for _, w:= range PSIntegers{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				b:= ordenacion.CreateRandomArray(v)
				t1 := time.Now()
				w.Sort(b)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Is sorted? ",ordenacion.IsSorted(b))
		}
		daos.GuardarResultadosTest(&test, results) 
		test.Fecha = time.Now()
		test.Nombre = fmt.Sprintf("Test paralelo array ascendente tamaño entrada: %d NumeroCores: %d", v, test.NCores)

		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a = ordenacion.CreateAscendingArray(v)
		results= results[:0]
		b = make([]int, len(a))
		for _, w:= range PSIntegers{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				w.Sort(b)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Is sorted? ",ordenacion.IsSorted(b))
		}
	
		daos.GuardarResultadosTest(&test, results) 

		test.Fecha = time.Now()
		test.Nombre = fmt.Sprintf("Test paralelo array descendente tamaño entrada: %d NumeroCores: %d", v, test.NCores)
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a = ordenacion.CreateDescendingArray(v)
		results= results[:0]
		b = make([]int, len(a))
		for _, w:= range PSIntegers{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				w.Sort(b)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			fmt.Println("Is sorted? ",ordenacion.IsSorted(b))
		}
		daos.GuardarResultadosTest(&test, results) 
		//	//ordenacion.AbrirBasePruebas()
		NCPU *=2
	}
}
