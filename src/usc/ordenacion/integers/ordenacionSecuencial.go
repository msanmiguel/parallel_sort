// María Sanmiguel Suarez. 2013

package integers

import(
	"sort"
	"reflect"
)


// Burbuja secuencial
func bubbleSortSequential(a []int){
	i:=0
	j:=0
	n:=len(a)
	for i=2; i<=n; i++{
		for j=0; j<=n-i; j++{
			if a[j]>a[j+1]{
				aux:= a[j]
				a[j]=a[j+1]
				a[j+1]= aux
			}
		}
	}

}

// Inserción secuencial
func insertionSortSequential(a []int){

	for i:=1; i<len(a);i++{
		p := a[i]
		j := i-1
		for ; j>=0 && a[j]>p; j--{
			a[j+1]=a[j]
		}
		a[j+1] = p
	}
}



// Quicksort secuencial
func quickSortSequential(a []int){
	if len(a) > 20 {
		pivot_pos := reposition(a)
		quickSortSequential(a[:pivot_pos]) // recoloco la lista de los menores
		quickSortSequential(a[(pivot_pos+1):]) // recoloco la lista de los mayores
	}else if len(a)>1 {
		insertionSortSequential(a)
	}
} 





// la funcion recolocar devuelve la lista recolocada y la posición en la que está el pivote
func reposition(a []int ) int {
	var left  int
	var right int
	var pivot int 
	n:= len(a)
	if a[0] >a [(n-1)/2]{
		a[0], a[(n-1)/2] =  a[(n-1)/2], a[0]
	}
	if(a[(n-1)/2] > a[n-1]){
		 a[(n-1)/2], a[n-1] = a[n-1], a[(n-1)/2]

	}
	if a[0] <= a [(n-1)/2]{
		a[0], a[(n-1)/2] = a[(n-1)/2], a[0]
	}
	pivot= a[0]

	left = 0
	right = len(a)-1

	// Hasta que los dos índices se crucen 
	for left < right {
		for a[right] > pivot{
			right--
		}
		for left < len(a) && a[left] <= pivot {
			left++
		}
		// si todavia no se cruzan los indices intercambiamos 
		if left < right {
			a[left],  a[right] =  a[right], a[left]
		}
	}
	// cuando se cruzaron los indices se coloca el pivote en el lugar que le corresponde
	a[right], a[0] = a[0],a[right]
	// se devuelve el la lista recolocada y la nueva posición del pivote
	return right
}



// MergeSort secuencial
func bitonicMergeSortSequential(a []int, direction bool){
	if len(a) <= 1{
		return;
	}
	s := len(a)
	bitonicMergeSortSequential(a[0:(s/2)], !direction)
	bitonicMergeSortSequential(a[(s/2):s], direction)
	merge(a, direction)
	

}

func merge(a[]int, direction bool){
	s:= len(a)
	m:= powOfTwo(s)
	halfclean(a, m, direction)
	bisort(a[0:m], direction)
	bisort(a[m:s], direction)

}
func bisort(a []int, direction bool){
	if len(a) == 1{
		return
	}
	s := len(a)
	m:= powOfTwo(s)
	halfclean(a, m, direction)

	bisort(a[0:m], direction)
	bisort(a[m:s], direction)
}

func powOfTwo(n int) int{
	i:=1
	for i < n {
		i=i<<1
	}
	lessPow:= i>>1
	return lessPow
}

func halfclean(a []int, m int, direction bool){
	s := len(a)
	for i := 0; i < s-m ; i++{
		compSwitch(a, i, i+m, direction)
	}

}


func compSwitch(a []int, i, j int, direction bool){
	c := a[i]
	b := a[j]
	if direction == (b < c) {
		a[i] = b
		a[j] = c
	}
}


// Shellsort secuencial
func shellSortSequential(a []int){
	jump:= len(a)/2
	for jump >= 1 {
		for k := 0; k < jump; k++ {
			for i:=k+jump; i<len(a);i+=jump {
				p := a[i]
				j := i-jump
				for ; j>=0 && a[j]>p; j-=jump{
					a[j+jump]=a[j]
				}
				a[j+jump] = p
			}
		}
		jump=jump/2
	}
}


// Radix sort secuencial
func radixSortSequential(a []int){
	if len(a) <= 1 {
		return
	}

	var k uint = 4
	kTotal  := uint(reflect.TypeOf(a[0]).Size())*8
	var t uint;
	// se desplaza el elemento k bits en cada iteracion para hacer la AND con la mascara y que devuelva el bucket que le corresponde en cada iteracion
	for t = 0; t < kTotal; t+=k {
		// en primer lugar se crea un slice que almacena e número de elementos que le corresponde a cada bucket
		nBuckets := 1<<k // esto es equivalente a 2^k, es el número de combinaciones de k bits, y da el número de buckets de k bits
		tamBuckets := make([]int, nBuckets)

		var pos int
		for i:=0; i< len(a); i++{
			// se calcula para cada elemento de la lista el bucket que le corresponde con los ultimos k bits
			pos = (a[i]>>t) & ((1<<k)-1) // se hace una AND con la máscara con los ultimos k bits puestos a 1 y devuelve el bucket que le corresponde al elemento
			//fmt.Println(pos)

			tamBuckets[pos]= tamBuckets[pos]+1 // se incrementa en 1 la posicion que corresponde al bucket en el que esta el elemento
		}
		// Se crea un slice de slices, 1 slice por bucket, y se introducen los elementos que corresponden a cada bucket en su slice
		buckets:= make([][]int, nBuckets)
		for i,_ := range  buckets{
			buckets[i] = make([]int, 0, tamBuckets[i])
		}
		for i:= 0; i< len(a); i++{
			pos = (a[i]>>t) & ((1<<k)-1)
			buckets[pos] = append(buckets[pos], a[i])
		}
		// Se pasan los elementos de los buckets al array original en el orden correcto empezando por los de primer bucket hasta el ultimo 
		a = a[:0]
		for _, v := range buckets{
			for _, w:= range v{
				a = append(a, w)
			}
		}
	}
}


// An interface which defines the methods of any sequential sorting algorithm implemented
// in this package.
type SequentialSort interface{
	Sort(a []int)
}

// Implementation of the sequential Quicksort algorithm. This implementation
// uses insertion sort when the size of the array is small.
type QuickSortSequential struct{}

// Implementation of the Shellsort algorithm. This implementation uses the gap
// sequence originally proposed by Shell.
type ShellSortSequential struct{}

// Implementation of the insertion sort algorithm.
type InsertionSortSequential struct{}

// Implementation of the Bubblesort algorithm.
type BubbleSortSequential struct{}

// Implementation of the sequential Bitonic mergesort algorithm, based on the paper
// 'Parallelizing the Merge Sorting Network Algorithm on a
// Multi-Core Computer Using Go and Cilk'. This implementation has been
// generalized to array sizes non power of two.
type BitonicMergeSortSequential struct{}

// Implementation of the Radixsort algorithm for integer keys.
type RadixSortSequential struct{}


// Proxy implementation to the sort algorithm provided by the Go library in the sort package.
type GolangSort struct{}
type  OrdenarSlice struct {
	a []int
}
func (s OrdenarSlice) Len() int{
	return len(s.a)
}
func (s OrdenarSlice) Swap(i,j int){
	s.a[i], s.a[j] = s.a[j], s.a[i]
}
func (s OrdenarSlice) Less(i, j int) bool{
	 return s.a[i]<s.a[j]		
}

func (o GolangSort) Sort(a []int){
	sort.Sort(OrdenarSlice{a})
}

func (o QuickSortSequential) Sort(a []int){
	quickSortSequential(a)
}

func (o ShellSortSequential) Sort(a []int){
	shellSortSequential(a)
}

func (o InsertionSortSequential) Sort(a []int){
	insertionSortSequential(a)
}

func (o BubbleSortSequential) Sort(a []int){
	bubbleSortSequential(a)
}

func (o BitonicMergeSortSequential) Sort(a []int){
	bitonicMergeSortSequential(a, true)
}

func (o RadixSortSequential) Sort(a []int){
	radixSortSequential(a)
}



// Una implementacion de Mergesort menos eficiente que bitonic mergesort

//func ordenaMergesort1(a []int){
//	b:= ordenaMergesort1_rec(a)
//	copy(a,b)
//}

//func ordenaMergesort1_rec(a []int) []int{
//	//a2:=make([]int, len(a)-len(a)/2)
//	// si a tiene 0 o 1 elementos esta ordenada
//	if len(a)==0 || len(a)==1{
//		return a
//	}
//	// si a tiene al menos dos elementos se divide en dos secuencias a1 y a2 	
//	a1:=a[:len(a)/2]	
//	a2:=a[len(a)/2:]
//
//	a1 = ordenaMergesort1_rec(a1)
//	a2 = ordenaMergesort1_rec(a2)
//
//	return mezcla(a1,a2)
//	
//}
// 
//
//
//func mezcla(a1 []int, a2 []int) []int{
//	arrayMezclado := make([]int, len(a1) + len(a2))
//	j:=0
//	k:=0
//	for k<len(a1) && j<len(a2){
//		if  k<len(a1) && a1[k] <= a2[j]{
//			arrayMezclado[j+k]=a1[k]
//			k++
//			
//		}else if j<len(a2) && a1[k] > a2[j]{
//			arrayMezclado[j+k]=a2[j]
//			j++
//		}
//	}
//	for k<len(a1){
//		arrayMezclado[j+k]=a1[k]
//		k++
//	}
//	for j<len(a2){
//		arrayMezclado[j+k]=a2[j]
//		j++
//	}
//	
//	return arrayMezclado
//}

