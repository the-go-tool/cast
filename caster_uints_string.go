package cast

import (
	"strconv"
)

func init() {
	MustRegister(func(in uint) (out string, err error) {
		return strconv.FormatUint(uint64(in), 10), nil
	})
	MustRegister(func(in uint8) (out string, err error) {
		return strconv.FormatUint(uint64(in), 10), nil
	})
	MustRegister(func(in uint16) (out string, err error) {
		return strconv.FormatUint(uint64(in), 10), nil
	})
	MustRegister(func(in uint32) (out string, err error) {
		return strconv.FormatUint(uint64(in), 10), nil
	})
	MustRegister(func(in uint64) (out string, err error) {
		return strconv.FormatUint(uint64(in), 10), nil
	})
}
