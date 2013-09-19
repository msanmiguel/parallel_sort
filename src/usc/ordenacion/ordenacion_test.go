// María Sanmiguel Suarez. 2013

package ordenacion

import (
	"testing"
	_"fmt"
	"reflect"
	"usc/ordenacion/generic"
	"usc/ordenacion/integers"
	"usc/ordenacion/reflection"

)
//runtime.GOMAXPROCS(NCPU)
var NCPU  int = 2
//algoritmos de ordenación secuencial enteros
var SSIntegers []integers.SequentialSort = []integers.SequentialSort{integers.GolangSort{},
                                                                     integers.QuickSortSequential {},
                                                                     integers.InsertionSortSequential{}, 
                                                                     integers.BubbleSortSequential{},
                                                                     integers.ShellSortSequential{},
                                                                     integers.RadixSortSequential{}, 
                                                                     integers.BitonicMergeSortSequential{}}
	
	//algoritmos de ordenación paralela enteros
 	var PSIntegers []integers.ParallelSort = []integers.ParallelSort { integers.QuickSortParallelized {NCPU}, 
 																	integers.ShellSortParallelized {NCPU},
 																	integers.BitonicMergeSortParallelized {NCPU},
 																	integers.RadixSortParallelized {NCPU}, 
 																	integers.ParallellSortRegularSampling {NCPU},
 																	integers.HistogramSort{NCPU}}

	// Ordenacion secuencial Para slices genericos usando la api de reflexion
	var SSReflect []reflection.SequentialSort = []reflection.SequentialSort{ reflection.BubbleSortSequential {}, 
																	reflection.InsertionSortSequential {}, 
																	reflection.BitonicMergesortSequential {},
																	reflection.QuickSortSequential {},
																	reflection.ShellSortSequential {} }
 
	// Ordenacion Paralela para slices genéricos usando la api de reflexión
	var PSReflect []reflection.ParallelSort = []reflection.ParallelSort{reflection.QuickSortParallelized {NCPU}, 
																	reflection.ShellSortParallelized {NCPU}, 
																	reflection.BitonicMergeSortParallelized {NCPU}, 
																	reflection.ParallellSortRegularSampling {NCPU}}
	
	// algoritmos de ordenacion secuencial generalizados

	var SSgeneric []generic.SequentialSort = []generic.SequentialSort { generic.QuickSortSequential {},
																	 generic.ShellSortSequential {}, 
																	 generic.BubbleSortSequential {},
																	 generic.InsertionSortSequential {}, 
																	 generic.BitonicMergesortSequential { }}
	

 	var PSgeneric []generic.ParallelSort = []generic.ParallelSort {generic.BitonicMergeSortParallelized { NCPU }, 
 																	generic.QuickSortParallelized { NCPU }, 
 																	generic.ShellSortParallelized { NCPU } }
func TestOrdenacion(t *testing.T){
	var a []int
	var sorted bool
	for _,o := range  SSIntegers {

		t.Log("Probando", reflect.TypeOf(o).Name())
		a = [] int{}
		o.Sort(a)
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The empty array isn't ordered.")
		}
		a = [] int{1}
		o.Sort(a)
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The size 1 array isn't ordered.")
		}
		a = CreateRandomArray(101)
		o.Sort(a)
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The prime number sized array isn't ordered.")
		}

		a = CreateRandomArray(128)
		o.Sort(a)
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The power of two sized array isn't ordered.")
		}

		a = CreateRandomArray(1000)
		o.Sort(a)
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("El array no esta ordenado con tamaño par") 
		}
		a = CreateRandomArray(1001)
		o.Sort(a)
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("El array no esta ordenado con tamaño impar") 
		}
	}
	for _,o := range PSIntegers {
		a = [] int{}
		o.Sort(a)
		sorted = IsSorted(a)
		t.Log("Probando", reflect.TypeOf(o).Name())
		if !sorted {
			t.Error("The empty array isn't ordered.")
		}
		a = [] int{1}
		o.Sort(a)
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The size 1 array isn't ordered.")
		}
		a = CreateRandomArray(101)
		o.Sort(a)
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The prime number sized array isn't ordered.")
		}

		a = CreateRandomArray(128)
		o.Sort(a)
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The power of two sized array isn't ordered.")
		}

		a = CreateRandomArray(1000)
		o.Sort(a)
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The even sized array isn't ordered.") 
		}
		a = CreateRandomArray(1001)
		o.Sort(a)
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The odd sized array isn't ordered.") 
		}
		a = CreateAscendingArray(1001)
		o.Sort(a)
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The odd sized array isn't ordered.") 
		}
		a = CreateDescendingArray(1001)
		o.Sort(a)
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The odd sized array isn't ordered.") 
		}

	}
}

func TestNumCPUS(t *testing.T) {
	for i := 0; i < 16; i++ {
		
		var PSIntegers []integers.ParallelSort = []integers.ParallelSort { integers.QuickSortParallelized {i},
																		 integers.ShellSortParallelized {i}, 
																		 integers.BitonicMergeSortParallelized {i}, 
																		 integers.RadixSortParallelized {i}, 
																		 integers.ParallellSortRegularSampling {i}, 
																		 integers.HistogramSort{i}}
		for _,o := range  PSIntegers {
			t.Logf("Algoritmo %s con %d CPUs", reflect.TypeOf(o).Name(), i)
			a := CreateRandomArray(10007)
			o.Sort(a)
			sorted := IsSorted(a)
			if !sorted{
				t.Errorf("El array no esta ordenado con algoritmo %s y numero CPUS %d", reflect.TypeOf(o).Name(), i) 
			}
		}
	}
}


type IntegersComparator struct {}

func (i IntegersComparator) Compare(i1, i2 interface{}) int {
	v1 := i1.(int)
	v2 := i2.(int)
	return v1-v2
}

func TestGenericSort(t *testing.T){
	var a []int
	var sorted bool
	for _,s := range SSReflect {
		t.Log("Probando", reflect.TypeOf(s).Name())
		a = [] int{}
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The empty array isn't ordered.")
		}
		a = [] int{1}
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The size 1 array isn't ordered.")
		}
		a = CreateRandomArray(101)
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The prime number sized array isn't ordered.")
		}

		a = CreateRandomArray(128)
		s.Sort(a, IntegersComparator{} )
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The power of two sized array isn't ordered.")
		}

		a = CreateRandomArray(1000)
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The even sized array isn't ordered.") 
		}
		a = CreateRandomArray(1001)
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The odd sized array isn't ordered.") 
		}
	}

	for _,s := range SSReflect {
		t.Log("Probando", reflect.TypeOf(s).Name())
		a = [] int{}
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The empty array isn't ordered.")
		}
		a = [] int{1}
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The size 1 array isn't ordered.")
		}
		a = CreateRandomArray(101)
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The prime number sized array isn't ordered.")
		}

		a = CreateRandomArray(128)
		s.Sort(a, IntegersComparator{} )
		sorted = IsSorted(a)
		if !sorted {
			t.Error("The power of two sized array isn't ordered.")
		}

		a = CreateRandomArray(1000)
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The even sized array isn't ordered.") 
		}
		a = CreateRandomArray(1001)
		s.Sort(a, IntegersComparator{})
		sorted = IsSorted(a)
		//fmt.Println(a)
		if !sorted{
			t.Error("The odd sized array isn't ordered.") 
		}
	}
}

// func TestMergesortParalelo(t *testing.T){
// 	var a []int
// 	a = CreateRandomArray(1000)
// 	//a = []int {3, 10, 1, 7, 5, 6, 8, 4, 9 ,2}
// 	ordenaMergesortParalelo(a)
// 	sorted := IsSorted(a)
// 	fmt.Println(a)
// 	if !sorted{
// 		t.Error("El array no esta ordenado") 
// 	}
// }

// func TestShellsortParalelo(t *testing.T){
// 	var a []int
// 	a = CreateRandomArray(1000)
// 	//a = []int {3, 10, 1, 7, 5, 6, 8, 4, 9 ,2}
// 	ordenaShellsortParalelo1(a)
// 	sorted := IsSorted(a)
// 	fmt.Println(a)
// 	if !sorted{
// 		t.Error("El array no esta ordenado") 
// 	}

// }

