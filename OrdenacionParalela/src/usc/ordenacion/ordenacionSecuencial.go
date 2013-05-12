package ordenacion

func OrdenaBurbuja1(a []int){
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

func OrdenaInsercion1(a []int){

	for i:=1; i<len(a);i++{
		p := a[i]
		j := i-1
		for ; j>=0 && a[j]>p; j--{
			a[j+1]=a[j]
		}
		a[j+1] = p
	}
}
func OrdenaMergesort1(a []int){
	b:= ordenaMergesort1(a)
	copy(a,b)
}

func ordenaMergesort1(a []int) []int{
	//a2:=make([]int, len(a)-len(a)/2)
	// si a tiene 0 o 1 elementos esta ordenada
	if len(a)==0 || len(a)==1{
		return a
	}
	// si a tiene al menos dos elementos se divide en dos secuencias a1 y a2 	
	a1:=a[:len(a)/2]	
	a2:=a[len(a)/2:]

	a1 = ordenaMergesort1(a1)
	a2 = ordenaMergesort1(a2)

	return mezcla(a1,a2)
	
}
 


func mezcla(a1 []int, a2 []int) []int{
	arrayMezclado := make([]int, len(a1) + len(a2))
	j:=0
	k:=0
	for k<len(a1) && j<len(a2){
		if  k<len(a1) && a1[k] <= a2[j]{
			arrayMezclado[j+k]=a1[k]
			k++
			
		}else if j<len(a2) && a1[k] > a2[j]{
			arrayMezclado[j+k]=a2[j]
			j++
		}
	}
	for k<len(a1){
		arrayMezclado[j+k]=a1[k]
		k++
	}
	for j<len(a2){
		arrayMezclado[j+k]=a2[j]
		j++
	}
	
	return arrayMezclado
}

func OrdenaQuicksort1(a []int){
	if len(a) > 1 {
		pos_pivote := recolocar(a)
		OrdenaQuicksort1(a[:pos_pivote]) // recoloco la lista de los menores
		OrdenaQuicksort1(a[(pos_pivote+1):]) // recoloco la lista de los mayores
	}
} 


// la funcion recolocar devuelve la lista recolocada y la posición en la que está el pivote
func recolocar(a []int ) int {
	var izquierdo int
	var derecho int
	var pivote int
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
			aux:= a[izquierdo]
			a[izquierdo]= a[derecho]
			a[derecho]= aux
		}
	}
	// cuando se cruzaron los indices se coloca el pivote en el lugar que le corresponde
	aux:= a[derecho]
	a[derecho]= a[0]
	a[0]= aux
	
	// se devuelve el la lista recolocada y la nueva posición del pivote
	return derecho
}
    

func OrdenaShellsort1(a []int){
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



