package cast

import (
	"testing"
)

func TestCasters(t *testing.T) {
	t.Run("string to ints", func(t *testing.T) {
		TestAssert("-69", int(-69), t)
		TestAssert("-69", int8(-69), t)
		TestAssert("-69", int16(-69), t)
		TestAssert("-69", int32(-69), t)
		TestAssert("-69", int64(-69), t)
	})
	t.Run("ints to string", func(t *testing.T) {
		TestAssert(int(-69), "-69", t)
		TestAssert(int8(-69), "-69", t)
		TestAssert(int16(-69), "-69", t)
		TestAssert(int32(-69), "-69", t)
		TestAssert(int64(-69), "-69", t)
	})

	t.Run("string to uints", func(t *testing.T) {
		TestAssert("69", uint(69), t)
		TestAssert("69", uint8(69), t)
		TestAssert("69", uint16(69), t)
		TestAssert("69", uint32(69), t)
		TestAssert("69", uint64(69), t)
	})
	t.Run("uints to string", func(t *testing.T) {
		TestAssert(uint(69), "69", t)
		TestAssert(uint8(69), "69", t)
		TestAssert(uint16(69), "69", t)
		TestAssert(uint32(69), "69", t)
		TestAssert(uint64(69), "69", t)
	})

	t.Run("string to floats", func(t *testing.T) {
		TestAssert("49.5", float32(49.5), t)
		TestAssert("49.5", float64(49.5), t)
	})
	t.Run("floats to string", func(t *testing.T) {
		TestAssert(float32(49.5), "49.5", t)
		TestAssert(float64(49.5), "49.5", t)
	})

	t.Run("string to complexes", func(t *testing.T) {
		TestAssert("49.5+5i", complex128(complex(49.5, 5)), t)
		TestAssert("49.5+5i", complex64(complex(49.5, 5)), t)
	})
	t.Run("complexes to string", func(t *testing.T) {
		TestAssert(complex128(complex(49.5, 5)), "(49.5+5i)", t)
		TestAssert(complex64(complex(49.5, 5)), "(49.5+5i)", t)
	})

	t.Run("string to bool", func(t *testing.T) {
		TestAssert("true", true, t)
		TestAssert("false", false, t)
	})
	t.Run("bool to string", func(t *testing.T) {
		TestAssert(true, "true", t)
		TestAssert(false, "false", t)
	})
}
