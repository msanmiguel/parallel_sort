package main


import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
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

func main() {
	// Ordenacion secuencial Para slices genericos usando la api de reflexion
	var SSReflect []reflection.SequentialSort = []reflection.SequentialSort{ reflection.BubbleSortSequential {},
																	reflection.InsertionSortSequential {}, 
																	reflection.BitonicMergesortSequential {},
																	reflection.QuickSortSequential {}, 
																	reflection.ShellSortSequential {} }
 
	

	var tamañosEntrada []int = []int{ 100 , 1000, 10000 }
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
		for _, w:= range SSReflect{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (reflexion)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				b := ordenacion.CreateRandomArray(v)
				t1 := time.Now()
				w.Sort(b, IntegerComparator{})
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
		for _, w:= range SSReflect{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (reflexion)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				w.Sort(b, IntegerComparator{})
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
		for _, w:= range SSReflect{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (reflexion)"
			fmt.Println(reflect.TypeOf(w).Name())
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				w.Sort(b, IntegerComparator{})
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
