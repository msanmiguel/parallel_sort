package main


import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
	"usc/ordenacion/reflection"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"daos"
	"reflect"
	"strconv"

)
type IntegerComparator struct {}

func (i IntegerComparator) Compare(i1, i2 interface{}) int {
	v1 := i1.(int)
	v2 := i2.(int)
	return v1-v2
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
	// Ordenacion Paralela para slices genéricos usando la api de reflexión
	var PSReflect []reflection.ParallelSort = []reflection.ParallelSort{reflection.QuickSortParallelized {NCPU}, 
																		reflection.ShellSortParallelized {NCPU},
																		reflection.BitonicMergeSortParallelized {NCPU}, 
																		reflection.ParallellSortRegularSampling {NCPU}}
	

	var tamañosEntrada []int = []int{ 50000, 100000, 150000, 200000, 250000, 300000 }
	var nPruebas int = 100
	var test daos.Test
	test.NCores = NCPU
	test.Fecha = time.Now()
	test.Nombre = "Test paralelo con 2 cores con arrays aleatorios"
	var results []daos.ResultadoTest = []daos.ResultadoTest{}
	var result daos.ResultadoTest
		for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("**************************** size array == %d  ***********************************\n", v)
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
			results = append(results, result)		}
	}
	daos.GuardarResultadosTest(&test, results) 
	test.Fecha = time.Now()
	test.Nombre = "Test paralelo con 2 cores arrays ascendentes"
	results= results[:0]
	for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CreateAscendingArray(v)
		b := make([]int, len(a))
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
	}
	daos.GuardarResultadosTest(&test, results) 

	test.Fecha = time.Now()
	test.Nombre = "Test paralelo con 2 cores con arrays descendentes"
	results= results[:0]
	for _, v := range tamañosEntrada{
		result.TamanhoEntrada = v
		fmt.Printf("*************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CreateDescendingArray(v)
		
		b := make([]int, len(a))
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
	}
	daos.GuardarResultadosTest(&test, results) 
	//	//ordenacion.AbrirBasePruebas()
}
