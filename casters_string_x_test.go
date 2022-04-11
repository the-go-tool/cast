package cast

import (
	"strconv"
	"testing"
)

func TestCastersStringX_StringInts(t *testing.T) {
	for i := int64(-1); i <= 1; i++ {
		v := strconv.FormatInt(i, 10)
		TestAssert(v, int(i), t)
		TestAssert(v, int8(i), t)
		TestAssert(v, int16(i), t)
		TestAssert(v, int32(i), t)
		TestAssert(v, int64(i), t)
	}

	for i := int64(-1); i <= 1; i++ {
		res := strconv.FormatInt(i, 10)
		ints := []any{int(i), int8(i), int16(i), int32(i), int64(i)}
		for _, v := range ints {
			TestAssert(v, res, t)
		}
	}
}

func TestCastersStringX_StringUints(t *testing.T) {
	for i := uint64(0); i <= 1; i++ {
		v := strconv.FormatUint(i, 10)
		TestAssert(v, uint(i), t)
		TestAssert(v, uint8(i), t)
		TestAssert(v, uint16(i), t)
		TestAssert(v, uint32(i), t)
		TestAssert(v, uint64(i), t)
	}

	for i := uint64(0); i <= 1; i++ {
		res := strconv.FormatUint(i, 10)
		ints := []any{uint(i), uint8(i), uint16(i), uint32(i), uint64(i)}
		for _, v := range ints {
			TestAssert(v, res, t)
		}
	}
}

func TestCastersStringX_StringFloats(t *testing.T) {
	for i := float64(-1); i <= 1; i += 0.5 {
		v := strconv.FormatFloat(i, 'f', -1, 64)
		TestAssert(v, float32(i), t)
		TestAssert(v, float64(i), t)
	}

	for i := float64(-1); i <= 1; i += 0.5 {
		res := strconv.FormatFloat(i, 'f', -1, 64)
		floats := []any{float32(i), float64(i)}
		for _, v := range floats {
			TestAssert(v, res, t)
		}
	}
}

func TestCastersStringX_StringComplexes(t *testing.T) {
	for i := float64(-1); i <= 1; i += 0.5 {
		for j := float64(-1); j <= 1; j += 0.5 {
			v := strconv.FormatComplex(complex(i, j), 'f', -1, 64)
			TestAssert(v, complex128(complex(i, j)), t)
			TestAssert(v, complex64(complex(i, j)), t)
		}
	}

	for i := float64(-1); i <= 1; i += 0.5 {
		for j := float64(-1); j <= 1; j += 0.5 {
			res := strconv.FormatComplex(complex(i, j), 'f', -1, 64)
			complexes := []any{complex128(complex(i, j)), complex64(complex(i, j))}
			for _, v := range complexes {
				TestAssert(v, res, t)
				TestAssert(v, res, t)
			}
		}
	}
}

func TestCastersStringX_StringBool(t *testing.T) {
	truthy := []string{"true", "yes", "1", "y", "t", "+"}
	for _, v := range truthy {
		TestAssert(v, true, t)
	}

	falsy := []string{"false", "no", "0", "n", "f", "-"}
	for _, v := range falsy {
		TestAssert(v, false, t)
	}

	TestAssert(true, "true", t)
	TestAssert(false, "false", t)
}
