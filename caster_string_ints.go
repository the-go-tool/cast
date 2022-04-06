package cast

import (
	"strconv"
)

func init() {
	MustRegister(func(in string) (out int, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return out, err
		}
		return int(val), nil
	})
	MustRegister(func(in string) (out int8, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return out, err
		}
		return int8(val), nil
	})
	MustRegister(func(in string) (out int16, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return out, err
		}
		return int16(val), nil
	})
	MustRegister(func(in string) (out int32, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return out, err
		}
		return int32(val), nil
	})
	MustRegister(func(in string) (out int64, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return out, err
		}
		return int64(val), nil
	})
}
