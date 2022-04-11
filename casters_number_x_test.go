package cast

import "testing"

func TestCastersNumberX_IntsUints(t *testing.T) {
	for i := 0; i <= 1; i++ {
		ints := []any{int(i), int8(i), int16(i), int32(i), int64(i)}
		for _, v := range ints {
			TestAssert(v, uint(i), t)
			TestAssert(v, uint8(i), t)
			TestAssert(v, uint16(i), t)
			TestAssert(v, uint32(i), t)
			TestAssert(v, uint64(i), t)
		}
	}

	for i := 0; i <= 1; i++ {
		uints := []any{uint(i), uint8(i), uint16(i), uint32(i), uint64(i)}
		for _, v := range uints {
			TestAssert(v, int(i), t)
			TestAssert(v, int8(i), t)
			TestAssert(v, int16(i), t)
			TestAssert(v, int32(i), t)
			TestAssert(v, int64(i), t)
		}
	}
}

func TestCastersNumberX_IntsFloats(t *testing.T) {
	for i := 0; i <= 1; i++ {
		ints := []any{int(i), int8(i), int16(i), int32(i), int64(i)}
		for _, v := range ints {
			TestAssert(v, float32(i), t)
			TestAssert(v, float64(i), t)
		}
	}

	for i := 0; i <= 1; i++ {
		floats := []any{float32(i), float64(i)}
		for _, v := range floats {
			TestAssert(v, int(i), t)
			TestAssert(v, int8(i), t)
			TestAssert(v, int16(i), t)
			TestAssert(v, int32(i), t)
			TestAssert(v, int64(i), t)
		}
	}
}

func TestCastersNumberX_FloatsComplexes(t *testing.T) {
	for i := float64(-1); i <= 1; i += 0.5 {
		floats := []any{float32(i), float64(i)}
		for _, v := range floats {
			TestAssert(v, complex64(complex(i, 0)), t)
			TestAssert(v, complex128(complex(i, 0)), t)
		}
	}

	for i := float64(-1); i <= 1; i += 0.5 {
		for j := float64(-1); j <= 1; j += 0.5 {
			complexes := []any{complex64(complex(i, j)), complex128(complex(i, j))}
			for _, v := range complexes {
				TestAssert(v, float32(i), t)
				TestAssert(v, float64(i), t)
			}
		}
	}
}

func TestCastersNumberX_IntsComplexes(t *testing.T) {
	for i := -1; i <= 1; i++ {
		ints := []any{int(i), int8(i), int16(i), int32(i), int64(i)}
		for _, v := range ints {
			TestAssert(v, complex64(complex(float64(i), 0)), t)
			TestAssert(v, complex128(complex(float64(i), 0)), t)
		}
	}

	for i := float64(-1); i <= 1; i += 0.5 {
		for j := float64(-1); j <= 1; j += 0.5 {
			complexes := []any{complex64(complex(i, j)), complex128(complex(i, j))}
			for _, v := range complexes {
				TestAssert(v, int(i), t)
				TestAssert(v, int8(i), t)
				TestAssert(v, int16(i), t)
				TestAssert(v, int32(i), t)
				TestAssert(v, int64(i), t)
			}
		}
	}
}

func TestCastersNumberX_UintsComplexes(t *testing.T) {
	for i := 0; i <= 1; i++ {
		uints := []any{uint(i), uint8(i), uint16(i), uint32(i), uint64(i)}
		for _, v := range uints {
			TestAssert(v, complex64(complex(float64(i), 0)), t)
			TestAssert(v, complex128(complex(float64(i), 0)), t)
		}
	}

	for i := float64(-1); i <= 1; i += 0.5 {
		for j := float64(-1); j <= 1; j += 0.5 {
			complexes := []any{complex64(complex(i, j)), complex128(complex(i, j))}
			for _, v := range complexes {
				TestAssert(v, uint(i), t)
				TestAssert(v, uint8(i), t)
				TestAssert(v, uint16(i), t)
				TestAssert(v, uint32(i), t)
				TestAssert(v, uint64(i), t)
			}
		}
	}
}
