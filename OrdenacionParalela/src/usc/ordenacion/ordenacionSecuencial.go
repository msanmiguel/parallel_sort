package ordenacion

import(
	_ "fmt"
	"reflect"
	"sort"
)

func ordenaBurbuja1(a []int){
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

func ordenaInsercion1(a []int){

	for i:=1; i<len(a);i++{
		p := a[i]
		j := i-1
		for ; j>=0 && a[j]>p; j--{
			a[j+1]=a[j]
		}
		a[j+1] = p
	}
}
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

func ordenaQuicksort1(a []int){
	if len(a) > 20 {
		pos_pivote := recolocar(a)
		ordenaQuicksort1(a[:pos_pivote]) // recoloco la lista de los menores
		ordenaQuicksort1(a[(pos_pivote+1):]) // recoloco la lista de los mayores
	}else if len(a)>1 {
		ordenaInsercion1(a)
	}
} 


// la funcion recolocar devuelve la lista recolocada y la posición en la que está el pivote
func recolocar(a []int ) int {
	var izquierdo int
	var derecho int
	var pivote int 
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
	pivote= a[0]

	izquierdo = 0
	derecho = len(a)-1

	// Hasta que los dos indices se crucen 
	for izquierdo < derecho {
		for a[derecho] > pivote{
			derecho--
		}
		for izquierdo < len(a) && a[izquierdo] <= pivote {
			izquierdo++
		} 
		// si todavia no se cruzan los indices intercambiamos 
		if izquierdo < derecho {
			a[izquierdo],  a[derecho] =  a[derecho], a[izquierdo]
		}
	}
	// cuando se cruzaron los indices se coloca el pivote en el lugar que le corresponde
	a[derecho], a[0] = a[0],a[derecho]
	// se devuelve el la lista recolocada y la nueva posición del pivote
	return derecho
}

func ordenaMergesortSecuencial(a []int, direccion bool){
	if len(a) <= 1{
		return;
	}
	s := len(a)
	ordenaMergesortSecuencial(a[0:(s/2)], !direccion)
	ordenaMergesortSecuencial(a[(s/2):s], direccion)
	merge(a, direccion)
	

}

func merge(a[]int, direccion bool){
	s:= len(a)
	m:= potenciaDe2(s)
	halfclean(a, m, direccion)
	bisort(a[0:m], direccion)
	bisort(a[m:s], direccion)

}
func bisort(a []int, direccion bool){
	if len(a) == 1{
		return
	}
	s := len(a)
	m:= potenciaDe2(s)
	halfclean(a, m, direccion)

	bisort(a[0:m], direccion)
	bisort(a[m:s], direccion)
}

func potenciaDe2(n int) int{
	i:=1
	for i < n {
		i=i<<1
	}
	menorPotencia:= i>>1
	return menorPotencia
}

func halfclean(a []int, m int, direccion bool){
	s := len(a)
	for i := 0; i < s-m ; i++{
		compSwitch(a, i, i+m, direccion)
	}

}


func compSwitch(a []int, i, j int, direccion bool){
	c := a[i]
	b := a[j]
	if direccion == (b < c) {
		a[i] = b
		a[j] = c
	}
}
    
func ordenaRankSort(a []int){
	b:= make([]int, len(a))
	for i:=0; i<len(a); i++{
		var x int = 0
		for j:= 0; j < len(a); j++{
			if a[i]> a[j]{
				x++
			}
		}
		b[x] = a[i]
	}
	copy(a,b)
}
    

func ordenaShellsort1(a []int){
	salto:= len(a)/2
	for salto >= 1 {
		for k := 0; k < salto; k++ {
			for i:=k+salto; i<len(a);i+=salto {
				p := a[i]
				j := i-salto
				for ; j>=0 && a[j]>p; j-=salto{
					a[j+salto]=a[j]
				}
				a[j+salto] = p
			}
		}
		salto=salto/2
	}
}

func OrdenaRadixSort(a []int){
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


type OrdenacionSec interface{
	Ordenar(a []int)
}

type QuicksortSec1 struct{}
type ShellsortSec1 struct{}
type Insercion1 struct{}
type Burbuja1 struct{}
type Mergesort1 struct{}
type RankSort1 struct{}
type RadixSort1 struct{}

type GolangSort struct{}

func (o GolangSort) Ordenar(a []int){
	sort.Sort(OrdenarSlice{a})
}

func (o QuicksortSec1) Ordenar(a []int){
	ordenaQuicksort1(a)
}

func (o ShellsortSec1) Ordenar(a []int){
	ordenaShellsort1(a)
}

func (o Insercion1) Ordenar(a []int){
	ordenaInsercion1(a)
}

func (o Burbuja1) Ordenar(a []int){
	ordenaBurbuja1(a)
}

func (o Mergesort1) Ordenar(a []int){
	ordenaMergesortSecuencial(a, true)
}

func (o RankSort1) Ordenar(a []int){
	ordenaRankSort(a)
}

func (o RadixSort1) Ordenar(a []int){
	OrdenaRadixSort(a)
}

