package ordenacion

import(
	_"fmt"
	"reflect"
	_"sort"
)


type Comparator interface{
	Compare(i1, i2 interface{})int
}


// Burbuja secuencial generalizado
func bubbleSortSecuential(i interface{}, c Comparator){
	k:=0
	j:=0
	v := reflect.ValueOf(i)
	for k=2; k<=v.Len(); k++{
		for j=0; j<=v.Len()-k; j++{
			v1 := v.Index(j)
			v2 := v.Index(j+1)
			if c.Compare(v1.Interface(), v2.Interface())>0{
				GenericSwap(v1,v2)
			}
		}
	}
}


// InsertionSort secuencial generalizado

func insertionSortSecuential(i interface{}, c Comparator){
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i).Elem()
	p := reflect.New(t).Elem()
	for i:=1; i < v.Len();i++{
		p.Set(v.Index(i))
		j := i-1
		for ; j>=0 && c.Compare(v.Index(j).Interface(), p.Interface())>0; j--{
			v.Index(j+1).Set(v.Index(j))
		}
		v.Index(j+1).Set(p)
	}
}

// Quicksort secuencial generalizado


func quickSortSecuential(i interface{}, c Comparator){
	v := reflect.ValueOf(i)
	if v.Len() > 20 {
		pos_pivote := genericReposition(i, c)
		quickSortSecuential(v.Slice(0,pos_pivote).Interface(), c) // recoloco la lista de los menores
		quickSortSecuential(v.Slice((pos_pivote+1),v.Len()).Interface(), c) // recoloco la lista de los mayores
	} else if  v.Len() > 1 {
		insertionSortSecuential(i, c)
	}
} 


// la funcion recolocar devuelve la lista recolocada y la posición en la que está el pivote
func genericReposition(i interface{}, c Comparator) int {
	var izquierdo int
	var derecho int
    	var pivote reflect.Value 
	v := reflect.ValueOf(i)	
	n := v.Len()
	c1 := v.Index(0)
	c2 := v.Index((n-1)/2)
	if c.Compare(c1.Interface(), c2.Interface()) > 0{
		GenericSwap(c1,c2)
	}
	c3 := v.Index(n-1)
	if c.Compare(c2.Interface(), c3.Interface()) > 0{
		 GenericSwap(c2 ,c3)

	}
	if c.Compare(c1.Interface(), c2.Interface()) <= 0 {
		GenericSwap(c1,c2)
	}
	pivote = v.Index(0)

	izquierdo = 0
	derecho = n-1

	// Hasta que los dos indices se crucen 
	for izquierdo < derecho {
		for c.Compare(v.Index(derecho).Interface(), pivote.Interface()) > 0{
			derecho--
		}
		for izquierdo < n && c.Compare(v.Index(izquierdo).Interface(),pivote.Interface()) <= 0{
			izquierdo++
		} 
		// si todavia no se cruzan los indices intercambiamos 
		if izquierdo < derecho {
			GenericSwap(v.Index(izquierdo),v.Index(derecho))
		}
	}
	// cuando se cruzaron los indices se coloca el pivote en el lugar que le corresponde
	GenericSwap(v.Index(derecho),v.Index(0))
	// se devuelve el la lista recolocada y la nueva posición del pivote
	return derecho
}



// Mergesort secuencial generalizado


func mergeSortSecuential(i interface{} , c Comparator, direction bool){
	v := reflect.ValueOf(i)
	s := v.Len()
	if s <= 1{
		return;
	}
	mergeSortSecuential(v.Slice(0,(s/2)).Interface(), c, !direction)
	mergeSortSecuential(v.Slice((s/2),s).Interface(), c, direction)
	genericMerge(v.Interface(), direction, c)

}

func genericMerge(i interface{}, direction bool,  c Comparator){
	v := reflect.ValueOf(i)
	s :=  v.Len()
	m := potenciaDe2(s)
	genericHalfclean(i, m, direction, c)
	genericBisort(v.Slice(0,m).Interface(), direction, c)
	genericBisort(v.Slice(m,s).Interface(), direction, c)

}
func genericBisort(i interface{}, direction bool,  c Comparator){
	v := reflect.ValueOf(i)
	if v.Len() == 1{
		return
	}
	s := v.Len()
	m:= potenciaDe2(s)
	genericHalfclean(i, m, direction, c)

	genericBisort(v.Slice(0,m).Interface(), direction, c)
	genericBisort(v.Slice(m,s).Interface(), direction, c)
}

func genericHalfclean (i interface{}, m int, direction bool, c Comparator){
	v := reflect.ValueOf(i)
	s := v.Len()
	for k := 0; k < s-m ; k++{
		genericCompSwitch(i, k, k+m, direction, c)
	}

}


func genericCompSwitch(i interface{}, k, j int, direction bool, c Comparator){
	v := reflect.ValueOf(i)
	e := v.Index(k)
	b := v.Index(j)
	
	if direction == (c.Compare(b.Interface(), e.Interface())<0) {
		GenericSwap(b,e)
	}
}

func shellSortSecuential(in interface{}, c Comparator){
	v := reflect.ValueOf(in)
	salto:= v.Len()/2
	t := reflect.TypeOf(in).Elem()
    p := reflect.New(t).Elem()
	for salto >= 1 {
		for k := 0; k < salto; k++ {
			for i := k+salto; i < v.Len(); i += salto {
                p.Set(v.Index(i))
				j := i-salto
				for ; j >= 0 && c.Compare(v.Index(j).Interface(), p.Interface()) > 0; j-=salto{
					v.Index(j+salto).Set(v.Index(j))
				}
				v.Index(j+salto).Set(p)
			}
		}
		salto=salto/2
	}
}

type SecuentialSort interface{
	Sort(i interface {}, c Comparator)
}


type BubbleSortSec struct{}

type InsertionSortSec struct{}

type MergeSortSec struct{}

type QuickSortSec struct{}

type ShellSortSec struct{}


func (o BubbleSortSec) Sort(i interface{}, c Comparator){
	bubbleSortSecuential(i, c)
}
func (o InsertionSortSec) Sort(i interface{}, c Comparator){
	insertionSortSecuential(i, c)
}
func (o MergeSortSec) Sort(i interface{}, c Comparator){
	mergeSortSecuential(i, c, true)
}

func (o QuickSortSec) Sort(i interface{}, c Comparator){
	quickSortSecuential(i, c)
}

func (o ShellSortSec) Sort(i interface{}, c Comparator){
	shellSortSecuential(i, c)
}

