package cast

import (
	"fmt"
	"reflect"
)

// Castable contract of values that can be casted each other.
// Byte supports also by ~uint8.
type Castable interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// casters contains set of ways to cast something from one type to other.
var casters = map[reflect.Kind]map[reflect.Kind]func(in reflect.Value) (out reflect.Value, err error){}

// cast general logic which helps to find and use available from-to casting function.
func cast[Out, In any](in In) (out Out, err error) {
	valueIn := reflect.ValueOf(deref(in))
	valueOut := reflect.ValueOf(deref(out))
	kindIn := valueIn.Kind()
	kindOut := valueOut.Kind()

	if kindIn == kindOut {
		return valueIn.Interface().(Out), nil
	}

	if castTo, ok := casters[kindIn]; ok {
		if caster, ok := castTo[kindOut]; ok {
			valueCasted, err := caster(valueIn)
			if err != nil {
				return *new(Out), err
			}
			return valueCasted.Interface().(Out), nil
		}
	}

	return *new(Out), fmt.Errorf("unsupported cast from %s to %s", kindIn, kindOut)
}

// register adds new caster.
func register(from reflect.Kind, to reflect.Kind, caster func(in reflect.Value) (out reflect.Value, err error)) {
	if castTo, ok := casters[from]; ok {
		if _, ok := castTo[to]; ok {
			panic(fmt.Sprintf("already defined cast from %s to %s", from, to))
		}
	} else {
		casters[from] = map[reflect.Kind]func(in reflect.Value) (out reflect.Value, err error){}
	}
	casters[from][to] = caster
}

// deref accepts any referenced value and deep dereferences it.
func deref(val any) any {
	vval := reflect.ValueOf(val)
	for vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	return vval.Interface()
}

// WithError casts one type to other and returns value and error.
func WithError[Out, In Castable](in In) (Out, error) {
	return cast[Out](in)
}

// To casts one type to other or returns default value for type in case of error.
func To[Out, In Castable](in Out) Out {
	val, _ := WithError[Out](in)
	return val
}

// MustTo casts one type to type and panics in case of error.
func MustTo[Out, In Castable](in Out) Out {
	val, err := WithError[Out](in)
	if err != nil {
		panic(err)
	}
	return val
}

// Possible tries cast one type to other and returns casting possibility flag.
func Possible[Out, In Castable](in Out) bool {
	if _, err := WithError[Out](in); err != nil {
		return false
	}
	return true
}
