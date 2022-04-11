package cast

import "testing"

func TestCastersRelatedX_CrossInts(t *testing.T) {
	for i := -1; i <= 1; i++ {
		ints := []any{int(i), int8(i), int16(i), int32(i), int64(i)}
		for _, v := range ints {
			TestAssert(v, int(i), t)
			TestAssert(v, int8(i), t)
			TestAssert(v, int16(i), t)
			TestAssert(v, int32(i), t)
			TestAssert(v, int64(i), t)
		}
	}
}

func TestCastersRelatedX_CrossUints(t *testing.T) {
	for i := 0; i <= 1; i++ {
		ints := []any{uint(i), uint8(i), uint16(i), uint32(i), uint64(i)}
		for _, v := range ints {
			TestAssert(v, uint(i), t)
			TestAssert(v, uint8(i), t)
			TestAssert(v, uint16(i), t)
			TestAssert(v, uint32(i), t)
			TestAssert(v, uint64(i), t)
		}
	}
}

func TestCastersRelatedX_CrossFloats(t *testing.T) {
	for i := float64(-1); i <= 1; i += 0.5 {
		ints := []any{float32(i), float64(i)}
		for _, v := range ints {
			TestAssert(v, float32(i), t)
			TestAssert(v, float64(i), t)
		}
	}
}

func TestCastersRelatedX_CrossComplexes(t *testing.T) {
	for i := float64(-1); i <= 1; i += 0.5 {
		for j := float64(-1); i <= 1; i += 0.5 {
			ints := []any{complex64(complex(i, j)), complex128(complex(i, j))}
			for _, v := range ints {
				TestAssert(v, complex64(complex(i, j)), t)
				TestAssert(v, complex128(complex(i, j)), t)
			}
		}
	}
}
