package cast

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/the-go-tool/cast/core/catchers"
)

func registerCatchers() {
	// Catcher to process same type values and pass their through.
	catchers.Add(func(in reflect.Value, to reflect.Type) (out reflect.Value, err error) {
		if in.Type() == to {
			return in, nil
		}
		return out, errors.New(catchers.NotSuitableError)
	})

	// Catcher to call String() method on structure if it is Stringer.
	catchers.Add(func(in reflect.Value, to reflect.Type) (out reflect.Value, err error) {
		if in.Kind() == reflect.Struct && to.Kind() == reflect.String {
			if stringer, ok := in.Interface().(fmt.Stringer); ok {
				return reflect.ValueOf(stringer.String()), nil
			}
		}
		return out, errors.New(catchers.NotSuitableError)
	})
}
