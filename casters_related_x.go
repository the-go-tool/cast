package cast

func registerRelatedXCast() {
	registerCrossInts()
	registerCrossUints()
	registerCrossFloats()
	registerCrossComplexes()
}

// casters: [ints <-> ints]
func registerCrossInts() {
	MustRegister(func(in int64) (out int32, err error) {
		return int32(in), nil
	})
	MustRegister(func(in int64) (out int16, err error) {
		return int16(in), nil
	})
	MustRegister(func(in int64) (out int8, err error) {
		return int8(in), nil
	})
	MustRegister(func(in int64) (out int, err error) {
		return int(in), nil
	})

	MustRegister(func(in int32) (out int64, err error) {
		return int64(in), nil
	})
	MustRegisterProxy[int32, int16, int64]()
	MustRegisterProxy[int32, int8, int64]()
	MustRegisterProxy[int32, int, int64]()

	MustRegister(func(in int16) (out int64, err error) {
		return int64(in), nil
	})
	MustRegisterProxy[int16, int32, int64]()
	MustRegisterProxy[int16, int8, int64]()
	MustRegisterProxy[int16, int, int64]()

	MustRegister(func(in int8) (out int64, err error) {
		return int64(in), nil
	})
	MustRegisterProxy[int8, int32, int64]()
	MustRegisterProxy[int8, int16, int64]()
	MustRegisterProxy[int8, int, int64]()

	MustRegister(func(in int) (out int64, err error) {
		return int64(in), nil
	})
	MustRegisterProxy[int, int32, int64]()
	MustRegisterProxy[int, int16, int64]()
	MustRegisterProxy[int, int8, int64]()
}

// casters: [uints <-> uints]
func registerCrossUints() {
	MustRegister(func(in uint64) (out uint32, err error) {
		return uint32(in), nil
	})
	MustRegister(func(in uint64) (out uint16, err error) {
		return uint16(in), nil
	})
	MustRegister(func(in uint64) (out uint8, err error) {
		return uint8(in), nil
	})
	MustRegister(func(in uint64) (out uint, err error) {
		return uint(in), nil
	})

	MustRegister(func(in uint32) (out uint64, err error) {
		return uint64(in), nil
	})
	MustRegisterProxy[uint32, uint16, uint64]()
	MustRegisterProxy[uint32, uint8, uint64]()
	MustRegisterProxy[uint32, uint, uint64]()

	MustRegister(func(in uint16) (out uint64, err error) {
		return uint64(in), nil
	})
	MustRegisterProxy[uint16, uint32, uint64]()
	MustRegisterProxy[uint16, uint8, uint64]()
	MustRegisterProxy[uint16, uint, uint64]()

	MustRegister(func(in uint8) (out uint64, err error) {
		return uint64(in), nil
	})
	MustRegisterProxy[uint8, uint32, uint64]()
	MustRegisterProxy[uint8, uint16, uint64]()
	MustRegisterProxy[uint8, uint, uint64]()

	MustRegister(func(in uint) (out uint64, err error) {
		return uint64(in), nil
	})
	MustRegisterProxy[uint, uint32, uint64]()
	MustRegisterProxy[uint, uint16, uint64]()
	MustRegisterProxy[uint, uint8, uint64]()
}

// casters: [floats <-> floats]
func registerCrossFloats() {
	MustRegister(func(in float64) (out float32, err error) {
		return float32(in), nil
	})

	MustRegister(func(in float32) (out float64, err error) {
		return float64(in), nil
	})
}

// casters: [complexes <-> complexes]
func registerCrossComplexes() {
	MustRegister(func(in complex128) (out complex64, err error) {
		return complex64(in), nil
	})

	MustRegister(func(in complex64) (out complex128, err error) {
		return complex128(in), nil
	})
}
