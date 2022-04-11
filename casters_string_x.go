package cast

import (
	"fmt"
	"strconv"
	"strings"
)

func registerStringXCast() {
	registerStringInts()
	registerStringUints()
	registerStringFloats()
	registerStringComplexes()
	registerStringBool()
}

// casters:    [string <-> ints]
// depends on: [ints <-> ints]
func registerStringInts() {
	MustRegister(func(in string) (out int64, err error) {
		val, err := strconv.ParseInt(in, 10, 64)
		return int64(val), err
	})
	MustRegisterProxy[string, int32, int64]()
	MustRegisterProxy[string, int16, int64]()
	MustRegisterProxy[string, int8, int64]()
	MustRegisterProxy[string, int, int64]()

	MustRegister(func(in int64) (out string, err error) {
		return strconv.FormatInt(int64(in), 10), nil
	})
	MustRegisterProxy[int32, string, int64]()
	MustRegisterProxy[int16, string, int64]()
	MustRegisterProxy[int8, string, int64]()
	MustRegisterProxy[int, string, int64]()
}

// casters:    [string <-> uints]
// depends on: [uints <-> uints]
func registerStringUints() {
	MustRegister(func(in string) (out uint64, err error) {
		val, err := strconv.ParseUint(in, 10, 64)
		return uint64(val), err
	})
	MustRegisterProxy[string, uint32, uint64]()
	MustRegisterProxy[string, uint16, uint64]()
	MustRegisterProxy[string, uint8, uint64]()
	MustRegisterProxy[string, uint, uint64]()

	MustRegister(func(in uint64) (out string, err error) {
		return strconv.FormatUint(uint64(in), 10), nil
	})
	MustRegisterProxy[uint32, string, uint64]()
	MustRegisterProxy[uint16, string, uint64]()
	MustRegisterProxy[uint8, string, uint64]()
	MustRegisterProxy[uint, string, uint64]()
}

// casters:    [string <-> floats]
// depends on: [floats <-> floats]
func registerStringFloats() {
	MustRegister(func(in string) (out float32, err error) {
		val, err := strconv.ParseFloat(in, 64)
		return float32(val), err
	})
	MustRegisterProxy[string, float64, float32]()

	MustRegister(func(in float32) (out string, err error) {
		return strconv.FormatFloat(float64(in), 'f', -1, 64), nil
	})
	MustRegisterProxy[float64, string, float32]()
}

// casters:    [string <-> complexes]
// depends on: [complexes <-> complexes]
func registerStringComplexes() {
	MustRegister(func(in string) (out complex128, err error) {
		val, err := strconv.ParseComplex(in, 128)
		return complex128(val), err
	})
	MustRegisterProxy[string, complex64, complex128]()

	MustRegister(func(in complex128) (out string, err error) {
		return strconv.FormatComplex(complex128(in), 'f', -1, 128), nil
	})
	MustRegisterProxy[complex64, string, complex128]()
}

// casters: [string <-> bool]
func registerStringBool() {
	MustRegister(func(in string) (out bool, err error) {
		truthy := []string{"true", "yes", "1", "y", "t", "+"}
		falsy := []string{"false", "no", "0", "n", "f", "-"}
		for _, t := range truthy {
			if strings.EqualFold(strings.TrimSpace(in), t) {
				return true, nil
			}
		}
		for _, f := range falsy {
			if strings.EqualFold(strings.TrimSpace(in), f) {
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
