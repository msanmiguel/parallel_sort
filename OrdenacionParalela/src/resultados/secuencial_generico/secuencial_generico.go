package main


import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
	"usc/ordenacion/generic"
	_ "github.com/mattn/go-sqlite3"
	_ "os"
	"daos"
	"reflect"
	//"runtime"
)
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
	// algoritmos de ordenacion secuencial generalizados

	var SSgeneric []generic.SequentialSort = []generic.SequentialSort { generic.QuickSortSequential {}, 
																		generic.ShellSortSequential {}, 
																		// generic.BubbleSortSequential {}, 
																		// generic.InsertionSortSequential {}, 
																		generic.BitonicMergesortSequential {}}
	
	

	var tamañosEntrada []int = []int{ 50000, 100000, 150000, 200000, 250000, 300000 }
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
		for _, w:= range SSgeneric{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (generics)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			is := IntSlice {}
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				is.a = ordenacion.CreateRandomArray(v)
				t1 := time.Now()
				w.Sort(is)
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
		for _, w:= range SSgeneric{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (generics)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			is := IntSlice { b }
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(is.a, a)
				t1 := time.Now()
				w.Sort(is)
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
		for _, w:= range SSgeneric{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (generics)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			is := IntSlice { b }
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(is.a, a)
				t1 := time.Now()
				w.Sort(is)
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
