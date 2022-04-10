package types

import "reflect"

// GetName returns full-named Type.
//
//  type Custom int
//  GetName(reflect.TypeOf(int(5))) //> int
//  GetName(reflect.TypeOf(Custom(5))) //> package.Custom
//  GetName(reflect.TypeOf(nil)) //> interface
func GetName(t reflect.Type) string {
	if t == nil {
		return "interface"
	}
	pkg := t.PkgPath()
	name := t.Name()
	if len(pkg) > 0 {
		return pkg + "." + name
	}
	return name
}

// GetKind returns reflect.Kind.
//
//  type Custom int
//  Get(reflect.TypeOf(int(5))) //> int
//  Get(reflect.TypeOf(Custom(5))) //> int
//  Get(reflect.TypeOf(nil)) //> interface
func GetKind(t reflect.Type) reflect.Kind {
	if t == nil {
		return reflect.Interface
	}
	return t.Kind()
}

// FromKind returns new type from kind name.
// It's useful for basetype search.
//
//  FromKind("int") //> reflect.Int
//  FromKind("float64") //> reflect.Float64
func FromKind(k reflect.Kind) reflect.Type {
	mapping := map[reflect.Kind]any{
		reflect.String:     string(""),
		reflect.Bool:       bool(false),
		reflect.Complex64:  complex64(0),
		reflect.Complex128: complex128(0),
		reflect.Float32:    float32(0),
		reflect.Float64:    float64(0),
		reflect.Int:        int(0),
		reflect.Int8:       int8(0),
		reflect.Int16:      int16(0),
		reflect.Int32:      int32(0),
		reflect.Int64:      int64(0),
		reflect.Uint:       uint(0),
		reflect.Uint8:      uint8(0),
		reflect.Uint16:     uint16(0),
		reflect.Uint32:     uint32(0),
		reflect.Uint64:     uint64(0),
	}

	t, ok := mapping[k]
	if !ok {
		return nil
	}

	return reflect.TypeOf(t)
}

// IsInt returns flag that indicates if this type can be converted to int.
//
//  type Custom int
//  IsInt(5) //> true
//  IsInt(Custom(5)) //> true
//  IsInt("5") //> false
func IsInt(t reflect.Type) bool {
	return reflect.New(t).CanInt()
}

// ToInt converts any intable type to int value.
// t must be one of int, int8, int16, int32 or int64.
// In case of incorrect types result will be reflect.Value(nil).
// Check it by IsInt before using.
//
//  ToInt(reflect.ValueOf(int(5)), reflect.Int64) //> reflect.Value(int64(5))
func ToInt(v reflect.Value, t reflect.Kind) reflect.Value {
	if v.CanInt() {
		switch v.Kind() {
		case reflect.Int:
			return reflect.ValueOf(int(v.Int()))
		case reflect.Int8:
			return reflect.ValueOf(int8(v.Int()))
		case reflect.Int16:
			return reflect.ValueOf(int16(v.Int()))
		case reflect.Int32:
			return reflect.ValueOf(int32(v.Int()))
		case reflect.Int64:
			return reflect.ValueOf(int64(v.Int()))
		}
	}
	return reflect.ValueOf(nil)
}

// IsUint returns flag that indicates if this type can be converted to uint.
//
//  type Custom uint
//  IsUint(5) //> true
//  IsUint(Custom(5)) //> true
//  IsUint("5") //> false
func IsUint(t reflect.Type) bool {
	return reflect.New(t).CanUint()
}

// ToUint converts any intable type to uint value.
// t must be one of uint, uint8, uint16, uint32 or uint64.
// In case of incorrect types result will be reflect.Value(nil).
// Check it by IsUint before using.
//
//  ToUint(reflect.ValueOf(uint(5)), reflect.Uint64) //> reflect.Value(uint64(5))
func ToUint(v reflect.Value, t reflect.Kind) reflect.Value {
	if v.CanUint() {
		switch v.Kind() {
		case reflect.Uint:
			return reflect.ValueOf(uint(v.Uint()))
		case reflect.Uint8:
			return reflect.ValueOf(uint8(v.Uint()))
		case reflect.Uint16:
			return reflect.ValueOf(uint16(v.Uint()))
		case reflect.Uint32:
			return reflect.ValueOf(uint32(v.Uint()))
		case reflect.Uint64:
			return reflect.ValueOf(uint64(v.Uint()))
		}
	}
	return reflect.ValueOf(nil)
}

// IsFloat returns flag that indicates if this type can be converted to float.
//
//  type Custom float
//  IsFloat(5.4) //> true
//  IsFloat(Custom(5.4)) //> true
//  IsFloat("5.4") //> false
func IsFloat(t reflect.Type) bool {
	return reflect.New(t).CanFloat()
}

// ToFloat converts any intable type to float value.
// t must be one of float32 or float64.
// In case of incorrect types result will be reflect.Value(nil).
// Check it by IsFloat before using.
//
//  ToFloat(reflect.ValueOf(float32(5)), reflect.Float64) //> reflect.Value(float64(5))
func ToFloat(v reflect.Value, t reflect.Kind) reflect.Value {
	if v.CanFloat() {
		switch v.Kind() {
		case reflect.Float32:
			return reflect.ValueOf(float32(v.Float()))
		case reflect.Float64:
			return reflect.ValueOf(float64(v.Float()))
		}
	}
	return reflect.ValueOf(nil)
}
