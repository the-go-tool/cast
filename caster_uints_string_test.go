package cast

import (
	"testing"
)

func TestUintsString(t *testing.T) {
	TestAssert(uint(69), "69", t)
	TestAssert(uint8(69), "69", t)
	TestAssert(uint16(69), "69", t)
	TestAssert(uint32(69), "69", t)
	TestAssert(uint64(69), "69", t)
}
