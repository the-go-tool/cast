package cast

import (
	"testing"
)

func TestStringUints(t *testing.T) {
	TestAssert("69", uint(69), t)
	TestAssert("69", uint8(69), t)
	TestAssert("69", uint16(69), t)
	TestAssert("69", uint32(69), t)
	TestAssert("69", uint64(69), t)
}
