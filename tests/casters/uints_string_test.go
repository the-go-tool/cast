package casting_tests

import (
	"testing"

	"github.com/the-go-tool/cast"
)

func TestUintsString(t *testing.T) {
	cast.TestAssert(uint(69), "69", t)
	cast.TestAssert(uint8(69), "69", t)
	cast.TestAssert(uint16(69), "69", t)
	cast.TestAssert(uint32(69), "69", t)
	cast.TestAssert(uint64(69), "69", t)
}
