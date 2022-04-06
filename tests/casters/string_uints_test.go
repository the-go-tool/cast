package casting_tests

import (
	"testing"

	"github.com/the-go-tool/cast"
)

func TestStringUints(t *testing.T) {
	cast.TestAssert("69", uint(69), t)
	cast.TestAssert("69", uint8(69), t)
	cast.TestAssert("69", uint16(69), t)
	cast.TestAssert("69", uint32(69), t)
	cast.TestAssert("69", uint64(69), t)
}
