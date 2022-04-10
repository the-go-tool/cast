package cast

import (
	"testing"
)

func setInts(val int64) []any {
	return []any{int(val), int8(val), int16(val), int32(val), int64(val)}
}
func setUints(val uint64) []any {
	return []any{uint(val), uint8(val), uint16(val), uint32(val), uint64(val)}
}
func setFloats(val float64) []any {
	return []any{float32(val), float64(val)}
}
func setComplexes(val complex128) []any {
	return []any{complex64(val), complex128(val)}
}

func TestCasters(t *testing.T) {
	t.Run("string to ints", func(t *testing.T) {
		s := "-69"
		TestAssert(s, int(-69), t)
		TestAssert(s, int8(-69), t)
		TestAssert(s, int16(-69), t)
		TestAssert(s, int32(-69), t)
		TestAssert(s, int64(-69), t)
	})
	t.Run("ints to string", func(t *testing.T) {
		for _, s := range setInts(-69) {
			TestAssert(s, "-69", t)
			TestAssert(s, "-69", t)
			TestAssert(s, "-69", t)
			TestAssert(s, "-69", t)
			TestAssert(s, "-69", t)
		}
	})

	t.Run("string to uints", func(t *testing.T) {
		s := "69"
		TestAssert(s, uint(69), t)
		TestAssert(s, uint8(69), t)
		TestAssert(s, uint16(69), t)
		TestAssert(s, uint32(69), t)
		TestAssert(s, uint64(69), t)
	})
	t.Run("uints to string", func(t *testing.T) {
		for _, s := range setUints(69) {
			TestAssert(s, "69", t)
		}
	})

	t.Run("string to floats", func(t *testing.T) {
		s := "49.5"
		TestAssert(s, float32(49.5), t)
		TestAssert(s, float64(49.5), t)
	})
	t.Run("floats to string", func(t *testing.T) {
		for _, s := range setFloats(49.5) {
			TestAssert(s, "49.5", t)
			TestAssert(s, "49.5", t)
		}
	})

	t.Run("string to complexes", func(t *testing.T) {
		s := "49.5+5i"
		TestAssert(s, complex128(complex(49.5, 5)), t)
		TestAssert(s, complex64(complex(49.5, 5)), t)
	})
	t.Run("complexes to string", func(t *testing.T) {
		for _, s := range setComplexes(complex(49.5, 5)) {
			TestAssert(s, "(49.5+5i)", t)
		}
	})

	t.Run("string to bool", func(t *testing.T) {
		TestAssert("true", true, t)
		TestAssert("false", false, t)
	})
	t.Run("bool to string", func(t *testing.T) {
		TestAssert(true, "true", t)
		TestAssert(false, "false", t)
	})

	t.Run("ints to ints", func(t *testing.T) {
		for _, s := range setInts(5) {
			TestAssert(s, int(5), t)
			TestAssert(s, int8(5), t)
			TestAssert(s, int16(5), t)
			TestAssert(s, int32(5), t)
			TestAssert(s, int64(5), t)
		}
	})

	t.Run("ints to uints", func(t *testing.T) {
		for _, s := range setInts(5) {
			TestAssert(s, uint(5), t)
			TestAssert(s, uint8(5), t)
			TestAssert(s, uint16(5), t)
			TestAssert(s, uint32(5), t)
			TestAssert(s, uint64(5), t)
		}
	})
	t.Run("uints to ints", func(t *testing.T) {
		for _, s := range setUints(5) {
			TestAssert(s, int(5), t)
			TestAssert(s, int8(5), t)
			TestAssert(s, int16(5), t)
			TestAssert(s, int32(5), t)
			TestAssert(s, int64(5), t)
		}
	})

	t.Run("uints to uints", func(t *testing.T) {
		for _, s := range setUints(5) {
			TestAssert(s, uint(5), t)
			TestAssert(s, uint8(5), t)
			TestAssert(s, uint16(5), t)
			TestAssert(s, uint32(5), t)
			TestAssert(s, uint64(5), t)
		}
	})

	t.Run("ints to bool", func(t *testing.T) {
		for _, s := range setInts(5) {
			TestAssert(s, true, t)
			TestAssert(s, true, t)
			TestAssert(s, true, t)
			TestAssert(s, true, t)
			TestAssert(s, true, t)
		}
	})
	t.Run("bool to ints", func(t *testing.T) {
		TestAssert(true, int(1), t)
		TestAssert(true, int8(1), t)
		TestAssert(true, int16(1), t)
		TestAssert(true, int32(1), t)
		TestAssert(true, int64(1), t)
	})

	t.Run("uints to bool", func(t *testing.T) {
		for _, s := range setUints(5) {
			TestAssert(s, true, t)
			TestAssert(s, true, t)
			TestAssert(s, true, t)
			TestAssert(s, true, t)
			TestAssert(s, true, t)
		}
	})
	t.Run("bool to uints", func(t *testing.T) {
		TestAssert(true, uint(1), t)
		TestAssert(true, uint8(1), t)
		TestAssert(true, uint16(1), t)
		TestAssert(true, uint32(1), t)
		TestAssert(true, uint64(1), t)
	})
}
