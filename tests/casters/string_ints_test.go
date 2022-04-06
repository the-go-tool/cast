package casting_tests

import (
	"testing"

	"github.com/the-go-tool/cast"
)

func TestStringInts(t *testing.T) {
	cast.TestAssert("-69", int(-69), t)
	cast.TestAssert("-69", int8(-69), t)
	cast.TestAssert("-69", int16(-69), t)
	cast.TestAssert("-69", int32(-69), t)
	cast.TestAssert("-69", int64(-69), t)
}
