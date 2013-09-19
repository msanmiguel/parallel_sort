// María Sanmiguel Suarez. 2013

/*
Package generic provides functions to sort collections of any type, in the same
style as the Go sort library. This package only asks for an implementation of the
interface SortInterface to provide proxy functions to the underlying collection.
Elements in this collection must be accesible by indexing with an integer value.

If a slice of integers is to be sorted with a call to a function of this package,
the SortInterface defined can have this form:

	type IntSlice []int
	func (s IntSlice) Len() {
		return len(s)
	}
	func (s IntSlice) Swap(i, j int) {
		s[i], s[j] = s[j], s[i]
	}
	func (s IntSlice) Compare(i, j int) int {
		return s[i]-s[j]
	}

Once this structure is defined, a call to an algorithm can be done by creating an
instance to any sorting interface, and calling the Sort method:

	s := ordenacion.CreateRandomArray(100)
	qs := generic.QuickSortSequential {}
	qs.Sort(s)

In case of a parallel algorithm:

	s := ordenacion.CreateRandomArray(100)
	qs := generic.QuickSortParallel {}
	qs.SetNumCPUs(4)
	qs.Sort(s)

If the number of parallel processes is not set, or is set to a value lower or equal than
zero, then the library will automatically detect the number of CPUs of the system, as
returned by the runtime.GOMAXPROCS function.

*/
package generic

import (
"runtime"
)

func parallelizedQuicksort_rec(s SortInterface, i, j, nGoroutines, NCPU int, c chan int) {
	length := j-i
	if length > 1 {
		if nGoroutines < NCPU {
			c2 := make(chan int)
			pivot_pos := reposition(s, i, j)
			go parallelizedQuicksort_rec(s, i, pivot_pos, 2*nGoroutines, NCPU, c2)
			go parallelizedQuicksort_rec(s, pivot_pos+1, j, 2*nGoroutines, NCPU, c2)
			<-c2
			<-c2
		} else {
			pivot_pos := reposition(s, i, j)
			quickSort(s, i, pivot_pos) // recoloco la lista de los menores
			quickSort(s, pivot_pos+1, j) // recoloco la lista de los mayores
		}
	}
	c<-0
}

// Quicksort 
func parallelizedQuickSort(s SortInterface, NCPU int){
	nGoroutines := 1
	c := make(chan int)
	go parallelizedQuicksort_rec(s, 0, s.Length(), nGoroutines, NCPU, c)
	<-c
}



func parallelizedShellSort(s SortInterface, NCPU int){
	jump:= s.Length()/2
	c := make(chan int)
	for jump >= NCPU {
		for k := 0; k < NCPU; k++ {
			go shellsort(s,jump,c,k, NCPU)
		}
		for k := 0; k < NCPU; k++ {
			<-c
		}
		jump=jump/2
	}
	insertionSortSequential(s)
}

func shellsort(s SortInterface, jump int, c chan int, k int, NCPU int){
	n := s.Length()
	for ; k< jump; k+=NCPU{
		for i:=k+jump; i< n;i+=jump {
			j:= i-jump
			for ; j>=0 && s.Compare(j, j+jump)>0; j-=jump{
				s.Swap(j, j+jump)
			}
		}
	}
	c <-0
}


// Implementacion de bitonic sort
func parallelizedBitonicMergeSort(s SortInterface, NCPU int){
	c := make(chan int)
	initialSize := s.Length()
	go sortMergesort(s, 0, s.Length(), initialSize, true, c, NCPU)
	<-c
}
func sortMergesort(s SortInterface, inf, sup, initialSize int, direction bool, c chan int, NCPU int){
	n := sup-inf
	if n <= 1{
	 	c<-0
		return
	}
	mid := inf+(sup-inf)/2
	if initialSize/n < NCPU{
		c1:= make(chan int)
		// ordena unha metade da lista en sentido ascendente e outra en sentido descendente
		go sortMergesort(s, inf, mid, initialSize, !direction, c1, NCPU)
		go sortMergesort(s, mid, sup, initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencialmente
		ordenaMergesortSequential_rec(s, inf, mid, !direction)
		ordenaMergesortSequential_rec(s, mid, sup, direction)
	}
	mergep(s, inf, sup, direction, initialSize, NCPU)
	c<-0
	
}

func mergep(s SortInterface, inf, sup int, direction bool, initialSize int, NCPU int){
	n:= sup-inf
	m:= powOfTwo(n)
	halfcleanp(s,inf,sup,initialSize, m, direction, NCPU)
	if initialSize/n < NCPU{
		c1:=make(chan int)
		go bisortp(s, inf, inf+m, initialSize, direction, c1, NCPU)
		go bisortp(s, inf+m, sup, initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencial
		bisort(s, inf, inf+m, direction)
		bisort(s, inf+m, sup, direction)
	}
	
}
func bisortp(s SortInterface, inf, sup, initialSize int, direction bool, c chan int, NCPU int){
	n := sup-inf
	if n == 1{
		c <- 0
		return
	}
	m:= powOfTwo(n)
	halfcleanp(s, inf, sup, initialSize, m, direction, NCPU)
	if initialSize/n < NCPU{
		c1:= make(chan int)
		go bisortp(s, inf, inf+m, initialSize, direction, c1, NCPU)
		go bisortp(s, inf+m, sup, initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{// secuencial
		bisort(s, inf, inf+m, direction)
		bisort(s, inf+m, sup, direction)
	}
	c<-0
}


func halfcleanp(s SortInterface, inf, sup, initialSize, m  int, direction bool, NCPU int){
	n := sup-inf
	if initialSize/n < NCPU{
		c := make(chan int)
		for j:= 0; j < n*NCPU/initialSize; j++{
			go halfCleanTrozo(s, inf, sup, m, direction, j, n*NCPU/initialSize, c)
		}
		for i := 0;  i < n*NCPU/initialSize ; i++{
			<-c
		}
	}else{
		for i := 0; i < n-m ; i++{
			compSwitch(s, inf+i, inf+i+m, direction)
		}
	}
	
}
func halfCleanTrozo(s SortInterface, inf, sup, m int, direction bool, j int, NCPU int, c chan int){
	n := sup-inf
	for i := j; i < n-m ; i+=NCPU{
		compSwitch(s, inf+i, inf+i+m, direction)
	}
	c<-0
}

// An interface which defines the methods of any parallel sorting algorithm implemented
// in this package.
type ParallelSort interface {
	// Sort sorts the array of slices received as a parameter.
	Sort(s SortInterface)
	// Sets de number of parallel processes to sort the slices. If n is <=0 the
	// number of processes created will be equal to the number of detected CPUs.
	SetNumCPUs(n int)
}

// Implementation of the parallel Bitonic mergesort algorithm, based on the paper
// 'Parallelizing the Merge Sorting Network Algorithm on a
// Multi-Core Computer Using Go and Cilk'. This implementation has been
// generalized to array sizes non power of two.
type BitonicMergeSortParallelized struct {
	NCPU int
}

// Implementation of a parallelized Quicksort algoritm. This is is based on the
// classical Quicksort algorithm, executing the recursive calls in parallel
// gorutines, creating them until its number achieves the maximum number of
// parallel processes configured.
type QuickSortParallelized struct {
	NCPU int
}

// Implementation of a parallelized Shellsort algorithm. In this implementation
// parallel processes are created to process different positions of the input slice
// for each gap. The gap sequence implemented is the original proposed by Shell.
type ShellSortParallelized struct{
	NCPU int
}

func (o BitonicMergeSortParallelized) Sort(s SortInterface){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedBitonicMergeSort(s, numCPU)
}

func (o BitonicMergeSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o QuickSortParallelized) Sort(s SortInterface){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedQuickSort(s, numCPU)
}

func (o QuickSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o ShellSortParallelized) Sort(s SortInterface){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedShellSort(s, numCPU)
}

func (o ShellSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}
