package cast

import (
	"strconv"
)

func init() {
	MustRegister(func(in int) (out string, err error) {
		return strconv.FormatInt(int64(in), 10), nil
	})
	MustRegister(func(in int8) (out string, err error) {
		return strconv.FormatInt(int64(in), 10), nil
	})
	MustRegister(func(in int16) (out string, err error) {
		return strconv.FormatInt(int64(in), 10), nil
	})
	MustRegister(func(in int32) (out string, err error) {
		return strconv.FormatInt(int64(in), 10), nil
	})
	MustRegister(func(in int64) (out string, err error) {
		return strconv.FormatInt(int64(in), 10), nil
	})
}
