package cast

import (
	"fmt"
	"strconv"
	"strings"
)

// string -> ints -> string
func init() {
	MustRegister(func(in string) (out int, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		return int(val), err
	})
	MustRegister(func(in string) (out int8, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		return int8(val), err
	})
	MustRegister(func(in string) (out int16, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		return int16(val), err
	})
	MustRegister(func(in string) (out int32, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		return int32(val), err
	})
	MustRegister(func(in string) (out int64, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		return int64(val), err
	})

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

// string -> uints -> string
func init() {
	MustRegister(func(in string) (out uint, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		return uint(val), err
	})
	MustRegister(func(in string) (out uint8, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		return uint8(val), err
	})
	MustRegister(func(in string) (out uint16, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		return uint16(val), err
	})
	MustRegister(func(in string) (out uint32, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		return uint32(val), err
	})
	MustRegister(func(in string) (out uint64, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		return uint64(val), err
	})

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

// string -> floats -> string
func init() {
	MustRegister(func(in string) (out float32, err error) {
		val, err := strconv.ParseFloat(in, 64)
		return float32(val), err
	})
	MustRegister(func(in string) (out float64, err error) {
		val, err := strconv.ParseFloat(in, 64)
		return float64(val), err
	})

	MustRegister(func(in float32) (out string, err error) {
		return strconv.FormatFloat(float64(in), 'f', -1, 64), nil
	})
	MustRegister(func(in float64) (out string, err error) {
		return strconv.FormatFloat(float64(in), 'f', -1, 64), nil
	})
}

// string -> complexes -> string
func init() {
	MustRegister(func(in string) (out complex128, err error) {
		val, err := strconv.ParseComplex(in, 128)
		return complex128(val), err
	})
	MustRegister(func(in string) (out complex64, err error) {
		val, err := strconv.ParseComplex(in, 128)
		return complex64(val), err
	})

	MustRegister(func(in complex128) (out string, err error) {
		return strconv.FormatComplex(complex128(in), 'f', -1, 128), nil
	})
	MustRegister(func(in complex64) (out string, err error) {
		return strconv.FormatComplex(complex128(in), 'f', -1, 128), nil
	})
}

// string -> bool -> string
func init() {
	MustRegister(func(in string) (out bool, err error) {
		truthy := []string{"1", "+", "true", "t", "yes", "y"}
		falsy := []string{"0", "-", "false", "f", "no", "n"}
		for _, t := range truthy {
			if strings.EqualFold(in, t) {
				return true, nil
			}
		}
		for _, f := range falsy {
			if strings.EqualFold(in, f) {
				return false, nil
			}
		}
		return false, fmt.Errorf("unknown value: %s", in)
	})

	MustRegister(func(in bool) (out string, err error) {
		if in {
			return "true", nil
		}
		return "false", nil
	})
}
