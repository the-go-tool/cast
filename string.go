package cast

import (
	"reflect"
	"strconv"
)

func init() {
	strToInt := func(in reflect.Value) (out int64, err error) {
		var val int64
		val, err = strconv.ParseInt(in.String(), 10, 64)
		if err != nil {
			return
		}
		return val, nil
	}

	register(reflect.String, reflect.Int, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToInt(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(int(val)), nil
		}
	})
	register(reflect.String, reflect.Int8, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToInt(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(int8(val)), nil
		}
	})
	register(reflect.String, reflect.Int16, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToInt(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(int16(val)), nil
		}
	})
	register(reflect.String, reflect.Int32, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToInt(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(int32(val)), nil
		}
	})
	register(reflect.String, reflect.Int64, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToInt(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(int64(val)), nil
		}
	})

	strToUint := func(in reflect.Value) (out uint64, err error) {
		var val uint64
		val, err = strconv.ParseUint(in.String(), 10, 64)
		if err != nil {
			return
		}
		return val, nil
	}

	register(reflect.String, reflect.Uint, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToUint(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(uint(val)), nil
		}
	})
	register(reflect.String, reflect.Uint8, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToUint(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(uint8(val)), nil
		}
	})
	register(reflect.String, reflect.Uint16, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToUint(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(uint16(val)), nil
		}
	})
	register(reflect.String, reflect.Uint32, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToUint(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(uint32(val)), nil
		}
	})
	register(reflect.String, reflect.Uint64, func(in reflect.Value) (out reflect.Value, err error) {
		if val, err := strToUint(in); err != nil {
			return *new(reflect.Value), err
		} else {
			return reflect.ValueOf(uint64(val)), nil
		}
	})
}
