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
	var NCPU int = 2
	var PSgeneric []generic.ParallelSort = []generic.ParallelSort { generic.BitonicMergeSortParallelized { NCPU }, 
 																	generic.QuickSortParallelized { NCPU },
 																	generic.ShellSortParallelized { NCPU } }


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
		for _, w:= range PSgeneric{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)" 
			var sumT int = 0
			var j int = 0
			is := IntSlice { }
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
		}
		daos.GuardarResultadosTest(&test, results) 
		test.Fecha = time.Now()
		test.Nombre = fmt.Sprintf("Test paralelo array ascendente tamaño entrada: %d NumeroCores: %d", v, test.NCores)

		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a = ordenacion.CreateAscendingArray(v)
		results= results[:0]
		b := make([]int, len(a))
		for _, w:= range PSgeneric{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
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
		}
	
		daos.GuardarResultadosTest(&test, results) 

		test.Fecha = time.Now()
		test.Nombre = fmt.Sprintf("Test paralelo array descendente tamaño entrada: %d NumeroCores: %d", v, test.NCores)
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a = ordenacion.CreateDescendingArray(v)
		results= results[:0]
		b = make([]int, len(a))
		for _, w:= range PSgeneric{
			result.Algoritmo = reflect.TypeOf(w).Name() +" (integers)"
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
		}
		daos.GuardarResultadosTest(&test, results) 
		//	//ordenacion.AbrirBasePruebas()
		NCPU *=2
	}
}
