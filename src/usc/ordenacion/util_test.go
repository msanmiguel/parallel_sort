// María Sanmiguel Suarez. 2013

package ordenacion

import (
	"testing"
)

func TestIsSorted(t *testing.T){

	arrayVacio :=[]int{}
	if !IsSorted(arrayVacio){
		t.Errorf("Empty array is not considered in order")
	}

}

func TestCreateRandomArray(t *testing.T) {
	array := CreateRandomArray(100)
	if len(array) != 100 {
		t.Errorf("La longitud del array creado no es correcta");
	}
}

