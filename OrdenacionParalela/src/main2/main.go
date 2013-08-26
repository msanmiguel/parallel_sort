package main

import (
	"usc/ordenacion/integers"
	"usc/ordenacion"
	"fmt"
)

type IntSlice struct {
	a []int
}

func (is IntSlice) Length() int {
	return len(is.a)
}

func (is IntSlice) Swap(i, j int) {
	is.a[i], is.a[j] = is.a[j], is.a[i]
}

func (is IntSlice) Compare(i, j int) int {
	return is.a[i]-is.a[j]
}

func main() {
	a := ordenacion.CreateAscendingArray(100)
	//a := []int{2,34,56,12,98,23,66,91,20,45,676,221,434,17,33,234,100, 5}
	b := integers.HistogramSort { 4 }
	//a := []int{2,34,5135126,12,98,23,124166,91,20,4512335,676,221,434}
	fmt.Println(a)
	//ordenacion.ParallellQuicksort(a)

	b.Sort(a)
	//ordenacion.OrdenaPSRS(a)
	fmt.Println(a)
	fmt.Println(ordenacion.IsSorted(a))

}
