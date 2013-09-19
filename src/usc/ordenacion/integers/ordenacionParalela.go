// María Sanmiguel Suarez. 2013

/*
Package integers provides functions to sort slices of ints. These
implementations are more efficient than using the generic package creating
an interface for slices of integers, since overhead by function calls is
avoided.

The functions in this package can be used creating an instance of any sorting interface
and calling then the Sort method. For instance, for the Quicksort algorithm:

	s := ordenacion.CreateRandomArray(100)
	qs := integers.QuickSortSequential {}
	qs.Sort(s)

In case of a parallel algorithm:

	s := ordenacion.CreateRandomArray(100)
	qs := integers.QuickSortParallel {}
	qs.SetNumCPUs(4)
	qs.Sort(s)

If the number of parallel processes is not set, or is set to a value lower or equal than
zero, then the library will automatically detect the number of CPUs of the system, as
returned by the runtime.GOMAXPROCS function.
*/
package integers

import(
	"runtime"
	"reflect"
	_"fmt"
	"math"
)

func parallelizedQuickSort(a []int, NCPU int) {
	nGoroutines := 1
	c := make(chan int)
	go parallelizedQuicksort_rec(a, c, nGoroutines, NCPU)
	<-c
} 


func parallelizedQuicksort_rec(a []int, c chan int, nGoroutines, NCPU int){
	if len(a) > 1 {
		if(nGoroutines< NCPU){ 
			c2 := make(chan int)
			pivot_pos := reposition(a)
			go parallelizedQuicksort_rec(a[:pivot_pos], c2, 2*nGoroutines, NCPU) // recoloco la lista de los menores
			go parallelizedQuicksort_rec(a[(pivot_pos+1):], c2, 2*nGoroutines, NCPU) // recoloco la lista de los mayores
			<- c2
			<- c2
		} else {
			pivot_pos := reposition(a) 
			quickSortSequential(a[:pivot_pos]) // recoloco la lista de los menores
			quickSortSequential(a[(pivot_pos+1):]) // recoloco la lista de los mayores
		}
	}
	c <- 0
	
} 


//Algoritmo parallelQuicksort


func parallelQuickSort(a []int, NCPU int ) {
	M := len(a)
	control := make([]chan int, NCPU)
	data := make([]chan int, NCPU)
	returnedData := make([]chan int, NCPU)
	//end := make(chan int)
	
	for i:= 0; i < len(control); i++ {
		control[i] = make(chan int)
		data[i] = make(chan int, len(a))
		returnedData[i] = make(chan int, len(a))
	}
	
	for i := 0; i < NCPU; i++ {
		go parallelQuicksort(i,NCPU,a[M*i/NCPU:M*(i+1)/NCPU], control, data, returnedData)
	}
	a = a[:0]
	
	for i := 0; i < NCPU; i++ {
		long:= <-returnedData[i]
		for j:= 0; j<long; j++{
			a = append(a, <-returnedData[i])
		}
	}
	//fmt.Println("ordenado", a)

		
}

func parallelQuicksort(id, NCPU int, a []int, control, data, returnedData []chan int) {
	lowers := make([]int, 0, len(returnedData))
	highers := make([]int, 0, len(a))
	b := make([]int, len(a), len(a)*NCPU)
	copy(b, a)
	for NCPU > 1 {
		var pivot int
		if id%NCPU == 0 {
			pivot = b[0]
			for i := id+1; i<id+NCPU; i++{
				control[i]<-pivot
			}
		} else {
			pivot = <- control[id]
		}
		for _, v := range b {
			if v <= pivot{
				lowers = append(lowers, v)
			} else{
				highers = append(highers, v)
			}
		}

		b = b[:0]
		w := NCPU*(id/NCPU)+NCPU/2

		if id < w {
			// enviar mayores canales[id+N/2]
			data[id+NCPU/2] <- len(highers)
			for _, v := range highers {
				data[id+NCPU/2] <- v
			}

			length := <- data[id]
			for i:= 0; i < length; i++ {
				b = append(b, <-data[id])
			}
			// leer lowers
			copy(b[len(b):cap(b)], lowers)
			b = b[:len(b)+len(lowers)]
			
		} else {
			// enviar lowers canales[id-N/2]
			data[id-NCPU/2] <- len(lowers)
			for _, v := range lowers {
				data[id-NCPU/2] <- v
			}
			
			length := <- data[id]
			for i:= 0; i < length; i++ {
				b = append(b, <-data[id])
			}
			copy(b[len(b):cap(b)], highers)
			b = b[:len(b)+len(highers)]
		}
		if id%NCPU == 0 {
			for i := id+1; i<id+NCPU; i++{
				control[i]<-0
			}
		} else {
			pivot = <- control[id]
		}
		lowers = lowers[:0]
		highers = highers[:0]

		NCPU = NCPU/2 
	}
	
	quickSortSequential(b)
	
		 returnedData[id]<-len(b)
		 for _,v := range b{
		 	returnedData[id] <- v
		 	}
			
	
	//fmt.Println("Soy ", id, b)
	//end <- 0
}





// Paralelizacion de shellsort secuencial, con tantas gorutinas como numero de cores
func parallelizedShellSort(a []int, NCPU int){
	salto:= len(a)/2
	c := make(chan int)
	for salto >= NCPU {
		for k := 0; k < NCPU; k++ {
				go shellsort(a,salto,c,k, NCPU)
		}
		for k := 0; k < NCPU; k++ {
			<-c
		}
		salto=salto/2
	}
	insertionSortSequential(a)
}

func shellsort(a []int, salto int, c chan int, k int, NCPU int){
	for ; k< salto; k+=NCPU{
		for i:=k+salto; i<len(a);i+=salto {
			p := a[i]
			j:=i-salto
			for ; j>=0 && a[j]>p; j-=salto{
				a[j+salto]=a[j]
			}
			a[j+salto] = p	
		}
	}
	c <-0
}


// Implementacion de bitonic sort
func parallelizedBitonicMergeSort(a []int, NCPU int){
	c := make(chan int)
	initialSize := len(a)
	go mergesort(a, initialSize, true, c, NCPU)
	<-c
}
func mergesort(a []int, initialSize int, direction bool, c chan int, NCPU int){
	if len(a) <= 1{
	 	c<-0
		return
	}
	s := len(a)
	if initialSize/s < NCPU{
		c1:= make(chan int)
		// ordena unha metade da lista en sentido ascendente e outra en sentido descendente
		go mergesort(a[0:(s/2)], initialSize, !direction, c1, NCPU)
		go mergesort(a[(s/2):s], initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencialmente
		bitonicMergeSortSequential(a[0:(s/2)], !direction)
		bitonicMergeSortSequential(a[(s/2):s], direction)
	}
	mergep(a, direction, initialSize, NCPU)
	c<-0
	
}

func mergep(a[]int, direction bool, initialSize int, NCPU int){
	s:= len(a)
	m:= powOfTwo(s)
	halfcleanp(a,initialSize, m, direction, NCPU)
	if initialSize/s < NCPU{
		c1:=make(chan int)
		go bisortp(a[0:m], initialSize, direction, c1, NCPU)
		go bisortp(a[m:s], initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencial
		bisort(a[0:m], direction)
		bisort(a[m:s], direction)
	}
	
}
func bisortp(a []int,initialSize int, direction bool, c chan int, NCPU int){
	if len(a) == 1{
		c <- 0
		return
	}
	s := len(a)
	m:= powOfTwo(s)
	halfcleanp(a, initialSize, m, direction, NCPU)
	if initialSize/s < NCPU{
		c1:= make(chan int)
		go bisortp(a[0:m], initialSize, direction, c1, NCPU)
		go bisortp(a[m:s], initialSize, direction, c1, NCPU)
		<-c1
		<-c1
	}else{// secuencial
		bisort(a[0:m], direction)
		bisort(a[m:s], direction)
	}
	c<-0
}


func halfcleanp(a []int, initialSize int,  m int, direction bool, NCPU int){
	s := len(a)
	if initialSize/s < NCPU{
		c := make(chan int)
		for j:= 0; j < s*NCPU/initialSize; j++{
			go halfCleanTrozo(a, m, direction, j, s*NCPU/initialSize, c)
		}
		for i := 0;  i < s*NCPU/initialSize ; i++{
			<-c
		}
	}else{
		for i := 0; i < s-m ; i++{
			compSwitch(a, i, i+m, direction)
		}
	}
	
}
func halfCleanTrozo(a []int, m int, direction bool, j int, NCPU int, c chan int){
	s := len(a)
	for i := j; i < s-m ; i+=NCPU{
		compSwitch(a, i, i+m, direction)
	}
	c<-0
}

func getBucketCapacity(a []int, t, k uint, c chan []int, c2 chan [][]int) {
	nBuckets := 1<<k
	bucketsSize := make([]int, nBuckets)

	for i:=0; i < len(a); i++ {
		pos := (a[i]>>t) & ((1<<k)-1) 
		bucketsSize[pos]++
	}
	c <- bucketsSize

	buckets := <- c2
	for i:= 0; i< len(a); i++{
		pos := (a[i]>>t) & ((1<<k)-1)// a[i] se desplaza t bits en cada iteración antes de hacer la AND con la máscara
		//fmt.Println("Metiendo", a[i], "en", pos)
		buckets[pos] = append(buckets[pos], a[i])
	}

	c <- []int{}
}

// Radix sort paralelizado
func parallelizedRadixSort(a []int, NCPU int){
	if len(a) <= 1 {
		return
	}
	var k uint = 4
	kTotal  := uint(reflect.TypeOf(a[0]).Size())*8
	var t uint
	// se desplaza el elemento k bits en cada iteracion para hacer la AND con la mascara y que devuelva el bucket que le corresponde en cada iteracion
	for t = 0; t < kTotal; t+=k {
		// en primer lugar se crea un slice que almacena e número de elementos que le corresponde a cada bucket
		nBuckets := 1<<k // esto es equivalente a 2^k, es el número de combinaciones de k bits, y da el número de buckets de k bits
		bucketsSize := make([]int, nBuckets)

		/*for i:=0; i< len(a); i++{
			// se calcula para cada elemento de la lista el bucket que le corresponde con los ultimos k bits
			pos = (a[i]>>t) & ((1<<k)-1) // se hace una AND con la máscara con los ultimos k bits puestos a 1 y devuelve el bucket que le corresponde al elemento
			//fmt.Println(pos)

			bucketsSize[pos]= bucketsSize[pos]+1 // se incrementa en 1 la posicion que corresponde al bucket en el que esta el elemento
		}*/
		cTam := make([]chan []int, NCPU)
		cBuckets := make([]chan [][]int, NCPU)
		bucketsSizePartial := make([][]int, 0, NCPU)
		for i:= 0; i< NCPU; i++ {
			cTam[i] = make(chan []int)
			cBuckets[i] = make(chan [][]int)
			//bucketsSizePartial[i] = make([]int, nBuckets)
			go getBucketCapacity(a[i*len(a)/NCPU:(i+1)*len(a)/NCPU], t, k, cTam[i], cBuckets[i])
		}
		//cTam[NCPU-1] = make(chan []int)
		//cBuckets[NCPU-1] = make(chan [][]int)
		//go calculaCapacidadBucket(a[(NCPU-1)*(len(a)/NCPU):], t, k, cTam[NCPU-1], cBuckets[NCPU-1])

		for i:= 0; i < NCPU; i++ {
			tamPartial := <- cTam[i]
			bucketsSizePartial = append(bucketsSizePartial, tamPartial)
			for j := range tamPartial {
				bucketsSize[j] += tamPartial[j]
			}
			if i > 0 {
				previous := bucketsSizePartial[i-1]
				for j := range tamPartial {
					tamPartial[j] += previous[j]
				}
			}
		}
		//fmt.Println("TAMS", t, bucketsSizePartial)

		// Se crea un slice de slices, 1 slice por bucket, y se introducen los elementos que corresponden a cada bucket en su slice
		buckets:= make([][]int, nBuckets)
		for i,_ := range  buckets{
			buckets[i] = make([]int, bucketsSize[i])
		}
		for i:= 0; i < NCPU; i++ {
			bucketsPartial := make([][]int, nBuckets)
			if i == 0 {
				for j:=0; j < nBuckets; j++ {
					bucketsPartial[j] = buckets[j][:0]
				}
			} else {
				t_previous := bucketsSizePartial[i-1]
				for j:=0; j < nBuckets; j++ {
					bucketsPartial[j] = buckets[j][t_previous[j]:t_previous[j]]
				}
			}
			cBuckets[i] <- bucketsPartial
		}

		for i:= 0; i < NCPU; i++ {
			<- cTam[i]
		}
		//fmt.Println("BUCKETS", buckets)
		// Se pasan los elementos de los buckets al array original en el orden correcto empezando por los de primer bucket hasta el ultimo 
		//a = a[:0]
		c := make(chan int)
		for i:= 0; i< NCPU; i++{
			go copyBucketsToList(a, buckets, i, NCPU, c)
		}

		//fmt.Println("ITERACION", a)
		for i:= 0; i< NCPU; i++{
			<-c
		}
	}
}



func copyBucketsToList(a []int, buckets [][]int, idProcessor, NCPU int, c chan int){
	var pos int =0
	for i:=0; i <idProcessor*len(buckets)/NCPU; i++{
		pos += len(buckets[i])
	}
	a = a[pos:pos]
	for j:=idProcessor*len(buckets)/NCPU; j<(idProcessor +1)*len(buckets)/NCPU; j++{
		for _, v := range buckets[j]{
			a = append(a, v)
		}
	}
	c<-0
}

// Algoritmo de ordenación por Histograma

func pivotSatisfies(h, n, p int) int {
	for j := 0; j < p-1; j++ {
		if h >= (j+1)*n/p-n/(20*p) && h <= (j+1)*n/p+n/(20*p) {
			return j
		}
	}
	return -1
}

func higherPivotNearest(h, n, p int) int {
	for j := 0; j < p; j++ {
		if h <= (j+1)*n/p {
			return j-1
		}
	}
	return -1
}

func countSatisfiedPivots(histogram []int, n, p int) int {
	last_satisfied := -1
	satisfied := 0
	for i := 0; i < len(histogram); i++ {
		//histogram[i] = histogram[i-1]+histogram[i]
		if j:= pivotSatisfies(histogram[i], n, p); j > -1 && j>last_satisfied {
			satisfied++
			last_satisfied = j
		}
	}
	return satisfied
}

func histogramSort(a []int, NCPU int){
	n := len(a)
	var b []int
	c:= make([]chan int, NCPU)
	end:= make(chan int)
	c2:= make([]chan []int, NCPU)
	dataChannel := make([]chan []int, NCPU)
	for i:=0; i < NCPU; i++{
		c[i] = make(chan int)
		c2[i] = make(chan [] int)
		dataChannel[i] = make(chan []int, NCPU)
	}
	for i:=0; i< NCPU-1; i++{
		b = a[i*(n/NCPU):i*(n/NCPU)+(n/NCPU)]
		//una go-rutina por procesador, a la que le pasamos la sublista que le corresponde
		go sortHist(b, end, c, c2, i, dataChannel, NCPU, n)
	}
	b = a[(NCPU-1)*(n/NCPU):]
	go sortHist(b, end, c, c2, (NCPU-1), dataChannel, NCPU, n)

	// Supondremos que un buen número de sondeos es pxp
	// tenemos que generar pxp sondeos dividiendo el espacio de claves (0-2³²-1) equitativamente
	k := 2*NCPU
	sounding := make([]int, k+1)
	for i:= 0; i< k+1; i++{
		sounding[i]= i*(math.MaxInt64/k)
	}
	var histogram []int
	for {
		//fmt.Println("SONDEO", sounding)
		// se envía el array con los sondeos a cada procesador
		for i:= 0; i<NCPU; i++{
			// Este 0 indica la fase en la que estamos (enviar sondeos)
			c[i]<-0
			c2[i] <-sounding
		}
		histogram = make([]int, len(sounding)-1)
		for i:= 0; i<NCPU; i++{
			data := <-dataChannel[i]
			//fmt.Println("receiving", datos)
			for j := range data {
				histogram[j] += data[j]
			}
		}
		//fmt.Println(histogram)
		satisfied := countSatisfiedPivots(histogram, n, NCPU)
		u := NCPU-1-satisfied
		if u == 0 {
			break
		}
		sounding2 := make([]int, 0, k)
		last_pivot_index := -1
		last_pivot := 0
		for i := range histogram {
			// Primero se calculan los posibles pivotes que quedan entre cada par
			// de valores del histograma.
			// Por un lado se calcula cuantos pivotes quedan a la izquierda del valor
			// de histograma actual.
			larges_pivot_index := higherPivotNearest(histogram[i], n, NCPU)
			//fmt.Println("MAYOR PIVOTE", histogram[i], "=>", larges_pivot_index)
			// Se le restan los que quedaron a la izquierda del valor de histograma
			// anterior. Esto nos da una estimacion de cuantos quedan encerrados.
			s := larges_pivot_index - last_pivot_index
			// Almacenamos el pivote asociado al valor de histograma que estamos procesando
			actual_pivot := sounding[i+1]
			if s > 0 { // Si hay pivotes encerrados en el intervalo
				// Si el propio valor del histograma satisface un pivote
				if j:= pivotSatisfies(histogram[i], n, NCPU); j != -1 {
					// Como estamos contando los valores a la izquierda del valor del histograma,
					// puede pasar que este ya sea válido (el valor del histograma
					// está muy cerca a la derecha del pivote teórico).
					// En ese caso el número de pivotes encerrados hay que decrementarlo en uno.
					if larges_pivot_index == j {
						s--
					}
					// Comprobamos de nuevo que el número de pivotes encerrados es mayor que 0
					// En ese caso, subdividimos el intervalo
					if s > 0 {
						soundingNum := s*k/u
						// No debemos subdividir mas el intervalo que el numero de elementos que haya
						// en el
						if soundingNum > (actual_pivot-last_pivot) {
							soundingNum = actual_pivot-last_pivot
						}
						gap := (actual_pivot-last_pivot)/soundingNum
						for j := 1; j < soundingNum; j++ {
							sounding2 = append(sounding2, last_pivot+gap*j)
						}
					}
					sounding2 = append(sounding2, actual_pivot)
					last_pivot_index = j
				} else {
					// Aquí sabemos que el valor de s es fiable, y nos da exactamente
					// el numero de pivotes encerrados.
					soundingNum := s*k/u
					if soundingNum > (actual_pivot-last_pivot) {
						soundingNum = actual_pivot-last_pivot
					}
					gap := (actual_pivot-last_pivot)/soundingNum
					// fmt.Println("Aqui esta", gap, actual_pivot, last_pivot)
					sounding2 = append(sounding2, last_pivot)
					for j := 1; j < soundingNum; j++ {
						sounding2 = append(sounding2, last_pivot+gap*j)
					}
					sounding2 = append(sounding2, actual_pivot)
					last_pivot_index = larges_pivot_index
				}
			} else if j:= pivotSatisfies(histogram[i], n, NCPU); j != -1  && j > last_pivot_index {
				// En este caso no hay pivotes en el intervalo.
				// Solo nos queda saber si este valor es en si mismo un pivote valido
				// y lo añadimos a la lista de sondeo.
				sounding2 = append(sounding2, last_pivot, actual_pivot)
				last_pivot_index = j
			}
			last_pivot = actual_pivot
		}
		sounding = sounding2
	}
	// Una vez se han calculado los mejores sondeos, se envian como pivotes definitivos
	pivots := make([]int, 0, NCPU-1)
	last_pivot := -1
	for i:= range histogram {
		if j := pivotSatisfies(histogram[i], n, NCPU); j != -1 && j > last_pivot {
			pivots = append(pivots, sounding[i+1])
			last_pivot = j
		}
	}
	//fmt.Println("P finales", pivots)

	// se hace broadcast de la lista con los pivotes 
	for i:= 0; i<NCPU; i++{
		c[i]<-1
		c2[i]<-pivots
	}
	for i:= 0; i<NCPU; i++{
		<-end
	}

	for i:= 0; i<NCPU; i++{
		<-end
	}
	a=a[:0]
	for i:= 0; i<NCPU; i++{
	 	a = append(a,<-dataChannel[i]...)
	}
}

func sortHist(b []int, end chan int, c []chan int, c2 []chan []int, nGoroutines int, dataChannels []chan []int, NCPU int, n int){
	localData:= make([]int, 0, len(b)*2)
	quickSortSequential(b)

	
	for {
		//fmt.Println("I am", nGoroutines)
		phase := <- c[nGoroutines]
		if phase == 0 {
			sounding := <- c2[nGoroutines]
			histogram := make([]int, len(sounding)-1)
			for i := 0; i < len(histogram); i++ {
				histogram[i] = searchPivot(b, sounding[i+1])
			}
			//fmt.Println(nGoroutines, histogram)
			dataChannels[nGoroutines]<- histogram
		} else {
			// fmt.Println("Here!")
			break
		}
	}

	pivots:= <-c2[nGoroutines]
	end<-0
	var pivot int
	i:=0
	for k,v := range pivots{
		pivot = searchPivot(b[i:],v)
		dataChannels[k]<-b[i:(i+pivot)]
		i = i+pivot
	}
	dataChannels[NCPU-1]<- b[i:]
	for i:=0; i<NCPU; i++{
		aux := <-dataChannels[nGoroutines]
		localData = append(localData, aux...)
	}
	end<-0
	quickSortSequential(localData)
	dataChannels[nGoroutines] <- localData
}


//Algoritmo Parallell sorting by regular sampling (PSRS)
func parallelSRS(a []int, NCPU int){
	n := len(a)
	if n <= 1 {
		return
	}
	var b []int
	c:= make(chan int)
	end:= make(chan int)
	c2:= make(chan []int)
	dataChannels := make([]chan []int, NCPU)
	for i:=0; i < NCPU; i++{
		dataChannels[i] = make(chan []int, NCPU)
	}
	for i:=0; i< NCPU-1; i++{
		b = a[i*(n/NCPU):i*(n/NCPU)+(n/NCPU)]
		//una go-rutina por procesador, a la que le pasamos la sublista que le corresponde
		go psRegularSampling(b, c, end, c2, i, dataChannels, NCPU, n)
	}
	b = a[(NCPU-1)*(n/NCPU):]
	go psRegularSampling(b, c, end, c2, (NCPU-1), dataChannels, NCPU, n)

	rsampleList:= make([]int, NCPU*NCPU)
	for i := 0; i < (NCPU*NCPU); i++{
		rsampleList[i] = <-c
	}
	quickSortSequential(rsampleList)
	
	for i:= 0; i<NCPU; i++{
		<-end
	}
	// seleccion los valores de los pivotes de la regular sampling lista anterior
	var j int
	pivots:=make([]int,NCPU-1)
	for i:= 1; i< NCPU; i++{
		j= i*NCPU+(NCPU/2)-1
		pivots[i-1]=rsampleList[j]
	}
	// se hace broadcast de la lista con los pivotes 
	for i:= 0; i<NCPU; i++{
		c2<-pivots
	}
	for i:= 0; i<NCPU; i++{
		<-end
	}

	for i:= 0; i<NCPU; i++{
		<-end
	}
	a=a[:0]
	for i:= 0; i<NCPU; i++{
	 	a = append(a,<-dataChannels[i]...)
	}
}

func psRegularSampling(b []int, c, end chan int, c2 chan []int, nGoroutines int, dataChannels []chan []int, NCPU int, n int){
	var j int
	localData:= make([]int, 0, len(b)*2)
	quickSortSequential(b)
	for i := 0; i < NCPU; i++{
		j = i*n/(NCPU*NCPU)
		c<-b[j]
	}	
	end<-0
	pivots:= <-c2
	end<-0
	var pivot int
	i:=0
	for k,v := range pivots{
		pivot = searchPivot(b[i:],v)
		dataChannels[k]<-b[i:(i+pivot)]
		i = i+pivot
	}
	dataChannels[NCPU-1]<- b[i:]
	for i:=0; i<NCPU; i++{
		aux := <-dataChannels[nGoroutines]
		localData = append(localData, aux...)
	}
	end<-0
	quickSortSequential(localData)
	dataChannels[nGoroutines] <- localData
}


func searchPivot(a []int, p int) int{
	inf := 0
	sup := len(a)-1
	for inf <= sup {
		mid := inf+(sup-inf)/2
		//fmt.Println("Buscando en", p, a[mid], "con", a[inf], a[sup])
		if a[mid] > p {
			sup = mid-1
		} else if a[mid] < p {
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
	// Sorts the slice of integers received as a parameter.
	Sort(a  []int)
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

// Implementation of the Parallel Quicksort algorithm as described in 
// 'Parallel Programming in C with MPI and OpenMP'. This algorithm only
// works for a power of two number of parallel processes.
type ParallelQuickSort struct{
	NCPU int
}

// Implementation of the parallel Bitonic mergesort algorithm, based on the paper
// 'Parallelizing the Merge Sorting Network Algorithm on a
// Multi-Core Computer Using Go and Cilk'. This implementation has been
// generalized to array sizes non power of two.
type BitonicMergeSortParallelized struct{
	NCPU int
}

// Implementation of a parallelized version of the Radix sort algorithm for
// integer keys.
type RadixSortParallelized struct{
	NCPU int
}

// Implementation of the Parallel Sort by Regular Sampling algorithm.
type ParallellSortRegularSampling struct{
	NCPU int
}

// Implementation of the Histogram sort algorithm.
type HistogramSort struct {
	NCPU int
}

func (o BitonicMergeSortParallelized) Sort(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedBitonicMergeSort(a, numCPU)
}

func (o BitonicMergeSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}


func (o QuickSortParallelized) Sort(a []int){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedQuickSort(a, numCPU)
}

func (o QuickSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}


func (o ShellSortParallelized) Sort(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedShellSort(a, numCPU)
}

func (o ShellSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o ParallelQuickSort) Sort(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelQuickSort(a, numCPU)
}

func (o ParallelQuickSort) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o RadixSortParallelized) Sort(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedRadixSort(a, numCPU)
}

func (o RadixSortParallelized) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o ParallellSortRegularSampling) Sort(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelSRS(a, numCPU)
}

func (o ParallellSortRegularSampling) SetNumCPUs(n int) {
	o.NCPU = n
}

func (o HistogramSort) Sort(a []int) {
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	histogramSort(a, numCPU)
}

func (o HistogramSort) SetNumCPUs(n int) {
	o.NCPU = n
}
