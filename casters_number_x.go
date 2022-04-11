package cast

func registerNumberXCast() {
	registerIntsUints()
	registerIntsFloats()
	registerUintsFloats()
	registerFloatsComplexes()
	registerIntsComplexes()
	registerUintsComplexes()
}

// casters:    [ints <-> uints]
// depends on: [ints <-> ints], [uints <-> uints]
func registerIntsUints() {
	MustRegister(func(in int64) (out uint64, err error) {
		return uint64(in), nil
	})
	MustRegisterProxy[int64, uint, uint64]()
	MustRegisterProxy[int64, uint8, uint64]()
	MustRegisterProxy[int64, uint16, uint64]()
	MustRegisterProxy[int64, uint32, uint64]()

	MustRegisterProxy[int32, uint, int64]()
	MustRegisterProxy[int32, uint8, int64]()
	MustRegisterProxy[int32, uint16, int64]()
	MustRegisterProxy[int32, uint32, int64]()
	MustRegisterProxy[int32, uint64, int64]()

	MustRegisterProxy[int16, uint, int64]()
	MustRegisterProxy[int16, uint8, int64]()
	MustRegisterProxy[int16, uint16, int64]()
	MustRegisterProxy[int16, uint32, int64]()
	MustRegisterProxy[int16, uint64, int64]()

	MustRegisterProxy[int8, uint, int64]()
	MustRegisterProxy[int8, uint8, int64]()
	MustRegisterProxy[int8, uint16, int64]()
	MustRegisterProxy[int8, uint32, int64]()
	MustRegisterProxy[int8, uint64, int64]()

	MustRegisterProxy[int, uint, int64]()
	MustRegisterProxy[int, uint8, int64]()
	MustRegisterProxy[int, uint16, int64]()
	MustRegisterProxy[int, uint32, int64]()
	MustRegisterProxy[int, uint64, int64]()

	MustRegister(func(in uint64) (out int64, err error) {
		return int64(in), nil
	})
	MustRegisterProxy[uint64, int, int64]()
	MustRegisterProxy[uint64, int8, int64]()
	MustRegisterProxy[uint64, int16, int64]()
	MustRegisterProxy[uint64, int32, int64]()

	MustRegisterProxy[uint32, int, uint64]()
	MustRegisterProxy[uint32, int8, uint64]()
	MustRegisterProxy[uint32, int16, uint64]()
	MustRegisterProxy[uint32, int32, uint64]()
	MustRegisterProxy[uint32, int64, uint64]()

	MustRegisterProxy[uint16, int, uint64]()
	MustRegisterProxy[uint16, int8, uint64]()
	MustRegisterProxy[uint16, int16, uint64]()
	MustRegisterProxy[uint16, int32, uint64]()
	MustRegisterProxy[uint16, int64, uint64]()

	MustRegisterProxy[uint8, int, uint64]()
	MustRegisterProxy[uint8, int8, uint64]()
	MustRegisterProxy[uint8, int16, uint64]()
	MustRegisterProxy[uint8, int32, uint64]()
	MustRegisterProxy[uint8, int64, uint64]()

	MustRegisterProxy[uint, int, uint64]()
	MustRegisterProxy[uint, int8, uint64]()
	MustRegisterProxy[uint, int16, uint64]()
	MustRegisterProxy[uint, int32, uint64]()
	MustRegisterProxy[uint, int64, uint64]()
}

// casters:    [ints <-> floats]
// depends on: [ints <-> ints]
func registerIntsFloats() {
	MustRegister(func(in int64) (out float64, err error) {
		return float64(in), nil
	})
	MustRegisterProxy[int32, float64, int64]()
	MustRegisterProxy[int16, float64, int64]()
	MustRegisterProxy[int8, float64, int64]()
	MustRegisterProxy[int, float64, int64]()

	MustRegisterProxy[int64, float32, float64]()
	MustRegisterProxy[int32, float32, float64]()
	MustRegisterProxy[int16, float32, float64]()
	MustRegisterProxy[int8, float32, float64]()
	MustRegisterProxy[int, float32, float64]()

	MustRegister(func(in float64) (out int64, err error) {
		return int64(in), nil
	})
	MustRegisterProxy[float64, int32, int64]()
	MustRegisterProxy[float64, int16, int64]()
	MustRegisterProxy[float64, int8, int64]()
	MustRegisterProxy[float64, int, int64]()

	MustRegisterProxy[float32, int64, float64]()
	MustRegisterProxy[float32, int32, float64]()
	MustRegisterProxy[float32, int16, float64]()
	MustRegisterProxy[float32, int8, float64]()
	MustRegisterProxy[float32, int, float64]()
}

// casters:    [uints <-> floats]
// depends on: [uints <-> uints]
func registerUintsFloats() {
	MustRegister(func(in uint64) (out float64, err error) {
		return float64(in), nil
	})
	MustRegisterProxy[uint32, float64, uint64]()
	MustRegisterProxy[uint16, float64, uint64]()
	MustRegisterProxy[uint8, float64, uint64]()
	MustRegisterProxy[uint, float64, uint64]()

	MustRegisterProxy[uint64, float32, float64]()
	MustRegisterProxy[uint32, float32, float64]()
	MustRegisterProxy[uint16, float32, float64]()
	MustRegisterProxy[uint8, float32, float64]()
	MustRegisterProxy[uint, float32, float64]()

	MustRegister(func(in float64) (out uint64, err error) {
		return uint64(in), nil
	})
	MustRegisterProxy[float64, uint32, uint64]()
	MustRegisterProxy[float64, uint16, uint64]()
	MustRegisterProxy[float64, uint8, uint64]()
	MustRegisterProxy[float64, uint, uint64]()

	MustRegisterProxy[float32, uint64, float64]()
	MustRegisterProxy[float32, uint32, float64]()
	MustRegisterProxy[float32, uint16, float64]()
	MustRegisterProxy[float32, uint8, float64]()
	MustRegisterProxy[float32, uint, float64]()
}

// casters: [floats <-> complexes]
// depends on: [floats <-> floats], [complexes <-> complexes]
func registerFloatsComplexes() {
	MustRegister(func(in float64) (out complex128, err error) {
		return complex(in, 0), nil
	})
	MustRegisterProxy[float32, complex128, float64]()

	MustRegisterProxy[float64, complex64, complex128]()
	MustRegisterProxy[float32, complex64, complex128]()

	MustRegister(func(in complex128) (out float64, err error) {
		return real(in), nil
	})
	MustRegisterProxy[complex64, float64, complex128]()

	MustRegisterProxy[complex128, float32, float64]()
	MustRegisterProxy[complex64, float32, float64]()
}

// casters:    [ints <-> complexes]
// depends on: [ints <-> ints]
func registerIntsComplexes() {
	MustRegister(func(in int64) (out complex128, err error) {
		return complex128(complex(float64(in), 0)), nil
	})
	MustRegisterProxy[int32, complex128, int64]()
	MustRegisterProxy[int16, complex128, int64]()
	MustRegisterProxy[int8, complex128, int64]()
	MustRegisterProxy[int, complex128, int64]()

	MustRegister(func(in int64) (out complex64, err error) {
		return complex64(complex(float64(in), 0)), nil
	})
	MustRegisterProxy[int32, complex64, int64]()
	MustRegisterProxy[int16, complex64, int64]()
	MustRegisterProxy[int8, complex64, int64]()
	MustRegisterProxy[int, complex64, int64]()

	MustRegister(func(in complex128) (out int64, err error) {
		return int64(real(in)), nil
	})
	MustRegisterProxy[complex128, int32, int64]()
	MustRegisterProxy[complex128, int16, int64]()
	MustRegisterProxy[complex128, int8, int64]()
	MustRegisterProxy[complex128, int, int64]()

	MustRegister(func(in complex64) (out int64, err error) {
		return int64(real(in)), nil
	})
	MustRegisterProxy[complex64, int32, int64]()
	MustRegisterProxy[complex64, int16, int64]()
	MustRegisterProxy[complex64, int8, int64]()
	MustRegisterProxy[complex64, int, int64]()
}

// casters:    [uints <-> complexes]
// depends on: [uints <-> uints]
func registerUintsComplexes() {
	MustRegister(func(in uint64) (out complex128, err error) {
		return complex128(complex(float64(in), 0)), nil
	})
	MustRegisterProxy[uint32, complex128, uint64]()
	MustRegisterProxy[uint16, complex128, uint64]()
	MustRegisterProxy[uint8, complex128, uint64]()
	MustRegisterProxy[uint, complex128, uint64]()

	MustRegister(func(in uint64) (out complex64, err error) {
		return complex64(complex(float64(in), 0)), nil
	})
	MustRegisterProxy[uint32, complex64, uint64]()
	MustRegisterProxy[uint16, complex64, uint64]()
	MustRegisterProxy[uint8, complex64, uint64]()
	MustRegisterProxy[uint, complex64, uint64]()

	MustRegister(func(in complex128) (out uint64, err error) {
		return uint64(real(in)), nil
	})
	MustRegisterProxy[complex128, uint32, uint64]()
	MustRegisterProxy[complex128, uint16, uint64]()
	MustRegisterProxy[complex128, uint8, uint64]()
	MustRegisterProxy[complex128, uint, uint64]()

	MustRegister(func(in complex64) (out uint64, err error) {
		return uint64(real(in)), nil
	})
	MustRegisterProxy[complex64, uint32, uint64]()
	MustRegisterProxy[complex64, uint16, uint64]()
	MustRegisterProxy[complex64, uint8, uint64]()
	MustRegisterProxy[complex64, uint, uint64]()
}
