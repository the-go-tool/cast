package cast

func registerBoolNumCast() {
	registerBoolInts()
	registerBoolUints()
	registerBoolFloats()
	registerBoolComplexes()
}

// casters:    [bool <-> ints]
// depends on: [ints <-> ints]
func registerBoolInts() {
	MustRegister(func(in int64) (out bool, err error) {
		return in != 0, nil
	})
	MustRegisterProxy[int32, bool, int64]()
	MustRegisterProxy[int16, bool, int64]()
	MustRegisterProxy[int8, bool, int64]()
	MustRegisterProxy[int, bool, int64]()

	MustRegister(func(in bool) (out int64, err error) {
		if in {
			return 1, nil
		}
		return 0, nil
	})
	MustRegisterProxy[bool, int32, int64]()
	MustRegisterProxy[bool, int16, int64]()
	MustRegisterProxy[bool, int8, int64]()
	MustRegisterProxy[bool, int, int64]()
}

// casters:    [bool <-> uints]
// depends on: [uints <-> uints]
func registerBoolUints() {
	MustRegister(func(in uint64) (out bool, err error) {
		return in != 0, nil
	})
	MustRegisterProxy[uint32, bool, uint64]()
	MustRegisterProxy[uint16, bool, uint64]()
	MustRegisterProxy[uint8, bool, uint64]()
	MustRegisterProxy[uint, bool, uint64]()

	MustRegister(func(in bool) (out uint64, err error) {
		if in {
			return 1, nil
		}
		return 0, nil
	})
	MustRegisterProxy[bool, uint32, uint64]()
	MustRegisterProxy[bool, uint16, uint64]()
	MustRegisterProxy[bool, uint8, uint64]()
	MustRegisterProxy[bool, uint, uint64]()
}

// casters:    [bool <-> floats]
// depends on: [floats <-> floats]
func registerBoolFloats() {
	MustRegister(func(in float64) (out bool, err error) {
		return in != 0, nil
	})
	MustRegisterProxy[float32, bool, float64]()

	MustRegister(func(in bool) (out float64, err error) {
		if in {
			return 1, nil
		}
		return 0, nil
	})
	MustRegisterProxy[bool, float32, float64]()
}

// casters:    [bool <-> complexes]
// depends on: [complexes <-> complexes]
func registerBoolComplexes() {
	MustRegister(func(in complex128) (out bool, err error) {
		return real(in) != 0, nil
	})
	MustRegisterProxy[complex64, bool, complex128]()

	MustRegister(func(in bool) (out complex128, err error) {
		if in {
			return complex(1, 0), nil
		}
		return complex(0, 0), nil
	})
	MustRegisterProxy[bool, complex64, complex128]()
}
