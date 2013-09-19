// María Sanmiguel Suarez. 2013

package reflection

import (
	"reflect"
	_"fmt"
)

// Function which swaps two Values, creating a temporary variable
// with the same type as the received values. Types of v1 and v2
// must be the same, or an exception will be raised.
func GenericSwap(v1, v2 reflect.Value) {
	tipo := reflect.TypeOf(v1.Interface())
	tmp := reflect.New(tipo)

	tmp.Elem().Set(v1)
	v1.Set(v2)
	v2.Set(tmp.Elem())
}
