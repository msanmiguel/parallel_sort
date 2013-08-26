package main


import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
	"usc/ordenacion/generic"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"daos"
	"reflect"
	"strconv"

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
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Número de parámetros incorrecto.")
		fmt.Println("Uso: parallel_generic_ncpu_const <NCPU>")
		os.Exit(-1)
	}
	NCPU, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: El parámetro NCPU no es un entero.")
		os.Exit(-1)
	}
	var PSgeneric []generic.ParallelSort = []generic.ParallelSort { generic.BitonicMergeSortParallelized { NCPU }, 
 																	generic.QuickSortParallelized { NCPU },
 																	generic.ShellSortParallelized { NCPU } }


	var tamañosEntrada []int = []int{ 100 , 1000, 10000 }
	var nPruebas int = 100
	var test daos.Test
	test.NCores = NCPU
	test.Fecha = time.Now()
	test.Nombre = fmt.Sprintf("Test paralelo con %d cores con arrays aleatorios", test.NCores)
	var results []daos.ResultadoTest = []daos.ResultadoTest{}
	var result daos.ResultadoTest
		for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("**************************** size array == %d  ***********************************\n", v)
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
			results = append(results, result)		}
	}
	daos.GuardarResultadosTest(&test, results) 
	test.Fecha = time.Now()
	test.Nombre = fmt.Sprintf("Test paralelo con %d cores arrays ascendentes", test.NCores)
	results= results[:0]
	for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CreateAscendingArray(v)
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
	}
	daos.GuardarResultadosTest(&test, results) 

	test.Fecha = time.Now()
	test.Nombre = fmt.Sprintf("Test paralelo con %d cores con arrays descendentes", test.NCores)
	results= results[:0]
	for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CreateDescendingArray(v)
		
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
	}
	daos.GuardarResultadosTest(&test, results) 
	//	//ordenacion.AbrirBasePruebas()
}
