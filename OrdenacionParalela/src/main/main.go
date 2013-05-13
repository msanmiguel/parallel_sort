package main

import (
	_ "database/sql"
	"fmt"
	"time"
	"usc/ordenacion"
	_ "github.com/mattn/go-sqlite3"
	_ "os"
)

func main() {

	//a:=[]int{2,54,7,98,56,51,0,23}
	//a := ordenacion.CrearArrayDescendente(100000)
	//a := ordenacion.CrearArrayAleatorio(10000)

	//a := ordenacion.CrearArrayAscendente(100000)

	//algoritmos de ordenación secuencial
	var os []ordenacion.OrdenacionSec = []ordenacion.OrdenacionSec{new(ordenacion.QuicksortSec1), new(ordenacion.ShellsortSec1),
	 new(ordenacion.Burbuja1), new(ordenacion.Insercion1), new(ordenacion.Mergesort1)}
	//algoritmos de ordenación secuencial
	var op []ordenacion.OrdenacionParal = []ordenacion.OrdenacionParal{new(ordenacion.QuicksortParal1), new(ordenacion.ShellsortParal1)}
	var tamañosEntrada []int = []int{10, 100, 1000, 10000}
	var nPruebas int64 = 5
	for _, v := range tamañosEntrada{
		fmt.Printf("****************************Tamaño entrada == %d  ***********************************\n", v)
		a := ordenacion.CrearArrayAleatorio(v)
		
		b := make([]int, len(a))
		for _, v:= range os{
			var sumaT int64 = 0
			var j int64 = 0
			for j = 0 ;j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Ordenar(b)
				t2 := time.Now()
				duracion := t2.Sub(t1)
				tiempo := duracion.Nanoseconds()
				sumaT =  sumaT + tiempo 
			}
			fmt.Printf("alogoritmo == %s   tiempo== %d \n", v.ObtenerNombreAlgoritmo() ,sumaT/(j+1))
			fmt.Println("Esta ordenado? ",ordenacion.EstaOrdenado(b))
		}
		for _, v:= range op{
			var sumaT int64 = 0
			var j int64 = 0
			for j = 0; j < nPruebas; j++{ //Se hacen n ejecuciones para calcular el promedio
				copy(b, a)
				t1 := time.Now()
				v.Ordenar(b)
				t2 := time.Now()
				duracion := t2.Sub(t1)
				tiempo := duracion.Nanoseconds()
				sumaT =  sumaT + tiempo
			}
			fmt.Printf("alogoritmo == %s   tiempo== %d \n", v.ObtenerNombreAlgoritmo() ,sumaT/(j+1))
			fmt.Println("Esta ordenado? ",ordenacion.EstaOrdenado(b))
		} 	
	}
	
	
	//	//ordenacion.AbrirBasePruebas()

}
