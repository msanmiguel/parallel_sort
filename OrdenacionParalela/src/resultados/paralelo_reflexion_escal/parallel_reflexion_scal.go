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
	var NCPU int = 2
	// Ordenacion Paralela para slices genéricos usando la api de reflexión
	var PSReflect []reflection.ParallelSort = []reflection.ParallelSort{reflection.QuickSortParallelized {NCPU}, 
																		reflection.ShellSortParallelized {NCPU},
																		reflection.BitonicMergeSortParallelized {NCPU}, 
																		reflection.ParallellSortRegularSampling {NCPU}}
	
	var tamañosEntrada []int = []int{1000, 2000, 4000, 8000}
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
		for _, w:= range PSReflect{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)" 
			var sumT int = 0
			var j int = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				b:= ordenacion.CreateRandomArray(v)
				t1 := time.Now()
				w.Sort(b, IntegerComparator{})
				t2 := time.Now()
				duration := t2.Sub(t1)
				time := int(duration.Nanoseconds())
				sumT =  sumT + time 
			}
			result.Tiempo = sumT/(j+1)
			results = append(results, result)
		}
		daos.GuardarResultadosTest(&test, results) 
		test.Fecha = time.Now()
		test.Nombre = fmt.Sprintf("Test paralelo array ascendente tamaño entrada: %d NumeroCores: %d", v, test.NCores)

		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a = ordenacion.CreateAscendingArray(v)
		results= results[:0]
		b = make([]int, len(a))
		for _, w:= range PSReflect{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
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
		}
	
		daos.GuardarResultadosTest(&test, results) 

		test.Fecha = time.Now()
		test.Nombre = fmt.Sprintf("Test paralelo array descendente tamaño entrada: %d NumeroCores: %d", v, test.NCores)
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a = ordenacion.CreateDescendingArray(v)
		results= results[:0]
		b = make([]int, len(a))
		for _, w:= range PSReflect{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
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
		}
		daos.GuardarResultadosTest(&test, results) 
		//	//ordenacion.AbrirBasePruebas()
		NCPU *=2
	}
}
