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
	//algoritmos de ordenación secuencial enteros
	var SSIntegers []integers.SequentialSort = []integers.SequentialSort{integers.GolangSort{}, 
																		integers.QuickSortSequential {},
																		integers.InsertionSortSequential{},
																		integers.BubbleSortSequential{}, 
																		integers.ShellSortSequential{},
																		integers.RadixSortSequential{},
																		integers.BitonicMergeSortSequential{}}
	
	

	var tamañosEntrada []int = []int{ 2000, 4000, 6000, 8000 }
	var nPruebas int = 100
	var test daos.Test
	test.NCores = 1
	test.Fecha = time.Now()
	test.Nombre = "Test secuencial con arrays aleatorios"
	var results []daos.ResultadoTest = []daos.ResultadoTest{}
	var result daos.ResultadoTest
		for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("****************************Tamaño entrada == %d  ***********************************\n", v)
		for _, w:= range SSIntegers{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				b := ordenacion.CreateRandomArray(v)
				t1 := time.Now()
				w.Sort(b)
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
			//fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		}
	}
	daos.GuardarResultadosTest(&test, results) 
	test.Fecha = time.Now()
	test.Nombre = "Test secuencial con arrays ascendentes"
	results= results[:0]
	for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CreateAscendingArray(v)
		
		b := make([]int, len(a))
		for _, w:= range SSIntegers{
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
			//fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		}
	}
	daos.GuardarResultadosTest(&test, results) 
	results= results[:0]
	test.Fecha = time.Now()
	test.Nombre = "Test secuencial con arrays descendentes"
	for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CreateDescendingArray(v)
		
		b := make([]int, len(a))
		for _, w:= range SSIntegers{
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
			//fmt.Println("Esta ordenado? ",ordenacion.IsSorted(b))
		}
	}
	daos.GuardarResultadosTest(&test, results) 
	//	//ordenacion.AbrirBasePruebas()
}
