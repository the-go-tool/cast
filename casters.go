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
		return false, fmt.Errorf("unknown bool value: %s", in)
	})

	MustRegister(func(in bool) (out string, err error) {
		if in {
			return "true", nil
		}
		return "false", nil
	})
}

// ints -> ints -> ints
func init() {
	MustRegister(func(in int) (out int8, err error) {
		return int8(in), nil
	})
	MustRegister(func(in int) (out int16, err error) {
		return int16(in), nil
	})
	MustRegister(func(in int) (out int32, err error) {
		return int32(in), nil
	})
	MustRegister(func(in int) (out int64, err error) {
		return int64(in), nil
	})

	MustRegister(func(in int8) (out int, err error) {
		return int(in), nil
	})
	MustRegister(func(in int8) (out int16, err error) {
		return int16(in), nil
	})
	MustRegister(func(in int8) (out int32, err error) {
		return int32(in), nil
	})
	MustRegister(func(in int8) (out int64, err error) {
		return int64(in), nil
	})

	MustRegister(func(in int16) (out int8, err error) {
		return int8(in), nil
	})
	MustRegister(func(in int16) (out int, err error) {
		return int(in), nil
	})
	MustRegister(func(in int16) (out int32, err error) {
		return int32(in), nil
	})
	MustRegister(func(in int16) (out int64, err error) {
		return int64(in), nil
	})

	MustRegister(func(in int32) (out int, err error) {
		return int(in), nil
	})
	MustRegister(func(in int32) (out int8, err error) {
		return int8(in), nil
	})
	MustRegister(func(in int32) (out int16, err error) {
		return int16(in), nil
	})
	MustRegister(func(in int32) (out int64, err error) {
		return int64(in), nil
	})

	MustRegister(func(in int64) (out int, err error) {
		return int(in), nil
	})
	MustRegister(func(in int64) (out int8, err error) {
		return int8(in), nil
	})
	MustRegister(func(in int64) (out int16, err error) {
		return int16(in), nil
	})
	MustRegister(func(in int64) (out int32, err error) {
		return int32(in), nil
	})
}

// ints -> uints -> ints
func init() {
	MustRegister(func(in int64) (out uint64, err error) {
		return uint64(in), nil
	})
	MustRegisterProxy[int64, uint, uint64]()
	MustRegisterProxy[int64, uint8, uint64]()
	MustRegisterProxy[int64, uint16, uint64]()
	MustRegisterProxy[int64, uint32, uint64]()

	MustRegisterProxy[int32, uint, int64]()
	MustRegisterProxy[int32, uint8, int64]()
	MustRegisterProxy[int32, uint16, int64]()
	MustRegisterProxy[int32, uint32, int64]()
	MustRegisterProxy[int32, uint64, int64]()

	MustRegisterProxy[int16, uint, int64]()
	MustRegisterProxy[int16, uint8, int64]()
	MustRegisterProxy[int16, uint16, int64]()
	MustRegisterProxy[int16, uint32, int64]()
	MustRegisterProxy[int16, uint64, int64]()

	MustRegisterProxy[int8, uint, int64]()
	MustRegisterProxy[int8, uint8, int64]()
	MustRegisterProxy[int8, uint16, int64]()
	MustRegisterProxy[int8, uint32, int64]()
	MustRegisterProxy[int8, uint64, int64]()

	MustRegisterProxy[int, uint, int64]()
	MustRegisterProxy[int, uint8, int64]()
	MustRegisterProxy[int, uint16, int64]()
	MustRegisterProxy[int, uint32, int64]()
	MustRegisterProxy[int, uint64, int64]()

	MustRegister(func(in uint64) (out int64, err error) {
		return int64(in), nil
	})
	MustRegisterProxy[uint64, int, int64]()
	MustRegisterProxy[uint64, int8, int64]()
	MustRegisterProxy[uint64, int16, int64]()
	MustRegisterProxy[uint64, int32, int64]()

	MustRegisterProxy[uint32, int, uint64]()
	MustRegisterProxy[uint32, int8, uint64]()
	MustRegisterProxy[uint32, int16, uint64]()
	MustRegisterProxy[uint32, int32, uint64]()
	MustRegisterProxy[uint32, int64, uint64]()

	MustRegisterProxy[uint16, int, uint64]()
	MustRegisterProxy[uint16, int8, uint64]()
	MustRegisterProxy[uint16, int16, uint64]()
	MustRegisterProxy[uint16, int32, uint64]()
	MustRegisterProxy[uint16, int64, uint64]()

	MustRegisterProxy[uint8, int, uint64]()
	MustRegisterProxy[uint8, int8, uint64]()
	MustRegisterProxy[uint8, int16, uint64]()
	MustRegisterProxy[uint8, int32, uint64]()
	MustRegisterProxy[uint8, int64, uint64]()

	MustRegisterProxy[uint, int, uint64]()
	MustRegisterProxy[uint, int8, uint64]()
	MustRegisterProxy[uint, int16, uint64]()
	MustRegisterProxy[uint, int32, uint64]()
	MustRegisterProxy[uint, int64, uint64]()
}

// uints -> uints -> uints
func init() {
	MustRegister(func(in uint) (out uint8, err error) {
		return uint8(in), nil
	})
	MustRegister(func(in uint) (out uint16, err error) {
		return uint16(in), nil
	})
	MustRegister(func(in uint) (out uint32, err error) {
		return uint32(in), nil
	})
	MustRegister(func(in uint) (out uint64, err error) {
		return uint64(in), nil
	})

	MustRegister(func(in uint8) (out uint, err error) {
		return uint(in), nil
	})
	MustRegister(func(in uint8) (out uint16, err error) {
		return uint16(in), nil
	})
	MustRegister(func(in uint8) (out uint32, err error) {
		return uint32(in), nil
	})
	MustRegister(func(in uint8) (out uint64, err error) {
		return uint64(in), nil
	})

	MustRegister(func(in uint16) (out uint8, err error) {
		return uint8(in), nil
	})
	MustRegister(func(in uint16) (out uint, err error) {
		return uint(in), nil
	})
	MustRegister(func(in uint16) (out uint32, err error) {
		return uint32(in), nil
	})
	MustRegister(func(in uint16) (out uint64, err error) {
		return uint64(in), nil
	})

	MustRegister(func(in uint32) (out uint, err error) {
		return uint(in), nil
	})
	MustRegister(func(in uint32) (out uint8, err error) {
		return uint8(in), nil
	})
	MustRegister(func(in uint32) (out uint16, err error) {
		return uint16(in), nil
	})
	MustRegister(func(in uint32) (out uint64, err error) {
		return uint64(in), nil
	})

	MustRegister(func(in uint64) (out uint, err error) {
		return uint(in), nil
	})
	MustRegister(func(in uint64) (out uint8, err error) {
		return uint8(in), nil
	})
	MustRegister(func(in uint64) (out uint16, err error) {
		return uint16(in), nil
	})
	MustRegister(func(in uint64) (out uint32, err error) {
		return uint32(in), nil
	})
}

// ints -> bool -> ints
func init() {
	MustRegister(func(in int64) (out bool, err error) {
		return in != 0, nil
	})
	MustRegisterProxy[int32, bool, int64]()
	MustRegisterProxy[int16, bool, int64]()
	MustRegisterProxy[int8, bool, int64]()
	MustRegisterProxy[int, bool, int64]()

	MustRegister(func(in bool) (out int64, err error) {
		if in {
			return 1, nil
		}
		return 0, nil
	})
	MustRegisterProxy[bool, int32, int64]()
	MustRegisterProxy[bool, int16, int64]()
	MustRegisterProxy[bool, int8, int64]()
	MustRegisterProxy[bool, int, int64]()
}

// uints -> bool -> uints
func init() {
	MustRegister(func(in uint64) (out bool, err error) {
		return in != 0, nil
	})
	MustRegisterProxy[uint32, bool, uint64]()
	MustRegisterProxy[uint16, bool, uint64]()
	MustRegisterProxy[uint8, bool, uint64]()
	MustRegisterProxy[uint, bool, uint64]()

	MustRegister(func(in bool) (out uint64, err error) {
		if in {
			return 1, nil
		}
		return 0, nil
	})
	MustRegisterProxy[bool, uint32, uint64]()
	MustRegisterProxy[bool, uint16, uint64]()
	MustRegisterProxy[bool, uint8, uint64]()
	MustRegisterProxy[bool, uint, uint64]()
}
