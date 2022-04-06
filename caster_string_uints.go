package cast

import (
	"strconv"
)

func init() {
	MustRegister(func(in string) (out uint, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			return out, err
		}
		return uint(val), nil
	})
	MustRegister(func(in string) (out uint8, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			return out, err
		}
		return uint8(val), nil
	})
	MustRegister(func(in string) (out uint16, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			return out, err
		}
		return uint16(val), nil
	})
	MustRegister(func(in string) (out uint32, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			return out, err
		}
		return uint32(val), nil
	})
	MustRegister(func(in string) (out uint64, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			return out, err
		}
		return uint64(val), nil
	})
}
