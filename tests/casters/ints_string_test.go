package casting_tests

import (
	"testing"

	"github.com/the-go-tool/cast"
)

func TestIntsString(t *testing.T) {
	cast.TestAssert(int(-69), "-69", t)
	cast.TestAssert(int8(-69), "-69", t)
	cast.TestAssert(int16(-69), "-69", t)
	cast.TestAssert(int32(-69), "-69", t)
	cast.TestAssert(int64(-69), "-69", t)
}
