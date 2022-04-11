package cast

import "testing"

func TestCastersBoolNum_BoolInts(t *testing.T) {
	for i := 0; i <= 1; i++ {
		var b bool
		if i != 0 {
			b = true
		}
		TestAssert(b, int(i), t)
		TestAssert(b, int8(i), t)
		TestAssert(b, int16(i), t)
		TestAssert(b, int32(i), t)
		TestAssert(b, int64(i), t)
	}

	for i := -1; i <= 1; i++ {
		ints := []any{int(i), int8(i), int16(i), int32(i), int64(i)}
		for _, v := range ints {
			var b bool
			if i != 0 {
				b = true
			}
			TestAssert(v, b, t)
		}
	}
}

func TestCastersBoolNum_BoolUints(t *testing.T) {
	for i := 0; i <= 1; i++ {
		var b bool
		if i != 0 {
			b = true
		}
		TestAssert(b, uint(i), t)
		TestAssert(b, uint8(i), t)
		TestAssert(b, uint16(i), t)
		TestAssert(b, uint32(i), t)
		TestAssert(b, uint64(i), t)
	}

	for i := 0; i <= 1; i++ {
		ints := []any{uint(i), uint8(i), uint16(i), uint32(i), uint64(i)}
		for _, v := range ints {
			var b bool
			if i != 0 {
				b = true
			}
			TestAssert(v, b, t)
		}
	}
}

func TestCastersBoolNum_BoolFloats(t *testing.T) {
	for i := 0; i <= 1; i++ {
		var b bool
		if i != 0 {
			b = true
		}
		TestAssert(b, float32(i), t)
		TestAssert(b, float64(i), t)
	}

	for i := float64(-1); i <= 1; i += 0.5 {
		ints := []any{float32(i), float64(i)}
		for _, v := range ints {
			var b bool
			if i != 0 {
				b = true
			}
			TestAssert(v, b, t)
		}
	}
}

func TestCastersBoolNum_BoolComplexes(t *testing.T) {
	for i := 0; i <= 1; i++ {
		var b bool
		if i != 0 {
			b = true
		}
		TestAssert(b, complex128(complex(float64(i), 0)), t)
		TestAssert(b, complex64(complex(float64(i), 0)), t)
	}

	for i := float64(-1); i <= 1; i += 0.5 {
		for j := float64(-1); j <= 1; j += 0.5 {
			complexes := []any{complex128(complex(i, j)), complex64(complex(i, j))}
			for _, v := range complexes {
				var b bool
				if i != 0 {
					b = true
				}
				TestAssert(v, b, t)
			}
		}
	}
}
