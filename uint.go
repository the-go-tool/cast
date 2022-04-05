package cast

import (
	"reflect"
	"strconv"
)

func init() {
	uintToStr := func(in reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(strconv.FormatUint(in.Uint(), 10)), nil
	}

	register(reflect.Uint, reflect.String, uintToStr)
	register(reflect.Uint8, reflect.String, uintToStr)
	register(reflect.Uint16, reflect.String, uintToStr)
	register(reflect.Uint32, reflect.String, uintToStr)
	register(reflect.Uint64, reflect.String, uintToStr)
}
