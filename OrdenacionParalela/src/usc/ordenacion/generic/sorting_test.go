// María Sanmiguel Suarez. 2013

package generic

import (
"testing"
"reflect"
"usc/ordenacion"
)

var ss []SequentialSort = []SequentialSort { QuickSortSequential {}, ShellSortSequential {}, BubbleSortSequential {}, InsertionSortSequential {}, BitonicMergesortSequential {}, BitonicMergesortSequential {} }

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

func TestOrdenacion(t *testing.T){
	is := IntSlice {}
	var ordered bool
	for _,o := range ss {

		t.Log("Probando", reflect.TypeOf(o).Name())
		is.a = [] int{}
		o.Sort(is)
		ordered = ordenacion.IsSorted(is.a)
		if !ordered {
			t.Error("El array no está ordenado con tamaño 0")
		}
		is.a = [] int{1}
		o.Sort(is)
		ordered = ordenacion.IsSorted(is.a)
		if !ordered {
			t.Error("El array no está ordenado con tamaño 1")
		}
		is.a = ordenacion.MakeRandomArray(101)
		o.Sort(is)
		ordered= ordenacion.IsSorted(is.a)
		if !ordered {
			t.Error("El array no está ordenado con tamaño primo")
		}

		is.a = ordenacion.MakeRandomArray(128)
		o.Sort(is)
		ordered = ordenacion.IsSorted(is.a)
		if !ordered {
			t.Error("El array no está ordenado con tamaño potencia de 2")
		}

		is.a = ordenacion.MakeRandomArray(1000)
		o.Sort(is)
		ordered = ordenacion.IsSorted(is.a)
		//fmt.Println(a)
		if !ordered{
			t.Error("El array no esta ordenado con tamaño par") 
		}
		is.a = ordenacion.MakeRandomArray(1001)
		o.Sort(is)
		ordered = ordenacion.IsSorted(is.a)
		//fmt.Println(a)
		if !ordered{
			t.Error("El array no esta ordenado con tamaño impar") 
		}
	}
}
