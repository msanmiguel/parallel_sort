// María Sanmiguel Suarez. 2013

/*
Package reflection provides functions to sort slices of any type, with the help
of the reflect package in Go library. These implementations are not particularily
efficient, since the reflect library adds a huge overhead on execution, but it is
useful to sort small sets of elements without the need of defining Len and Swap functions,
as it is required with the generic package of this library, or the Go's sort package.

The functions in this package can be used creating an instance of any sorting interface
and calling then the Sort method. Also a type which implements the Comparator interface
must be defined. This interfaces receives two values of the collection and returns
-1, 0 or 1 depending of the first value is lower, equal or greater than the second. An
example implementation can be:

	type IntSorter struct {}
	func (i IntSorter) Compare(i1, i2 interface{})int {
		v1 := i1.(int)
		v2 := i2.(int)
		return v1-v2
	}

Once this comparator is defined, it can be used to sort any slice of integers:

	s := ordenacion.CreateRandomArray(100)
	is := IntSorter {}
	qs := reflection.QuickSortSequential {}
	qs.Sort(s, is)

In case of the parallel version of any algorithm:

	s := ordenacion.CreateRandomArray(100)
	is := IntSorter {}
	qs := reflection.QuickSortParallel {}
	qs.SetNumCPUs(4)
	qs.Sort(s, is)

If the number of parallel processes is not set, or is set to a value lower or equal than
zero, then the library will automatically detect the number of CPUs of the system, as
returned by the runtime.GOMAXPROCS function.

*/
package reflection

import(
	"runtime"
	"reflect"
	_"fmt"
	_"math"
)

// Quicksort paralelo generalizado

func parallelizedQuickSort(in interface{}, comp Comparator, NCPU int) {
	nGoroutines := 1
	c := make(chan int)
	go parallelizedQuicksort_rec(in, comp, c, nGoroutines, NCPU)
	<-c
} 


func parallelizedQuicksort_rec(in interface{}, comp Comparator, c chan int, nGoroutines, NCPU int){
	v := reflect.ValueOf(in)
	if v.Len() > 1 {
		if(nGoroutines < NCPU){ 
			c2 := make(chan int)
			pivot_pos := reposition(in , comp)
			go parallelizedQuicksort_rec(v.Slice(0,pivot_pos).Interface(), comp, c2, 2*nGoroutines, NCPU) // recoloco la lista de los menores
			go parallelizedQuicksort_rec(v.Slice(pivot_pos+1, v.Len()).Interface(), comp, c2, 2*nGoroutines, NCPU) // recoloco la lista de los mayores
			<- c2
			<- c2
		} else {
			pivot_pos := reposition(in , comp)
			quickSortSequential(v.Slice(0,pivot_pos).Interface(), comp) // recoloco la lista de los menores
			quickSortSequential(v.Slice(pivot_pos+1, v.Len()).Interface(), comp) // recoloco la lista de los mayores
		}
	}
	c <- 0
	
} 


// ShellSort paralelizado generalizado

func parallelizedShellSort(in interface{}, comp Comparator, NCPU int){
	v := reflect.ValueOf(in)
	jump:= v.Len()/2
	c := make(chan int)
	for jump >= NCPU {
		for k := 0; k < NCPU; k++ {
				go shellSort(in, comp, jump, c, k, NCPU)
		}
		for k := 0; k < NCPU; k++ {
			<-c
		}
		jump = jump/2
	}
	insertionSortSequential(in, comp)
}

func shellSort(in interface{}, comp Comparator, jump int, c chan int, k int, NCPU int){
	v := reflect.ValueOf(in)
	t := reflect.TypeOf(in).Elem()
	p := reflect.New(t).Elem()
	for ; k< jump; k+=NCPU{
		for i := k+jump; i < v.Len(); i+= jump {
			p.Set(v.Index(i))
			j:=i-jump
			for ; j>=0 && comp.Compare(v.Index(j).Interface(), p.Interface()) > 0; j-=jump{
				v.Index(j+jump).Set(v.Index(j))
			}
			v.Index(j+jump).Set(p)	
		}
	}
	c <-0
}

// BitonicMergeSort paralelizado generalizado

func parallelizedBitonicMergeSort(in interface{}, comp Comparator, NCPU int){
	c := make(chan int)
	v := reflect.ValueOf(in)
	initialSize := v.Len()
	go sortMergesort(in, comp, initialSize, true, c, NCPU)
	<-c
}
func sortMergesort(in interface{}, comp Comparator, initialSize int, direction bool, c chan int, NCPU int){
	v := reflect.ValueOf(in)
	if v.Len() <= 1{
	 	c<-0
		return
	}
	s := v.Len()
	if initialSize/s < NCPU{
		c1:= make(chan int)
		// ordena unha metade da lista en sentido ascendente e outra en sentido descendente
		go sortMergesort(v.Slice(0,(s/2)).Interface(), comp, initialSize, !direction, c1, NCPU)
		go sortMergesort(v.Slice((s/2),s).Interface(), comp, initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencialmente
		mergeSortSequential(v.Slice(0,(s/2)).Interface(), comp, !direction)
		mergeSortSequential(v.Slice((s/2),s).Interface(), comp, direction)
	}
	mergep(in, comp, direction, initialSize, NCPU)
	c<-0
	
}

func mergep(in interface{}, comp Comparator, direction bool, initialSize int, NCPU int){
	v := reflect.ValueOf(in)
	s:= v.Len()
	m:= powOfTwo(s)
	halfcleanp(in, comp, initialSize, m, direction, NCPU)
	if initialSize/s < NCPU{
		c1:=make(chan int)
		go bisortp(v.Slice(0,m).Interface(), comp, initialSize, direction, c1, NCPU)
		go bisortp(v.Slice(m,s).Interface(), comp, initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencial
		bisort(v.Slice(0,m).Interface(), direction, comp)
		bisort(v.Slice(m,s).Interface(), direction, comp)
	}
	
}
func bisortp(in interface{}, comp Comparator, initialSize int, direction bool, c chan int, NCPU int){
	v := reflect.ValueOf(in)
	if v.Len() == 1{
		c <- 0
		return
	}
	s := v.Len() 
	m:= powOfTwo(s)
	halfcleanp(in, comp, initialSize, m, direction, NCPU)
	if initialSize/s < NCPU{
		c1:= make(chan int)
		go bisortp(v.Slice(0,m).Interface(), comp, initialSize, direction, c1, NCPU)
		go bisortp(v.Slice(m,s).Interface(), comp, initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{// secuencial
		bisort(v.Slice(0,m).Interface(), direction, comp)
		bisort(v.Slice(m,s).Interface(), direction, comp)
	}
	c<-0
}

func halfcleanp(in interface{}, comp Comparator, initialSize int,  m int, direction bool, NCPU int){
	v := reflect.ValueOf(in)
	s := v.Len() 
	if initialSize/s < NCPU{
		c := make(chan int)
		for j:= 0; j < s*NCPU/initialSize; j++{
			go halfCleanTrozo(in, comp, m, direction, j, s*NCPU/initialSize, c)
		}
		for i := 0;  i < s*NCPU/initialSize ; i++{
			<-c
		}
	}else{
		for i := 0; i < s-m ; i++{
			compSwitch(in, i, i+m, direction, comp)
		}
	}
	
}
func halfCleanTrozo(in interface{}, comp Comparator, m int, direction bool, j int, NCPU int, c chan int){
	v := reflect.ValueOf(in)
	s := v.Len() 
	for i := j; i < s-m ; i+=NCPU{
		compSwitch(in, i, i+m, direction, comp)
	}
	c<-0
}

// ParallelSortByRegularSampling 

//Algoritmo Parallell sorting by regular sampling (PSRS)
func parallelSRS(in interface{}, comp Comparator, NCPU int){
	v:= reflect.ValueOf(in)
	n := v.Len()
	if n <= 1 {
		return
	}
	t := reflect.TypeOf(in)
	tElem := t.Elem()
	tChan := reflect.ChanOf(reflect.BothDir, tElem)
	tChanSlice := reflect.ChanOf(reflect.BothDir, t)
	c:= reflect.MakeChan(tChan,0)
	c2:= reflect.MakeChan(tChanSlice,0)
	end:= make(chan int)
	dataChanels := make([]reflect.Value, NCPU) // slice de canales de slices genericos
	for i:=0; i < NCPU; i++{
		dataChanels[i] = reflect.MakeChan(tChanSlice,NCPU)
	}
	var b interface {}
	for i:=0; i< NCPU-1; i++{
		b = v.Slice(i*(n/NCPU), i*(n/NCPU)+(n/NCPU)).Interface()
		//una go-rutina por procesador, a la que le pasamos la sublista que le corresponde
		go psRegularSampling(b, comp, end, c, c2, dataChanels, i, NCPU, n)
	}
	b = v.Slice((NCPU-1)*(n/NCPU), v.Len()).Interface()
	go psRegularSampling(b, comp, end, c, c2, dataChanels, (NCPU-1), NCPU, n)

	rsampleList:= reflect.MakeSlice(t, NCPU*NCPU, NCPU*NCPU)
	for i := 0; i < (NCPU*NCPU); i++{
		rsample,_ := c.Recv()
		rsampleList.Index(i).Set(rsample)
	}
	quickSortSequential(rsampleList.Interface(), comp)
	
	for i:= 0; i<NCPU; i++{
		<-end
	}
	// seleccion los valores de los pivotes de la regular sampling lista anterior
	var j int
	pivots := reflect.MakeSlice(t, NCPU-1,NCPU-1)
	for i:= 1; i< NCPU; i++{
		j= i*NCPU+(NCPU/2)-1
		pivots.Index(i-1).Set(rsampleList.Index(j))
	}
	// se hace broadcast de la lista con los pivotes 
	for i:= 0; i<NCPU; i++{
		c2.Send(pivots)
	}
	for i:= 0; i<NCPU; i++{
		<-end
	}
	for i:= 0; i<NCPU; i++{
		<-end
	}
	v = v.Slice(0,0)
	for i:= 0; i<NCPU; i++{
		d,_ := dataChanels[i].Recv()
	 	v = reflect.AppendSlice(v,d)
	}
}

func psRegularSampling(b interface {}, comp Comparator, end chan int, c, c2 reflect.Value, dataChanels []reflect.Value, nGorutina int, NCPU int, n int){
	var j int
	v := reflect.ValueOf(b)
	t := reflect.TypeOf(b)
	localData:= reflect.MakeSlice(t, 0, v.Len()*2)
	quickSortSequential(b, comp)
	for i := 0; i < NCPU; i++{
		j = i*n/(NCPU*NCPU)
		c.Send(v.Index(j))
	}	
	end<-0
	pivots, _ := c2.Recv()
	end<-0
	var pivot int
	i:=0
	for j := 0; j < pivots.Len(); j++ {
		pivot = searchPivot(v.Slice(i, v.Len()).Interface(), pivots.Index(j).Interface(), comp)
		dataChanels[j].Send(v.Slice(i, (i+pivot)))
		i = i+pivot
	}
	dataChanels[NCPU-1].Send(v.Slice(i, v.Len()))
	for i:=0; i<NCPU; i++{
		aux, _ := dataChanels[nGorutina].Recv()
		localData = reflect.AppendSlice(localData, aux)
	}
	end<-0
	quickSortSequential(localData.Interface(), comp)
	dataChanels[nGorutina].Send(localData)
}


func searchPivot(a interface {}, p interface {}, comp Comparator) int{
	v := reflect.ValueOf(a)
	inf := 0
	sup := v.Len()-1
	for inf <= sup {
		mid := inf+(sup-inf)/2
		//fmt.Println("Buscando en", p, a[mid], "con", a[inf], a[sup])
		r := comp.Compare(v.Index(mid).Interface(), p)
		if r > 0 {
			sup = mid-1
		} else if r < 0 {
			inf = mid+1
		} else {
			return mid
		}
	}
	return inf
}

// An interface which defines the methods of any parallel sorting algorithm implemented
// in this package.
type ParallelSort interface{
	// Sorts the slice received as a parameter. This slice can be of any type, as long as
	// the Comparator received as a parameter works with the type of the elements of the slice.
	Sort(a interface{}, comp Comparator)
	// Sets de number of parallel processes to sort the slices. If n is <=0 the
	// number of processes created will be equal to the number of detected CPUs.
	SetNumCPUs(n int)
}

// Implementation of a parallelized Quicksort algoritm. This is is based on the
// classical Quicksort algorithm, executing the recursive calls in parallel
// gorutines, creating them until its number achieves the maximum number of
// parallel processes configured.
type QuickSortParallelized struct{
	NCPU  int
}

// Implementation of a parallelized Shellsort algorithm. In this implementation
// parallel processes are created to process different positions of the input slice
// for each gap. The gap sequence implemented is the original proposed by Shell.
type ShellSortParallelized struct{
	NCPU int
}

// Implementation of the parallel Bitonic mergesort algorithm, based on the paper
// 'Parallelizing the Merge Sorting Network Algorithm on a
// Multi-Core Computer Using Go and Cilk'. This implementation has been
// generalized to array sizes non power of two.
type BitonicMergeSortParallelized struct{
	NCPU int
}

// Implementation of the Parallel Sort by Regular Sampling algorithm.
type ParallellSortRegularSampling struct{
	NCPU int
}

func (o QuickSortParallelized) Sort(in interface{}, comp Comparator){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedQuickSort(in, comp, numCPU)
}

func (o QuickSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o ShellSortParallelized) Sort(in interface{}, comp Comparator){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedShellSort(in, comp, numCPU)
}

func (o ShellSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o BitonicMergeSortParallelized) Sort(in interface{}, comp Comparator){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedBitonicMergeSort(in, comp, numCPU)
}

func (o BitonicMergeSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o ParallellSortRegularSampling) Sort(in interface{}, comp Comparator){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelSRS(in, comp, numCPU)
}

func (o ParallellSortRegularSampling) SetNumCPUs(n int) {
	o.NCPU = n
}
