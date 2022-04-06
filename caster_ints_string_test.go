package cast

import (
	"testing"
)

func TestIntsString(t *testing.T) {
	TestAssert(int(-69), "-69", t)
	TestAssert(int8(-69), "-69", t)
	TestAssert(int16(-69), "-69", t)
	TestAssert(int32(-69), "-69", t)
	TestAssert(int64(-69), "-69", t)
}
