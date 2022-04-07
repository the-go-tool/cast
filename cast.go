package cast

import (
	"fmt"
	"reflect"
	"testing"
)

// Register adds new way to cast one type to another.
//
//   type Custom int
//   // ...
//   err := Register(func(in Custom) (int, error) {
//     return int(Custom), nil
//   })
//   // ...
//   c := Custom(5)
//   i := To[int](c) // i is int(5)
func Register[In, Out any](caster func(in In) (out Out, err error)) error {
	inT := reflect.TypeOf(*new(In))
	outT := reflect.TypeOf(*new(Out))

	return register(inT, outT, func(in reflect.Value) (out reflect.Value, err error) {
		val, err := caster(in.Interface().(In))
		if err != nil {
			return out, err
		}
		return reflect.ValueOf(val), nil
	})
}

// MustRegister adds new way to cast one type to another.
// It panics if caster with this pair of types already registered.
//
//   type Custom int
//   // ...
//   MustRegister(func(in Custom) (int, error) {
//     return int(Custom), nil
//   })
//   // ...
//   c := Custom(5)
//   i := To[int](c) // i is int(5)
func MustRegister[In, Out any](caster func(in In) (out Out, err error)) {
	if err := Register(caster); err != nil {
		panic(err)
	}
}

// RegisterProxy adds new way to cast one type to another through existing ways.
// If you have already registered casters for A->B and B->C then this proxy
// allows to register A->C casting way.
//
//   type Custom int
//   // ...
//   MustRegister(func(in Custom) (int, error) {
//     return int(Custom), nil
//   })
//   err := RegisterProxy[Custom, string, int]() // Custom->int->string
//   // ...
//   c := Custom(5)
//   s := To[string](c) // s is string("5")
func RegisterProxy[In, Out, Proxy any]() error {
	inT := reflect.TypeOf(*new(In))
	outT := reflect.TypeOf(*new(Out))
	proxyT := reflect.TypeOf(*new(Proxy))
	return registerProxy(inT, outT, proxyT)
}

// MustRegisterProxy adds new way to cast one type to another through existing ways.
// It panics if caster with this types already registered.
// If you have already registered casters for A->B and B->C then this proxy
// allows to register A->C casting way.
//
//   type Custom int
//   // ...
//   MustRegister(func(in Custom) (int, error) {
//     return int(Custom), nil
//   })
//   MustRegisterProxy[Custom, string, int]() // Custom->int->string
//   // ...
//   c := Custom(5)
//   s := To[string](c) // s is string("5")
func MustRegisterProxy[In, Out, Proxy any]() {
	if err := RegisterProxy[In, Out, Proxy](); err != nil {
		panic(err)
	}
}

// Assert checks that one type can be cast to another with correct value.
//
//   // ok, it's possible casting and valid value
//   err := Assert(5, "5")
//   // err, it's possible casting but invalid value
//   err := Assert(5, "6")
//   // err, unknown casting way Custom->int
//   err := Assert(Custom(5), 5)
func Assert[In, Out any](in In, out Out) error {
	val, err := WithError[Out](in)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(val, out) {
		return fmt.Errorf("expected %#v, got %#v from %#v", out, val, in)
	}
	return nil
}

// MustAssert checks that one type can be cast to another with correct value.
// It panics in case of error.
//
//   // ok, it's possible casting and valid value
//   MustAssert(5, "5")
//   // panic, it's possible casting but invalid value
//   MustAssert(5, "6")
//   // panic, unknown casting way Custom->int
//   MustAssert(Custom(5), 5)
func MustAssert[In, Out any](in In, out Out) {
	if err := Assert(in, out); err != nil {
		panic(err)
	}
}

// TestAssert checks that one type can be cast to another with correct value.
// It calls t.Fatal in case of error.
//
//   // ok, it's possible casting and valid value
//   TestAssert(5, "5", t)
//   // test failed, it's possible casting but invalid value
//   TestAssert(5, "6", t)
//   // test failed, unknown casting way Custom->int
//   TestAssert(Custom(5), 5, t)
func TestAssert[In, Out any](in In, out Out, t *testing.T) {
	name := fmt.Sprintf("cast %#v to %#v", in, out)
	t.Run(name, func(t *testing.T) {
		if err := Assert(in, out); err != nil {
			t.Fatal(err)
		}
	})
}

// WithError casts one type to other and returns value and error.
//
//   val, err := WithError[int]("5") // int(5), nil
//   val, err := WithError[int]("text") // int(0), "can't parse"
//   val, err := WithError[Custom]("5") // Custom{}, "unknown caster"
func WithError[Out any](in any) (Out, error) {
	out := *new(Out)
	val, err := cast(reflect.ValueOf(in), reflect.TypeOf(out))
	if err != nil {
		return out, err
	}
	return val.Interface().(Out), nil
}

// To casts one type to other.
// Returns value but ignores error.
//
//   val := To[int]("5") // int(5)
//   val := To[int]("text") // int(0)
//   val := To[Custom]("5") // Custom{}
func To[Out any](in any) Out {
	val, _ := WithError[Out](in)
	return val
}

// MustTo casts one type to other.
// Returns value but ignores error.
// It panics in case of error.
//
//   val := MustTo[int]("5") // int(5)
//   val := MustTo[int]("text") // panic
//   val := MustTo[Custom]("5") // panic
func MustTo[Out any](in any) Out {
	val, err := WithError[Out](in)
	if err != nil {
		panic(err)
	}
	return val
}

// Possible checks casting possibility.
// It check that casting way registered, and specified value can be cast.
//
//   val := Possible[int]("5") // true
//   val := Possible[int]("text") // false
//   val := Possible[Custom]("5") // false
func Possible[Out any](in any) bool {
	if _, err := WithError[Out](in); err != nil {
		return false
	}
	return true
}
