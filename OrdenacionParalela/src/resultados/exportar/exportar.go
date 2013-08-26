package main

import (
"daos"
"os"
"fmt"
"strconv"
)

func VolcarDatos(nombre string, test *daos.Test, resultados []*daos.ResultadoTest) {
	ultimoTam := resultados[0].TamanhoEntrada
	os.Remove(nombre)
	fichero, err := os.Create(nombre)
	if err != nil {
		fmt.Println("Error abriendo el fichero ", nombre)
		os.Exit(-1)
	}
	first := true
	for _,v := range resultados {
		if first {
			first = !first
		} else {
			if v.TamanhoEntrada != ultimoTam {
				fichero.WriteString("\n")
				ultimoTam = v.TamanhoEntrada
			} else {
				fichero.WriteString(",")
			}
		}
		cadena := fmt.Sprintf("%d", v.Tiempo)
		fichero.WriteString(cadena)
	}
	fichero.Close()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ERROR: Es necesario indicar un id inicial.")
		os.Exit(-1)
	}
	idInicial, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: Es necesario indicar un id inicial.")
		os.Exit(-1)
	}

	// Datos para grafica de algoritmos secuenciales con enteros (aleatorio)
	test, resultados := daos.CargarResultadosTest(idInicial)
	VolcarDatos("secuencial_int_aleatorio.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (ascendente)
	test, resultados = daos.CargarResultadosTest(idInicial+1)
	VolcarDatos("secuencial_int_ascendente.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (descendente)
	test, resultados = daos.CargarResultadosTest(idInicial+2)
	VolcarDatos("secuencial_int_descendente.txt", test, resultados)

	// Datos para grafica de algoritmos secuenciales con enteros (aleatorio)
	test, resultados = daos.CargarResultadosTest(idInicial+3)
	VolcarDatos("paralelo_int_aleatorio_escal.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (ascendente)
	test, resultados = daos.CargarResultadosTest(idInicial+4)
	VolcarDatos("paralelo_int_ascendente_escal.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (descendente)
	test, resultados = daos.CargarResultadosTest(idInicial+5)
	VolcarDatos("paralelo_int_descendente_escal.txt", test, resultados)

	// Datos para grafica de algoritmos secuenciales con enteros (aleatorio)
	test, resultados = daos.CargarResultadosTest(idInicial+6)
	VolcarDatos("paralelo_int_aleatorio_2.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (ascendente)
	test, resultados = daos.CargarResultadosTest(idInicial+7)
	VolcarDatos("paralelo_int_ascendente_2.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (descendente)
	test, resultados = daos.CargarResultadosTest(idInicial+8)
	VolcarDatos("paralelo_int_descendente_2.txt", test, resultados)

		// Datos para grafica de algoritmos secuenciales con enteros (aleatorio)
	test, resultados = daos.CargarResultadosTest(idInicial+9)
	VolcarDatos("paralelo_int_aleatorio_4.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (ascendente)
	test, resultados = daos.CargarResultadosTest(idInicial+10)
	VolcarDatos("paralelo_int_ascendente_4.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (descendente)
	test, resultados = daos.CargarResultadosTest(idInicial+11)
	VolcarDatos("paralelo_int_descendente_4.txt", test, resultados)

		// Datos para grafica de algoritmos secuenciales con enteros (aleatorio)
	test, resultados = daos.CargarResultadosTest(idInicial+12)
	VolcarDatos("paralelo_int_aleatorio_8.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (ascendente)
	test, resultados = daos.CargarResultadosTest(idInicial+13)
	VolcarDatos("paralelo_int_ascendente_8.txt", test, resultados)
	// Datos para grafica de algoritmos secuenciales con enteros (descendente)
	test, resultados = daos.CargarResultadosTest(idInicial+14)
	VolcarDatos("paralelo_int_descendente_8.txt", test, resultados)
}
