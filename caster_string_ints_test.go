package cast

import (
	"testing"
)

func TestStringInts(t *testing.T) {
	TestAssert("-69", int(-69), t)
	TestAssert("-69", int8(-69), t)
	TestAssert("-69", int16(-69), t)
	TestAssert("-69", int32(-69), t)
	TestAssert("-69", int64(-69), t)
}
