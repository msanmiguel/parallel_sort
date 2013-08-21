package ordenacion

import(
	"runtime"
	"reflect"
	_"fmt"
	_"math"
)

// Quicksort paralelo generalizado

func parallelizedQuicksort(in interface{}, comp Comparator, NCPU int) {
	nGorutinas := 1
	c := make(chan int)
	go parallelizedQuicksort_rec(in, comp, c, nGorutinas, NCPU)
	<-c
} 


func parallelizedQuicksort_rec(in interface{}, comp Comparator, c chan int, nGorutinas, NCPU int){
	v := reflect.ValueOf(in)
	if v.Len() > 1 {
		if(nGorutinas < NCPU){ 
			c2 := make(chan int)
			pos_pivote := genericReposition(in , comp)
			go parallelizedQuicksort_rec(v.Slice(0,pos_pivote).Interface(), comp, c2, 2*nGorutinas, NCPU) // recoloco la lista de los menores
			go parallelizedQuicksort_rec(v.Slice(pos_pivote+1, v.Len()).Interface(), comp, c2, 2*nGorutinas, NCPU) // recoloco la lista de los mayores
			<- c2
			<- c2
		} else {
			pos_pivote := genericReposition(in , comp)
			quickSortSecuential(v.Slice(0,pos_pivote).Interface(), comp) // recoloco la lista de los menores
			quickSortSecuential(v.Slice(pos_pivote+1, v.Len()).Interface(), comp) // recoloco la lista de los mayores
		}
	}
	c <- 0
	
} 


// ShellSort paralelizado generalizado

func parallelizedShellsort(in interface{}, comp Comparator, NCPU int){
	v := reflect.ValueOf(in)
	salto:= v.Len()/2
	c := make(chan int)
	for salto >= 1 {
		for k := 0; k < NCPU; k++ {
				go shellSort(in, comp, salto, c, k, NCPU)
		}
		for k := 0; k < NCPU; k++ {
			<-c
		}
		salto = salto/2
	}
}

func shellSort(in interface{}, comp Comparator, salto int, c chan int, k int, NCPU int){
	v := reflect.ValueOf(in)
	t := reflect.TypeOf(in).Elem()
	p := reflect.New(t).Elem()
	for ; k< salto; k+=NCPU{
		for i := k+salto; i < v.Len(); i+= salto {
			p.Set(v.Index(i))
			j:=i-salto
			for ; j>=0 && comp.Compare(v.Index(j).Interface(), p.Interface()) > 0; j-=salto{
				v.Index(j+salto).Set(v.Index(j))
			}
			v.Index(j+salto).Set(p)	
		}
	}
	c <-0
}

// BitonicMergeSort paralelizado generalizado

func parallelizedBitonicMergeSort(in interface{}, comp Comparator, NCPU int){
	c := make(chan int)
	v := reflect.ValueOf(in)
	tamanhoInicial := v.Len()
	go sortMergesort(in, comp, tamanhoInicial, true, c, NCPU)
	<-c
}
func sortMergesort(in interface{}, comp Comparator, tamanhoInicial int, direccion bool, c chan int, NCPU int){
	v := reflect.ValueOf(in)
	if v.Len() <= 1{
	 	c<-0
		return
	}
	s := v.Len()
	if tamanhoInicial/s < NCPU{
		c1:= make(chan int)
		// ordena unha metade da lista en sentido ascendente e outra en sentido descendente
		go sortMergesort(v.Slice(0,(s/2)).Interface(), comp, tamanhoInicial, !direccion, c1, NCPU)
		go sortMergesort(v.Slice((s/2),s).Interface(), comp, tamanhoInicial, direccion, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencialmente
		mergeSortSecuential(v.Slice(0,(s/2)).Interface(), comp, !direccion)
		mergeSortSecuential(v.Slice((s/2),s).Interface(), comp, direccion)
	}
	genricMergep(in, comp, direccion, tamanhoInicial, NCPU)
	c<-0
	
}

func genricMergep(in interface{}, comp Comparator, direccion bool, tamanhoInicial int, NCPU int){
	v := reflect.ValueOf(in)
	s:= v.Len()
	m:= potenciaDe2(s)
	geniricHalfcleanp(in, comp, tamanhoInicial, m, direccion, NCPU)
	if tamanhoInicial/s < NCPU{
		c1:=make(chan int)
		go genericBisortp(v.Slice(0,m).Interface(), comp, tamanhoInicial, direccion, c1, NCPU)
		go genericBisortp(v.Slice(m,s).Interface(), comp, tamanhoInicial, direccion, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencial
		genericBisort(v.Slice(0,m).Interface(), direccion, comp)
		genericBisort(v.Slice(m,s).Interface(), direccion, comp)
	}
	
}
func genericBisortp(in interface{}, comp Comparator, tamanhoInicial int, direccion bool, c chan int, NCPU int){
	v := reflect.ValueOf(in)
	if v.Len() == 1{
		c <- 0
		return
	}
	s := v.Len() 
	m:= potenciaDe2(s)
	geniricHalfcleanp(in, comp, tamanhoInicial, m, direccion, NCPU)
	if tamanhoInicial/s < NCPU{
		c1:= make(chan int)
		go genericBisortp(v.Slice(0,m).Interface(), comp, tamanhoInicial, direccion, c1, NCPU)
		go genericBisortp(v.Slice(m,s).Interface(), comp, tamanhoInicial, direccion, c1, NCPU)
		<-c1
		<-c1
	}else{// secuencial
		genericBisort(v.Slice(0,m).Interface(), direccion, comp)
		genericBisort(v.Slice(m,s).Interface(), direccion, comp)
	}
	c<-0
}


func geniricHalfcleanp(in interface{}, comp Comparator, tamanhoInicial int,  m int, direccion bool, NCPU int){
	v := reflect.ValueOf(in)
	s := v.Len() 
	if tamanhoInicial/s < NCPU{
		c := make(chan int)
		for j:= 0; j < s*NCPU/tamanhoInicial; j++{
			go genericHalfCleanTrozo(in, comp, m, direccion, j, s*NCPU/tamanhoInicial, c)
		}
		for i := 0;  i < s*NCPU/tamanhoInicial ; i++{
			<-c
		}
	}else{
		for i := 0; i < s-m ; i++{
			genericCompSwitch(in, i, i+m, direccion, comp)
		}
	}
	
}
func genericHalfCleanTrozo(in interface{}, comp Comparator, m int, direccion bool, j int, NCPU int, c chan int){
	v := reflect.ValueOf(in)
	s := v.Len() 
	for i := j; i < s-m ; i+=NCPU{
		genericCompSwitch(in, i, i+m, direccion, comp)
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
	fin:= make(chan int)
	canalesDatos := make([]reflect.Value, NCPU) // slice de canales de slices genericos
	for i:=0; i < NCPU; i++{
		canalesDatos[i] = reflect.MakeChan(tChanSlice,NCPU)
	}
	var b interface {}
	for i:=0; i< NCPU-1; i++{
		b = v.Slice(i*(n/NCPU), i*(n/NCPU)+(n/NCPU)).Interface()
		//una go-rutina por procesador, a la que le pasamos la sublista que le corresponde
		go genericPsRegularSampling(b, comp, fin, c, c2, canalesDatos, i, NCPU, n)
	}
	b = v.Slice((NCPU-1)*(n/NCPU), v.Len()).Interface()
	go genericPsRegularSampling(b, comp, fin, c, c2, canalesDatos, (NCPU-1), NCPU, n)

	rsampleList:= reflect.MakeSlice(t, NCPU*NCPU, NCPU*NCPU)
	for i := 0; i < (NCPU*NCPU); i++{
		rsample,_ := c.Recv()
		rsampleList.Index(i).Set(rsample)
	}
	quickSortSecuential(rsampleList.Interface(), comp)
	
	for i:= 0; i<NCPU; i++{
		<-fin
	}
	// seleccion los valores de los pivotes de la regular sampling lista anterior
	var j int
	pivotes := reflect.MakeSlice(t, NCPU-1,NCPU-1)
	for i:= 1; i< NCPU; i++{
		j= i*NCPU+(NCPU/2)-1
		pivotes.Index(i-1).Set(rsampleList.Index(j))
	}
	// se hace broadcast de la lista con los pivotes 
	for i:= 0; i<NCPU; i++{
		c2.Send(pivotes)
	}
	for i:= 0; i<NCPU; i++{
		<-fin
	}
	for i:= 0; i<NCPU; i++{
		<-fin
	}
	v = v.Slice(0,0)
	for i:= 0; i<NCPU; i++{
		d,_ := canalesDatos[i].Recv()
	 	v = reflect.AppendSlice(v,d)
	}
}

func genericPsRegularSampling(b interface {}, comp Comparator, fin chan int, c, c2 reflect.Value, canalesDatos []reflect.Value, nGorutina int, NCPU int, n int){
	var j int
	v := reflect.ValueOf(b)
	t := reflect.TypeOf(b)
	datosLocales:= reflect.MakeSlice(t, 0, v.Len()*2)
	quickSortSecuential(b, comp)
	for i := 0; i < NCPU; i++{
		j = i*n/(NCPU*NCPU)
		c.Send(v.Index(j))
	}	
	fin<-0
	pivotes, _ := c2.Recv()
	fin<-0
	var pivote int
	i:=0
	for j := 0; j < pivotes.Len(); j++ {
		pivote = genericBuscarPivote(v.Slice(i, v.Len()).Interface(), pivotes.Index(j).Interface(), comp)
		canalesDatos[j].Send(v.Slice(i, (i+pivote)))
		i = i+pivote
	}
	canalesDatos[NCPU-1].Send(v.Slice(i, v.Len()))
	for i:=0; i<NCPU; i++{
		aux, _ := canalesDatos[nGorutina].Recv()
		datosLocales = reflect.AppendSlice(datosLocales, aux)
	}
	fin<-0
	quickSortSecuential(datosLocales.Interface(), comp)
	canalesDatos[nGorutina].Send(datosLocales)
}


func genericBuscarPivote(a interface {}, p interface {}, comp Comparator) int{
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

type ParallelSort interface{
	Sort(a interface{}, comp Comparator)
}

type QuickSortParallelized struct{
	NCPU  int
}

type ShellSortParallelized struct{
	NCPU int
}

type BitonicMergesortParallelized struct{
	NCPU int
}

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
	parallelizedQuicksort(in, comp, numCPU)
}

func (o ShellSortParallelized) Sort(in interface{}, comp Comparator){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedShellsort(in, comp, numCPU)
}

func (o BitonicMergesortParallelized) Sort(in interface{}, comp Comparator){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	parallelizedBitonicMergeSort(in, comp, numCPU)
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
