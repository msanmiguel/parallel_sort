// María Sanmiguel Suarez. 2013

package generic


func bubbleSortSequential(s SortInterface){
	n:=s.Length()
	for i:=2; i <= n; i++{
		for j:=0; j <= n-i; j++{
			if s.Compare(j, j+1)>0 {
				s.Swap(j, j+1)
			}
		}
	}

}

// InsertionSort secuencial
func insertionSortSequential(s SortInterface){
	insertionSortRange(s, 0, s.Length())
}

func insertionSortRange(s SortInterface, inf, sup int){
	//fmt.Println("Intentalo!", inf, sup)
	for i:=inf+1; i<sup; i++{
		j := i-1
		for ; j>=inf && s.Compare(j, j+1) > 0; j-- {
			//fmt.Println("HEHEHE!")
			s.Swap(j, j+1)
		}
	}
}

func quickSort(s SortInterface, i, j int) {
	length := j-i
	//fmt.Println(i, j)
	if length > 20 {
		pivot_pos := reposition(s, i, j)
		quickSort(s, i, pivot_pos)
		quickSort(s, pivot_pos+1, j)
	} else if length>1 {
		//fmt.Println("Inserta!")
		insertionSortRange(s, i, j)
		//fmt.Println("Insertado!")
	}
}

// Quicksort secuencial
func quickSortSequential(s SortInterface){
	quickSort(s, 0, s.Length())
}

// la funcion reposition devuelve la lista recolocada y la posición en la que está el pivote
func reposition(s SortInterface, i, j int) int {
	var left int
	var right int
	var pivot int

	mid := i+(j-i)/2
	if s.Compare(i, mid)>0 {
		s.Swap(i, mid)
	}
	if s.Compare(mid, j-1) > 0 {
		s.Swap(mid, j-1)
	}
	if s.Compare(i, mid) <= 0 {
		s.Swap(i, mid)
	}
	pivot= i

	left = i
	right = j-1

	// Hasta que los dos indices se crucen 
	for left < right {
		for s.Compare(right, pivot) > 0 {
			right--
		}
		for left < j && s.Compare(left, pivot) <= 0 {
			left++
		}
		// si todavia no se cruzan los indices intercambiamos 
		if left < right {
			s.Swap(left, right)
		}
	}
	// cuando se cruzaron los indices se coloca el pivote en el lugar que le corresponde
	s.Swap(right, i)
	// se devuelve el la lista recolocada y la nueva posición del pivote
	return right
}



func ordenaMergesortSequential_rec(s SortInterface, inf, sup int, direction bool) {
	n := sup-inf
	if n <= 1{
		return;
	}
	mid := inf+(sup-inf)/2
	ordenaMergesortSequential_rec(s, inf, mid, !direction)
	ordenaMergesortSequential_rec(s, mid, sup, direction)
	merge(s, inf, sup, direction)
}

// MergeSort secuencial
func bitonicMergeSortSequential(s SortInterface, direction bool){
	ordenaMergesortSequential_rec(s, 0, s.Length(), direction)
}

func merge(s SortInterface, inf, sup int, direction bool){
	n:= sup-inf
	m:= powOfTwo(n)
	halfclean(s, inf, sup, m, direction)
	bisort(s, inf, inf+m, direction)
	bisort(s, inf+m, sup, direction)

}
func bisort(s SortInterface, inf, sup int, direction bool){
	n := sup-inf
	if n == 1{
		return
	}
	m:= powOfTwo(n)
	halfclean(s, inf, sup, m, direction)
	bisort(s, inf, inf+m, direction)
	bisort(s, inf+m, sup, direction)
}

func powOfTwo(n int) int{
	i:=1
	for i < n {
		i=i<<1
	}
	lessPow:= i>>1
	return lessPow
}

func halfclean(s SortInterface, inf, sup, m int, direction bool){
	n := sup-inf
	for i := 0; i < n-m ; i++{
		compSwitch(s, inf+i, inf+i+m, direction)
	}

}

func compSwitch(s SortInterface, i, j int, direction bool){
	if direction == (s.Compare(i,j)>0) {
		s.Swap(i, j)
	}
}

func shellSortRange(s SortInterface, inf, sup int) {
	n := sup-inf
	jump:= n/2
	for jump >= 1 {
		for k := inf; k < jump; k++ {
			for i:=k+jump; i<n;i+=jump {
				j := i-jump
				for ; j>=inf && s.Compare(j, j+jump)>0; j-=jump{
					s.Swap(j, j+jump)
				}
			}
		}
		jump=jump/2
	}
}

// Shellsort secuencial
func shellSortSequential(s SortInterface){
	shellSortRange(s, 0, s.Length())
}

// Interface to be implemented on top of the collection to be sorted. This collection
// bust be accesible by an integer index.
type SortInterface interface {
	// Compare function to be used by the sorting algorithms of this package.
	// The returned value of must be:
	// <=0 if i < j,
	// = 0 if i == j,
	// >=0 if i > j.
	Compare(i, j int) int
	// Must return the length of the collection to be sorted.
	Length() int
	// Swap function which works on the underlying collection.
	Swap(i, j int)
}

// An interface which defines the methods of any sequential sorting algorithm implemented
// in this package.
type SequentialSort interface {
	Sort(s SortInterface)
}

// Implementation of the sequential Quicksort algorithm. This implementation
// uses insertion sort when the size of the array is small.
type QuickSortSequential struct {}

// Implementation of the Shellsort algorithm. This implementation uses the gap
// sequence originally proposed by Shell.
type ShellSortSequential struct{}

// Implementation of the Bubblesort algorithm.
type BubbleSortSequential struct {}

// Implementation of the insertion sort algorithm.
type InsertionSortSequential struct {}

// Implementation of the sequential Bitonic mergesort algorithm, based on the paper
// 'Parallelizing the Merge Sorting Network Algorithm on a
// Multi-Core Computer Using Go and Cilk'. This implementation has been
// generalized to array sizes non power of two.
type BitonicMergesortSequential struct {}

func (q QuickSortSequential) Sort(s SortInterface) {
	quickSortSequential(s)
}

func (q ShellSortSequential) Sort(s SortInterface) {
	shellSortSequential(s)
}

func (q BubbleSortSequential) Sort(s SortInterface) {
	bubbleSortSequential(s)
}

func (q InsertionSortSequential) Sort(s SortInterface) {
	insertionSortSequential(s)
}

func (q BitonicMergesortSequential) Sort(s SortInterface) {
	bitonicMergeSortSequential(s, true)
}
