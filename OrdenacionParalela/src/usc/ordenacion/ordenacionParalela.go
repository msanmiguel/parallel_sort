package ordenacion

import(
	"runtime"
	"reflect"
	_"fmt"
	"math"
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
* Nesta implementación lánzase unha gorutina para cada chamada recursiva sempre que o número de gorrutinas
sexa menor que o número de cores
**/
func ordenaQuicksortParalelo1(a []int, NCPU int) {
	nGorutinas := 1
	c := make(chan int)
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

func ParallellQuicksort(a []int, NCPU int ) {
	M := len(a)
	control := make([]chan int, NCPU)
	datos := make([]chan int, NCPU)
	datosDevueltos := make([]chan int, NCPU)
	//fin := make(chan int)
	
	for i:= 0; i < len(control); i++ {
		control[i] = make(chan int)
		datos[i] = make(chan int, len(a))
		datosDevueltos[i] = make(chan int, len(a))
	}
	
	for i := 0; i < NCPU; i++ {
		go parallellQuicksort(i,NCPU,a[M*i/NCPU:M*(i+1)/NCPU], control, datos, datosDevueltos)
	}
	a = a[:0]
	
	for i := 0; i < NCPU; i++ {
		long:= <-datosDevueltos[i]
		for j:= 0; j<long; j++{
			a = append(a, <-datosDevueltos[i])
		}
	}
	//fmt.Println("ordenado", a)

		
}

func parallellQuicksort(id, NCPU int, a []int, control, datos, datosDevueltos []chan int,) {
	menores := make([]int, 0, len(a))
	mayores := make([]int, 0, len(a))
	b := make([]int, len(a), len(a)*NCPU)
	copy(b, a)
	for NCPU > 1 {
		var pivote int
		if id%NCPU == 0 {
			pivote = b[0]
			for i := id+1; i<id+NCPU; i++{
				control[i]<-pivote
			}
		} else {
			pivote = <- control[id]
		}
		for _, v := range b {
			if v <= pivote{
				menores = append(menores, v)
			} else{
				mayores = append(mayores, v)
			}
		}

		b = b[:0]
		w := NCPU*(id/NCPU)+NCPU/2

		if id < w {
			// enviar mayores canales[id+N/2]
			datos[id+NCPU/2] <- len(mayores)
			for _, v := range mayores {
				datos[id+NCPU/2] <- v
			}

			longitud := <- datos[id]
			for i:= 0; i < longitud; i++ {
				b = append(b, <-datos[id])
			}
			// leer menores
			copy(b[len(b):cap(b)], menores)
			b = b[:len(b)+len(menores)]
			
		} else {
			// enviar menores canales[id-N/2]
			datos[id-NCPU/2] <- len(menores)
			for _, v := range menores {
				datos[id-NCPU/2] <- v
			}
			
			longitud := <- datos[id]
			for i:= 0; i < longitud; i++ {
				b = append(b, <-datos[id])
			}
			copy(b[len(b):cap(b)], mayores)
			b = b[:len(b)+len(mayores)]
		}
		if id%NCPU == 0 {
			for i := id+1; i<id+NCPU; i++{
				control[i]<-0
			}
		} else {
			pivote = <- control[id]
		}
		menores = menores[:0]
		mayores = mayores[:0]

		NCPU = NCPU/2 
	}
	
	ordenaQuicksort1(b)
	
		 datosDevueltos[id]<-len(b)
		 for _,v := range b{
		 	datosDevueltos[id] <- v
		 	}
			
	
	//fmt.Println("Soy ", id, b)
	//fin <- 0
}





// Paralelizacion de shellsort secuencial, con tantas gorutinas como numero de cores
func ordenaShellsortParalelo1(a []int, NCPU int){
	salto:= len(a)/2
	c := make(chan int)
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


// Implementacion de bitonic sort
func OrdenaBitonicMergesortParalelo(a []int, NCPU int){
	c := make(chan int)
	tamanhoInicial := len(a)
	go ordenaMergesort(a,tamanhoInicial, true, c, NCPU)
	<-c
}
func ordenaMergesort(a []int, tamanhoInicial int, direccion bool, c chan int, NCPU int){
	if len(a) <= 1{
	 	c<-0
		return
	}
	s := len(a)
	if tamanhoInicial/s < NCPU{
		c1:= make(chan int)
		// ordena unha metade da lista en sentido ascendente e outra en sentido descendente
		go ordenaMergesort(a[0:(s/2)], tamanhoInicial, !direccion, c1, NCPU)
		go ordenaMergesort(a[(s/2):s], tamanhoInicial, direccion, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencialmente
		ordenaMergesortSecuencial(a[0:(s/2)], !direccion)
		ordenaMergesortSecuencial(a[(s/2):s], direccion)
	}
	mergep(a, direccion, tamanhoInicial, NCPU)
	c<-0
	
}

func mergep(a[]int, direccion bool, tamanhoInicial int, NCPU int){
	s:= len(a)
	m:= potenciaDe2(s)
	halfcleanp(a,tamanhoInicial, m, direccion, NCPU)
	if tamanhoInicial/s < NCPU{
		c1:=make(chan int)
		go bisortp(a[0:m], tamanhoInicial, direccion, c1, NCPU)
		go bisortp(a[m:s], tamanhoInicial, direccion, c1, NCPU)
		<-c1
		<-c1
	}else{ // secuencial
		bisort(a[0:m], direccion)
		bisort(a[m:s], direccion)
	}
	
}
func bisortp(a []int,tamanhoInicial int, direccion bool, c chan int, NCPU int){
	if len(a) == 1{
		c <- 0
		return
	}
	s := len(a)
	m:= potenciaDe2(s)
	halfcleanp(a, tamanhoInicial, m, direccion, NCPU)
	if tamanhoInicial/s < NCPU{
		c1:= make(chan int)
		go bisortp(a[0:m], tamanhoInicial, direccion, c1, NCPU)
		go bisortp(a[m:s], tamanhoInicial, direccion, c1, NCPU)
		<-c1
		<-c1
	}else{// secuencial
		bisort(a[0:m], direccion)
		bisort(a[m:s], direccion)
	}
	c<-0
}


func halfcleanp(a []int, tamanhoInicial int,  m int, direccion bool, NCPU int){
	s := len(a)
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

func calculaCapacidadBucket(a []int, t, k uint, c chan []int, c2 chan [][]int) {
	nBuckets := 1<<k
	tamBuckets := make([]int, nBuckets)

	for i:=0; i < len(a); i++ {
		pos := (a[i]>>t) & ((1<<k)-1) 
		tamBuckets[pos]++
	}
	c <- tamBuckets

	buckets := <- c2
	for i:= 0; i< len(a); i++{
		pos := (a[i]>>t) & ((1<<k)-1)// a[i] se desplaza t bits en cada iteración antes de hacer la AND con la máscara
		//fmt.Println("Metiendo", a[i], "en", pos)
		buckets[pos] = append(buckets[pos], a[i])
	}

	c <- []int{}
}

// Radix sort paralelizado
func OrdenaRadixSortParalelo(a []int, NCPU int){
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
		tamBuckets := make([]int, nBuckets)

		/*for i:=0; i< len(a); i++{
			// se calcula para cada elemento de la lista el bucket que le corresponde con los ultimos k bits
			pos = (a[i]>>t) & ((1<<k)-1) // se hace una AND con la máscara con los ultimos k bits puestos a 1 y devuelve el bucket que le corresponde al elemento
			//fmt.Println(pos)

			tamBuckets[pos]= tamBuckets[pos]+1 // se incrementa en 1 la posicion que corresponde al bucket en el que esta el elemento
		}*/
		cTam := make([]chan []int, NCPU)
		cBuckets := make([]chan [][]int, NCPU)
		tamBucketsParcial := make([][]int, 0, NCPU)
		for i:= 0; i< NCPU; i++ {
			cTam[i] = make(chan []int)
			cBuckets[i] = make(chan [][]int)
			//tamBucketsParcial[i] = make([]int, nBuckets)
			go calculaCapacidadBucket(a[i*len(a)/NCPU:(i+1)*len(a)/NCPU], t, k, cTam[i], cBuckets[i])
		}
		//cTam[NCPU-1] = make(chan []int)
		//cBuckets[NCPU-1] = make(chan [][]int)
		//go calculaCapacidadBucket(a[(NCPU-1)*(len(a)/NCPU):], t, k, cTam[NCPU-1], cBuckets[NCPU-1])

		for i:= 0; i < NCPU; i++ {
			tamParcial := <- cTam[i]
			tamBucketsParcial = append(tamBucketsParcial, tamParcial)
			for j := range tamParcial {
				tamBuckets[j] += tamParcial[j]
			}
			if i > 0 {
				previo := tamBucketsParcial[i-1]
				for j := range tamParcial {
					tamParcial[j] += previo[j]
				}
			}
		}
		//fmt.Println("TAMS", t, tamBucketsParcial)

		// Se crea un slice de slices, 1 slice por bucket, y se introducen los elementos que corresponden a cada bucket en su slice
		buckets:= make([][]int, nBuckets)
		for i,_ := range  buckets{
			buckets[i] = make([]int, tamBuckets[i])
		}
		for i:= 0; i < NCPU; i++ {
			bucketsParcial := make([][]int, nBuckets)
			if i == 0 {
				for j:=0; j < nBuckets; j++ {
					bucketsParcial[j] = buckets[j][:0]
				}
			} else {
				t_previo := tamBucketsParcial[i-1]
				for j:=0; j < nBuckets; j++ {
					bucketsParcial[j] = buckets[j][t_previo[j]:t_previo[j]]
				}
			}
			cBuckets[i] <- bucketsParcial
		}

		for i:= 0; i < NCPU; i++ {
			<- cTam[i]
		}
		//fmt.Println("BUCKETS", buckets)
		// Se pasan los elementos de los buckets al array original en el orden correcto empezando por los de primer bucket hasta el ultimo 
		//a = a[:0]
		c := make(chan int)
		for i:= 0; i< NCPU; i++{
			go copiarBucketsaLista(a, buckets, i, NCPU, c)
		}

		//fmt.Println("ITERACION", a)
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
	for j:=idProcesador*len(buckets)/NCPU; j<(idProcesador +1)*len(buckets)/NCPU; j++{
		for _, v := range buckets[j]{
			a = append(a, v)
		}
	}
	c<-0
}

// Algoritmo de ordenación por Histograma

func satisfacePivote(h, n, p int) int {
	for j := 0; j < p-1; j++ {
		if h >= (j+1)*n/p-n/(20*p) && h <= (j+1)*n/p+n/(20*p) {
			return j
		}
	}
	return -1
}

func mayorPivoteCercano(h, n, p int) int {
	for j := 0; j < p; j++ {
		if h <= (j+1)*n/p {
			return j-1
		}
	}
	return -1
}

func contarPivotesSatisfechos(histograma []int, n, p int) int {
	ultimo_satisfecho := -1
	satisfechos := 0
	for i := 0; i < len(histograma); i++ {
		//histograma[i] = histograma[i-1]+histograma[i]
		if j:= satisfacePivote(histograma[i], n, p); j > -1 && j>ultimo_satisfecho {
			satisfechos++
			ultimo_satisfecho = j
		}
	}
	return satisfechos
}

func OrdenaHistogram(a []int, NCPU int){
	n := len(a)
	var b []int
	c:= make([]chan int, NCPU)
	fin:= make(chan int)
	c2:= make([]chan []int, NCPU)
	canalesDatos := make([]chan []int, NCPU)
	for i:=0; i < NCPU; i++{
		c[i] = make(chan int)
		c2[i] = make(chan [] int)
		canalesDatos[i] = make(chan []int, NCPU)
	}
	for i:=0; i< NCPU-1; i++{
		b = a[i*(n/NCPU):i*(n/NCPU)+(n/NCPU)]
		//una go-rutina por procesador, a la que le pasamos la sublista que le corresponde
		go ordenaHist(b, fin, c, c2, i, canalesDatos, NCPU, n)
	}
	b = a[(NCPU-1)*(n/NCPU):]
	go ordenaHist(b, fin, c, c2, (NCPU-1), canalesDatos, NCPU, n)

	// Supondremos que un buen número de sondeos es pxp
	// tenemos que generar pxp sondeos dividiendo el espacio de claves (0-2³²-1) equitativamente
	k := 2*NCPU
	sondeo := make([]int, k+1)
	for i:= 0; i< k+1; i++{
		sondeo[i]= i*(math.MaxInt64/k)
	}
	var histograma []int
	for {
		//fmt.Println("SONDEO", sondeo)
		// se envía el array con los sondeos a cada procesador
		for i:= 0; i<NCPU; i++{
			c[i]<-0
			c2[i] <-sondeo
		}
		histograma = make([]int, len(sondeo)-1)
		for i:= 0; i<NCPU; i++{
			datos := <-canalesDatos[i]
			//fmt.Println("receiving", datos)
			for j := range datos {
				histograma[j] += datos[j]
			}
		}
		// fmt.Println(histograma)
		satisfechos := contarPivotesSatisfechos(histograma, n, NCPU)
		u := NCPU-1-satisfechos
		// fmt.Println("Insatisfechos", u)
		if u == 0 {
			break
		}
		sondeo2 := make([]int, 0, k)
		ultimo_indice_pivote := -1
		ultimo_pivote := 0
		for i := range histograma {
			mayor_indice_pivote := mayorPivoteCercano(histograma[i], n, NCPU)
			// fmt.Println("MAYOR PIVOTE", histograma[i], "=>", mayor_indice_pivote)
			s := mayor_indice_pivote - ultimo_indice_pivote
			pivote_actual := sondeo[i+1]
			if s > 0 {
				if j:= satisfacePivote(histograma[i], n, NCPU); j != -1 {
					if mayor_indice_pivote == j {
						s--
					}
					if s > 0 {
						numSondeos := s*k/u
						gap := (pivote_actual-ultimo_pivote)/numSondeos
						for j := 1; j < numSondeos; j++ {
							sondeo2 = append(sondeo2, ultimo_pivote+gap*j)
						}
					}
					sondeo2 = append(sondeo2, pivote_actual)
					ultimo_indice_pivote = j
				} else {
					numSondeos := s*k/u
					gap := (pivote_actual-ultimo_pivote)/numSondeos
					// fmt.Println("Aqui esta", gap, pivote_actual, ultimo_pivote)
					sondeo2 = append(sondeo2, ultimo_pivote)
					for j := 1; j < numSondeos; j++ {
						sondeo2 = append(sondeo2, ultimo_pivote+gap*j)
					}
					sondeo2 = append(sondeo2, pivote_actual)
					ultimo_indice_pivote = mayor_indice_pivote
				}
			} else if j:= satisfacePivote(histograma[i], n, NCPU); j != -1  && j > ultimo_indice_pivote {
				sondeo2 = append(sondeo2, ultimo_pivote, pivote_actual)
				ultimo_indice_pivote = j
			}
			ultimo_pivote = pivote_actual
		}
		sondeo = sondeo2
	}
	// Una vez se han calculado los mejores sondeos, se envian como pivotes definitivos
	pivotes := make([]int, 0, NCPU-1)
	ultimo_pivote := -1
	for i:= range histograma {
		if j := satisfacePivote(histograma[i], n, NCPU); j != -1 && j > ultimo_pivote {
			pivotes = append(pivotes, sondeo[i+1])
			ultimo_pivote = j
		}
	}
	// fmt.Println("P finales", pivotes)

	// se hace broadcast de la lista con los pivotes 
	for i:= 0; i<NCPU; i++{
		c[i]<-1
		c2[i]<-pivotes
	}
	for i:= 0; i<NCPU; i++{
		<-fin
	}

	for i:= 0; i<NCPU; i++{
		<-fin
	}
	a=a[:0]
	for i:= 0; i<NCPU; i++{
	 	a = append(a,<-canalesDatos[i]...)
	}
}

func ordenaHist(b []int, fin chan int, c []chan int, c2 []chan []int, nGorutina int, canalesDatos []chan []int, NCPU int, n int){
	datosLocales:= make([]int, 0, len(b)*2)
	ordenaQuicksort1(b)

	//fmt.Println("I am", nGorutina)
	for {
		fase := <- c[nGorutina]
		if fase == 0 {
			sondeo := <- c2[nGorutina]
			histograma := make([]int, len(sondeo)-1)
			for i := 0; i < len(histograma); i++ {
				histograma[i] = buscarPivote(b, sondeo[i+1])
			}
			//fmt.Println(nGorutina, histograma)
			canalesDatos[nGorutina]<- histograma
		} else {
			// fmt.Println("Here!")
			break
		}
	}

	pivotes:= <-c2[nGorutina]
	fin<-0
	var pivote int
	i:=0
	for k,v := range pivotes{
		pivote = buscarPivote(b[i:],v)
		canalesDatos[k]<-b[i:(i+pivote)]
		i = i+pivote
	}
	canalesDatos[NCPU-1]<- b[i:]
	for i:=0; i<NCPU; i++{
		aux := <-canalesDatos[nGorutina]
		datosLocales = append(datosLocales, aux...)
	}
	fin<-0
	ordenaQuicksort1(datosLocales)
	canalesDatos[nGorutina] <- datosLocales
}


//Algoritmo Parallell sorting by regular sampling (PSRS)
func OrdenaPSRS(a []int, NCPU int){
	n := len(a)
	if n <= 1 {
		return
	}
	var b []int
	c:= make(chan int)
	fin:= make(chan int)
	c2:= make(chan []int)
	canalesDatos := make([]chan []int, NCPU)
	for i:=0; i < NCPU; i++{
		canalesDatos[i] = make(chan []int, NCPU)
	}
	for i:=0; i< NCPU-1; i++{
		b = a[i*(n/NCPU):i*(n/NCPU)+(n/NCPU)]
		//una go-rutina por procesador, a la que le pasamos la sublista que le corresponde
		go psRegularSampling(b, c, fin, c2, i, canalesDatos, NCPU, n)
	}
	b = a[(NCPU-1)*(n/NCPU):]
	go psRegularSampling(b, c, fin, c2, (NCPU-1), canalesDatos, NCPU, n)

	rsampleList:= make([]int, NCPU*NCPU)
	for i := 0; i < (NCPU*NCPU); i++{
		rsampleList[i] = <-c
	}
	ordenaQuicksort1(rsampleList)
	
	for i:= 0; i<NCPU; i++{
		<-fin
	}
	// seleccion los valores de los pivotes de la regular sampling lista anterior
	var j int
	pivotes:=make([]int,NCPU-1)
	for i:= 1; i< NCPU; i++{
		j= i*NCPU+(NCPU/2)-1
		pivotes[i-1]=rsampleList[j]
	}
	// se hace broadcast de la lista con los pivotes 
	for i:= 0; i<NCPU; i++{
		c2<-pivotes
	}
	for i:= 0; i<NCPU; i++{
		<-fin
	}

	for i:= 0; i<NCPU; i++{
		<-fin
	}
	a=a[:0]
	for i:= 0; i<NCPU; i++{
	 	a = append(a,<-canalesDatos[i]...)
	}
}

func psRegularSampling(b []int, c, fin chan int, c2 chan []int, nGorutina int, canalesDatos []chan []int, NCPU int, n int){
	var j int
	datosLocales:= make([]int, 0, len(b)*2)
	ordenaQuicksort1(b)
	for i := 0; i < NCPU; i++{
		j = i*n/(NCPU*NCPU)
		c<-b[j]
	}	
	fin<-0
	pivotes:= <-c2
	fin<-0
	var pivote int
	i:=0
	for k,v := range pivotes{
		pivote = buscarPivote(b[i:],v)
		canalesDatos[k]<-b[i:(i+pivote)]
		i = i+pivote
	}
	canalesDatos[NCPU-1]<- b[i:]
	for i:=0; i<NCPU; i++{
		aux := <-canalesDatos[nGorutina]
		datosLocales = append(datosLocales, aux...)
	}
	fin<-0
	ordenaQuicksort1(datosLocales)
	canalesDatos[nGorutina] <- datosLocales
}


func buscarPivote(a []int, p int) int{
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

// func Funcion(a []int){
// 	n:= len(a)
// 	NCPU := runtime.NumCPU()
// 	var b []int
// 	c:= make(chan []int)
// 	cSink := make(chan int)
// 	fin := make(chan int)
// 	for i:=0; i< NCPU; i++{
// 		b = a[i*n/NCPU:i*n/NCPU+(n/NCPU)]
// 		fmt.Println("gorrutina %d trozo %d",i,b)
// 		go gorrutina(b, c, cSink, fin,  i)
// 	}
// 	for i:=0;i<NCPU; i++{
// 		b = <-c
// 		for j:=0; j<len(b); j++ {
// 			b[j]= 1
// 		}
// 	}
// 	for i:=0;i<NCPU; i++{
// 		cSink<-0
// 	}
// 	for i:=0;i<NCPU; i++{
// 		<-fin
// 	}

// }

// func gorrutina(b []int, c chan []int, cSink, fin chan int, i int){
// 	 ordenaQuicksort1(b)
// 	 fmt.Println("gorrutina %d trozo ordenado %d",i,b)
// 	 c<-b
// 	 //fmt.Println(b)
// 	 <-cSink
// 	 fmt.Println("despues de modificarlo en funcion gorrutina %d, trozo %d", i, b)
// 	 fin <- 0	
// }


type OrdenacionParal interface{
	Ordenar(a []int)
}

type QuicksortParal1 struct{
	NCPU  int
}
type ShellsortParal1 struct{
	NCPU int
}
type ParallellQuicksort1 struct{
	NCPU int
}
type BitonicMergesortParallell struct{
	NCPU int
}
type RadixSortParalelo struct{
	NCPU int
}
type ParallellSRegularSampling struct{
	NCPU int
}
type HistogramSort struct {
	NCPU int
}

func (o BitonicMergesortParallell) Ordenar(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	OrdenaBitonicMergesortParalelo(a, numCPU)
}


func (o QuicksortParal1) Ordenar(a []int){ 
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	ordenaQuicksortParalelo1(a, numCPU)
}


func (o ShellsortParal1) Ordenar(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	ordenaShellsortParalelo1(a, numCPU)
}

func (o ParallellQuicksort1) Ordenar(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	ParallellQuicksort(a, numCPU)
}

func (o RadixSortParalelo) Ordenar(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	OrdenaRadixSortParalelo(a, numCPU)
}

func (o ParallellSRegularSampling) Ordenar(a []int){
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	OrdenaPSRS(a, numCPU)
}


func (o HistogramSort) Ordenar(a []int) {
	var numCPU int
	if o.NCPU <=0 {
		numCPU = runtime.NumCPU()
	} else {
		numCPU = o.NCPU
	}
	runtime.GOMAXPROCS(numCPU)
	OrdenaHistogram(a, numCPU)
}
