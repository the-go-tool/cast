package cast

import (
	"reflect"
	"strconv"
)

func init() {

	intToStr := func(in reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(strconv.FormatInt(in.Int(), 10)), nil
	}

	register(reflect.Int, reflect.String, intToStr)
	register(reflect.Int8, reflect.String, intToStr)
	register(reflect.Int16, reflect.String, intToStr)
	register(reflect.Int32, reflect.String, intToStr)
	register(reflect.Int64, reflect.String, intToStr)
}
