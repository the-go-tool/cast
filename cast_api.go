package cast

import (
	"fmt"
	"reflect"
	"testing"
)

// Register adds new caster.
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

// MustRegister add new caster.
// It panics if caster with this types already defined.
func MustRegister[In, Out any](caster func(in In) (out Out, err error)) {
	if err := Register(caster); err != nil {
		panic(err)
	}
}

// registerProxy adds new caster from existing proxy custer.
// from -> proxy -> to.
func RegisterProxy[In, Out, Proxy any]() error {
	inT := reflect.TypeOf(*new(In))
	outT := reflect.TypeOf(*new(Out))
	proxyT := reflect.TypeOf(*new(Proxy))
	return registerProxy(inT, outT, proxyT)
}

// registerProxy adds new caster from existing proxy custer.
// from -> proxy -> to.
// It panics if caster with this types already defined.
func MustRegisterProxy[In, Out, Proxy any]() {
	if err := RegisterProxy[In, Out, Proxy](); err != nil {
		panic(err)
	}
}

// Assert checks that one value can be cast to other with correct value data.
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

// Assert checks that one value can be cast to other with correct value data.
// It panics in case of error.
func MustAssert[In, Out any](in In, out Out) {
	if err := Assert(in, out); err != nil {
		panic(err)
	}
}

// TestAssert checks that one value can be cast to other with correct value data.
func TestAssert[In, Out any](in In, out Out, t *testing.T) {
	name := fmt.Sprintf("cast %#v to %#v", in, out)
	t.Run(name, func(t *testing.T) {
		if err := Assert(in, out); err != nil {
			t.Fatal(err)
		}
	})
}

// WithError casts one type to other and returns value and error.
func WithError[Out any](in any) (Out, error) {
	out := *new(Out)
	val, err := cast(reflect.ValueOf(in), reflect.TypeOf(out))
	if err != nil {
		return out, err
	}
	return val.Interface().(Out), nil
}

// To casts one type to other or returns default value for type in case of error.
func To[Out any](in any) Out {
	val, _ := WithError[Out](in)
	return val
}

// MustTo casts one type to type and panics in case of error.
func MustTo[Out any](in any) Out {
	val, err := WithError[Out](in)
	if err != nil {
		panic(err)
	}
	return val
}

// Possible tries cast one type to other and returns casting possibility flag.
func Possible[Out any](in any) bool {
	if _, err := WithError[Out](in); err != nil {
		return false
	}
	return true
}
