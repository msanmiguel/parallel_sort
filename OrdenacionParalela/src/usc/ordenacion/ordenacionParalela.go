package ordenacion

import(
	_"fmt"
	"runtime"
	"reflect"
)

//Primeira implementación de Quicksort paralelo 
// Nesta implementación lánzase nunha gourutina cada chamada recursiva do algoritmo.
func ordenaQuicksortParalelo(a []int) {
	c := make(chan int)
	go ordenaQuicksortParalelo_rec(a, c)
	<-c
} 

func ordenaQuicksortParalelo_rec(a []int, c chan int){
	if len(a) > 1 {
		c2 := make(chan int)
		pos_pivote := recolocar(a) 
		go ordenaQuicksortParalelo_rec(a[:pos_pivote], c2) // recoloco la lista de los menores
		go ordenaQuicksortParalelo_rec(a[(pos_pivote+1):], c2) // recoloco la lista de los mayores
		<- c2
		<- c2
	}
	c <- 0
} 
// Segunda implementación de Quicksort paralelo
/**
* Nesta implementación lánzase unha gorutina para cada chamada recursiva sempre que o número de elementos
* da lista a ordenar sexa maior que un determinado valor, se o número de elementos da sublista é menor que o límite
* establecido chámase a Quicksort secuencial de forma recursiva
**/
func ordenaQuicksortParalelo1(a []int) {
	nGorutinas := 1
	c := make(chan int)
	NCPU := runtime.NumCPU()
	go ordenaQuicksortParalelo_rec1(a, c, nGorutinas, NCPU)
	<-c
} 


func ordenaQuicksortParalelo_rec1(a []int, c chan int, nGorutinas, NCPU int){
	if len(a) > 1 {
		if(nGorutinas< NCPU){ 
			c2 := make(chan int)
			pos_pivote := recolocar(a)
			go ordenaQuicksortParalelo_rec1(a[:pos_pivote], c2, 2*nGorutinas, NCPU) // recoloco la lista de los menores
			go ordenaQuicksortParalelo_rec1(a[(pos_pivote+1):], c2, 2*nGorutinas, NCPU) // recoloco la lista de los mayores
			<- c2
			<- c2
		} else {
			pos_pivote := recolocar(a) 
			ordenaQuicksort1(a[:pos_pivote]) // recoloco la lista de los menores
			ordenaQuicksort1(a[(pos_pivote+1):]) // recoloco la lista de los mayores
		}
	}
	c <- 0
	
} 


//Algoritmo parallellQuicksort

/**
* 
*
*/

func ParallellQuicksort(a []int) {
	N := runtime.NumCPU()
	M := len(a)
	control := make([]chan int, N)
	datos := make([]chan int, N)
	datosDevueltos := make([]chan int, N)
	//fin := make(chan int)
	
	for i:= 0; i < len(control); i++ {
		control[i] = make(chan int)
		datos[i] = make(chan int, len(a))
		datosDevueltos[i] = make(chan int, len(a))
	}
	
	for i := 0; i < N; i++ {
		go parallellQuicksort(i,N,a[M*i/N:M*(i+1)/N], control, datos, datosDevueltos)
	}
	a = a[:0]
	
	for i := 0; i < N; i++ {
		long:= <-datosDevueltos[i]
		for j:= 0; j<long; j++{
			a = append(a, <-datosDevueltos[i])
		}
	}
	//fmt.Println("ordenado", a)

		
}

func parallellQuicksort(id, N int, a []int, control, datos, datosDevueltos []chan int,) {
	menores := make([]int, 0, len(a))
	mayores := make([]int, 0, len(a))
	b := make([]int, len(a), len(a)*N)
	copy(b, a)
	for N > 1 {
		var pivote int
		if id%N == 0 {
			pivote = b[0]
			for i := id+1; i<id+N; i++{
				control[i]<-pivote
			}
		} else {
			pivote = <- control[id]
		}
//		fmt.Printf("Soy %d y mi pivote es %d (%d)\n", id, pivote, N)

		for _, v := range b {
			if v <= pivote{
				menores = append(menores, v)
			} else{
				mayores = append(mayores, v)
			}
		}

		b = b[:0]
		w := N*(id/N)+N/2

		if id < w {
			// enviar mayores canales[id+N/2]
			datos[id+N/2] <- len(mayores)
			for _, v := range mayores {
				datos[id+N/2] <- v
			}

			longitud := <- datos[id]
			for i:= 0; i < longitud; i++ {
				b = append(b, <-datos[id])
			}
			copy(b[len(b):cap(b)], menores)
			b = b[:len(b)+len(menores)]
			// leer menores
		} else {
			// enviar menores canales[id-N/2]
			datos[id-N/2] <- len(menores)
			for _, v := range menores {
				datos[id-N/2] <- v
			}
			
			longitud := <- datos[id]
			for i:= 0; i < longitud; i++ {
				b = append(b, <-datos[id])
			}
			copy(b[len(b):cap(b)], mayores)
			b = b[:len(b)+len(mayores)]
		}
		if id%N == 0 {
			for i := id+1; i<id+N; i++{
				control[i]<-0
			}
		} else {
			pivote = <- control[id]
		}
		menores = menores[:0]
		mayores = mayores[:0]

		N = N/2 
	}
	
	ordenaQuicksort1(b)
	
		 datosDevueltos[id]<-len(b)
		 for _,v := range b{
		 	datosDevueltos[id] <- v
		 	}
			
	
	//fmt.Println("Soy ", id, b)
	//fin <- 0


}

 



// Primeira implementación de Shellsort paralelo
func ordenaShellsortParalelo(a []int){
	salto:= len(a)/2
	c := make(chan int)
	for salto >= 1 {
		for k := 0; k < salto; k++ {
			go funcion(a,salto,c,k)
				
		}
		for k := 0; k < salto; k++ {
			<-c
		}
		salto=salto/2
	}
}

func funcion(a []int, salto int, c chan int, k int){
	for i:=k+salto; i<len(a);i+=salto {
		p := a[i]
		j:=i-salto
		for ; j>=0 && a[j]>p; j-=salto{
				a[j+salto]=a[j]
			}
				a[j+salto] = p	
	}
	c <-0
}



// Paralelizacion de shellsort Paralelo, con tantas gorutinas como numero de cores
func ordenaShellsortParalelo1(a []int){
	salto:= len(a)/2
	c := make(chan int)
	NCPU := runtime.NumCPU()
	for salto >= 1 {
		for k := 0; k < NCPU; k++ {
				go shellsort(a,salto,c,k, NCPU)
		}
		for k := 0; k < NCPU; k++ {
			<-c
		}
		salto=salto/2
	}
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

func ordenaMergesortParalelo(a []int){
	c := make(chan int)
	tamanhoInicial := len(a)
	go ordenaMergesort(a,tamanhoInicial, true, c)
	<-c
}
func ordenaMergesort(a []int, tamanhoInicial int, direccion bool, c chan int){

	if len(a) == 1{
	 	c<-0
		return
	}
	s := len(a)
	NCPU := runtime.NumCPU()
	if tamanhoInicial/s < NCPU{
		c1:= make(chan int)
		go ordenaMergesort(a[0:(s/2)], tamanhoInicial, !direccion, c1)
		go ordenaMergesort(a[(s/2):s], tamanhoInicial, direccion, c1)
		<-c1
		<-c1
	}else{
		ordenaMergesortSecuencial(a[0:(s/2)], !direccion)
		ordenaMergesortSecuencial(a[(s/2):s], direccion)
	}
	mergep(a, direccion, tamanhoInicial)
	c<-0
	
}

func mergep(a[]int, direccion bool, tamanhoInicial int){
	s:= len(a)
	NCPU := runtime.NumCPU()
	m:= potenciaDe2(s)
	halfcleanp(a,tamanhoInicial, m, direccion)
	if tamanhoInicial/s < NCPU{
		c1:=make(chan int)
		go bisortp(a[0:m], tamanhoInicial, direccion, c1)
		go bisortp(a[m:s], tamanhoInicial, direccion, c1)
		<-c1
		<-c1
	}else{
		bisort(a[0:m], direccion)
		bisort(a[m:s], direccion)
	}
	
}
func bisortp(a []int,tamanhoInicial int, direccion bool, c chan int){
	if len(a) == 1{
		c <- 0
		return
	}
	s := len(a)
	NCPU := runtime.NumCPU()
	m:= potenciaDe2(s)
	halfcleanp(a, tamanhoInicial, m, direccion)
	if tamanhoInicial/s < NCPU{
		c1:= make(chan int)
		go bisortp(a[0:m], tamanhoInicial, direccion, c1)
		go bisortp(a[m:s], tamanhoInicial, direccion, c1)
		<-c1
		<-c1
	}else{
		bisort(a[0:m], direccion)
		bisort(a[m:s], direccion)
	}
	c<-0
}


func halfcleanp(a []int, tamanhoInicial int,  m int, direccion bool){
	s := len(a)
	NCPU := runtime.NumCPU()
	if tamanhoInicial/s < NCPU{
		c := make(chan int)
		for j:= 0; j < s*NCPU/tamanhoInicial; j++{
			go halfCleanTrozo(a, m, direccion, j, s*NCPU/tamanhoInicial, c)
		}
		for i := 0;  i < s*NCPU/tamanhoInicial ; i++{
			<-c
		}
	}else{
		for i := 0; i < s-m ; i++{
			compSwitch(a, i, i+m, direccion)
		}
	}
	
}
func halfCleanTrozo(a []int, m int, direccion bool, j int, NCPU int, c chan int){
	s := len(a)
	for i := j; i < s-m ; i+=NCPU{
		compSwitch(a, i, i+m, direccion)
	}
	c<-0
}

// Radix sort paralelizado
func OrdenaRadixSortParalelo(a []int){
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
			pos = (a[i]>>t) & ((1<<k)-1)// a[i] se desplaza t bits en cada iteración antes de hacer la AND con la máscara
			buckets[pos] = append(buckets[pos], a[i])
		}
		// Se pasan los elementos de los buckets al array original en el orden correcto empezando por los de primer bucket hasta el ultimo 
		//a = a[:0]
		c := make(chan int)
		NCPU := runtime.NumCPU()
		for i:= 0; i< NCPU; i++{
			go copiarBucketsaLista(a, buckets, i, NCPU, c)
		}
		for i:= 0; i< NCPU; i++{
			<-c
		}
	}
}
func copiarBucketsaLista(a []int, buckets [][]int, idProcesador, NCPU int, c chan int){
	var pos int =0
	for i:=0; i <idProcesador*len(buckets)/NCPU; i++{
		pos += len(buckets[i])
	}
	a = a[pos:pos]
	for j:=idProcesador*len(buckets)/NCPU; j<(idProcesador +1)*(len(buckets)/NCPU); j++{
		for _, v := range buckets[j]{
				a = append(a, v)
		}
	}
	c<-0			
}


type OrdenacionParal interface{
	Ordenar(a []int)
}

type QuicksortParal1 struct{}
type ShellsortParal1 struct{}
type ParallellQuicksort1 struct{}
type MergesortParallel struct{}
type RadixSortParalelo struct{}

func (o MergesortParallel) Ordenar(a []int){
	ordenaMergesortParalelo(a)
}


func (o QuicksortParal1) Ordenar(a []int){
	ordenaQuicksortParalelo1(a)
}


func (o ShellsortParal1) Ordenar(a []int){
	ordenaShellsortParalelo1(a)
}

func (o ParallellQuicksort1) Ordenar(a []int){
	ParallellQuicksort(a)	
}

func (o RadixSortParalelo) Ordenar(a []int){
	OrdenaRadixSortParalelo(a)
}



