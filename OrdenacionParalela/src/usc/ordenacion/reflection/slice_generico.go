package reflection

import (
	"reflect"
	_"fmt"
)

func GenericSwap(v1, v2 reflect.Value) {
	tipo := reflect.TypeOf(v1.Interface())
	tmp := reflect.New(tipo)

	tmp.Elem().Set(v1)
	v1.Set(v2)
	v2.Set(tmp.Elem())
}
