// María Sanmiguel Suarez. 2013

package reflection

import(
	_"fmt"
	"reflect"
	_"sort"
)

// Interface to be implemented to provide a comparator between values
// of the slice.
type Comparator interface{
	// Compare function to be used by the sorting algorithms of this package.
	// The returned value of must be:
	// <=0 if i1 < i2,
	// = 0 if i1 == i2,
	// >=0 if i1 > i2.
	Compare(i1, i2 interface{})int
}


// Burbuja secuencial generalizado
func bubbleSortSequential(i interface{}, c Comparator){
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

func insertionSortSequential(i interface{}, c Comparator){
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


func quickSortSequential(i interface{}, c Comparator){
	v := reflect.ValueOf(i)
	if v.Len() > 20 {
		pivot_pos := reposition(i, c)
		quickSortSequential(v.Slice(0,pivot_pos).Interface(), c) // recoloco la lista de los menores
		quickSortSequential(v.Slice((pivot_pos+1),v.Len()).Interface(), c) // recoloco la lista de los mayores
	} else if  v.Len() > 1 {
		insertionSortSequential(i, c)
	}
} 


// la funcion recolocar devuelve la lista recolocada y la posición en la que está el pivote
func reposition(i interface{}, c Comparator) int {
	var left int
	var right int
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

	left = 0
	right = n-1

	// Hasta que los dos indices se crucen 
	for left < right {
		for c.Compare(v.Index(right).Interface(), pivote.Interface()) > 0{
			right--
		}
		for left < n && c.Compare(v.Index(left).Interface(),pivote.Interface()) <= 0{
			left++
		} 
		// si todavia no se cruzan los indices intercambiamos 
		if left < right {
			GenericSwap(v.Index(left),v.Index(right))
		}
	}
	// cuando se cruzaron los indices se coloca el pivote en el lugar que le corresponde
	GenericSwap(v.Index(right),v.Index(0))
	// se devuelve el la lista recolocada y la nueva posición del pivote
	return right
}



// Mergesort secuencial generalizado


func mergeSortSequential(i interface{} , c Comparator, direction bool){
	v := reflect.ValueOf(i)
	s := v.Len()
	if s <= 1{
		return;
	}
	mergeSortSequential(v.Slice(0,(s/2)).Interface(), c, !direction)
	mergeSortSequential(v.Slice((s/2),s).Interface(), c, direction)
	merge(v.Interface(), direction, c)

}

func merge(i interface{}, direction bool,  c Comparator){
	v := reflect.ValueOf(i)
	s :=  v.Len()
	m := powOfTwo(s)
	halfclean(i, m, direction, c)
	bisort(v.Slice(0,m).Interface(), direction, c)
	bisort(v.Slice(m,s).Interface(), direction, c)

}
func bisort(i interface{}, direction bool,  c Comparator){
	v := reflect.ValueOf(i)
	if v.Len() == 1{
		return
	}
	s := v.Len()
	m:= powOfTwo(s)
	halfclean(i, m, direction, c)

	bisort(v.Slice(0,m).Interface(), direction, c)
	bisort(v.Slice(m,s).Interface(), direction, c)
}

func powOfTwo(n int) int{
	i:=1
	for i < n {
		i=i<<1
	}
	lessPow:= i>>1
	return lessPow
}

func halfclean (i interface{}, m int, direction bool, c Comparator){
	v := reflect.ValueOf(i)
	s := v.Len()
	for k := 0; k < s-m ; k++{
		compSwitch(i, k, k+m, direction, c)
	}

}


func compSwitch(i interface{}, k, j int, direction bool, c Comparator){
	v := reflect.ValueOf(i)
	e := v.Index(k)
	b := v.Index(j)
	
	if direction == (c.Compare(b.Interface(), e.Interface())<0) {
		GenericSwap(b,e)
	}
}

func shellSortSequential(in interface{}, c Comparator){
	v := reflect.ValueOf(in)
	jump:= v.Len()/2
	t := reflect.TypeOf(in).Elem()
    p := reflect.New(t).Elem()
	for jump >= 1 {
		for k := 0; k < jump; k++ {
			for i := k+jump; i < v.Len(); i += jump {
                p.Set(v.Index(i))
				j := i-jump
				for ; j >= 0 && c.Compare(v.Index(j).Interface(), p.Interface()) > 0; j-=jump{
					v.Index(j+jump).Set(v.Index(j))
				}
				v.Index(j+jump).Set(p)
			}
		}
		jump=jump/2
	}
}

// An interface which defines the methods of any sequential sorting algorithm implemented
// in this package.
type SequentialSort interface{
	// Sorts the slice received as a parameter. This slice can be of any type, as long as
	// the Comparator received as a parameter works with the type of the elements of the slice.
	Sort(i interface {}, c Comparator)
}

// Implementation of the Bubblesort algorithm.
type BubbleSortSequential struct{}

// Implementation of the insertion sort algorithm.
type InsertionSortSequential struct{}

// Implementation of the sequential Bitonic mergesort algorithm, based on the paper
// 'Parallelizing the Merge Sorting Network Algorithm on a
// Multi-Core Computer Using Go and Cilk'. This implementation has been
// generalized to array sizes non power of two.
type BitonicMergesortSequential struct{}

// Implementation of the sequential Quicksort algorithm. This implementation
// uses insertion sort when the size of the array is small.
type QuickSortSequential struct{}

// Implementation of the Shellsort algorithm. This implementation uses the gap
// sequence originally proposed by Shell.
type ShellSortSequential struct{}


func (o BubbleSortSequential) Sort(i interface{}, c Comparator){
	bubbleSortSequential(i, c)
}
func (o InsertionSortSequential) Sort(i interface{}, c Comparator){
	insertionSortSequential(i, c)
}
func (o BitonicMergesortSequential) Sort(i interface{}, c Comparator){
	mergeSortSequential(i, c, true)
}

func (o QuickSortSequential) Sort(i interface{}, c Comparator){
	quickSortSequential(i, c)
}

func (o ShellSortSequential) Sort(i interface{}, c Comparator){
	shellSortSequential(i, c)
}

