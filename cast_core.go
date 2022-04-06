package cast

import (
	"fmt"
	"reflect"
)

// tag is specific type's name.
type tag string

// casters contains set of ways to cast something from one type to other.
var casters = map[tag]map[tag]func(in reflect.Value) (out reflect.Value, err error){}

// tag returns type tag as string.
// For example:
//   fetchTag(reflect.TypeOf(int(5))) //> int
//   fetchTag(reflect.TypeOf(Custom(5))) //> main.Custom
func fetchTag(typ reflect.Type) tag {
	pkg := typ.PkgPath()
	name := typ.Name()
	if len(pkg) > 0 {
		return tag(pkg + "." + name)
	}
	return tag(name)
}

// caster finds and returns existing custer for specified types.
func caster(from reflect.Type, to reflect.Type) (cst func(in reflect.Value) (out reflect.Value, err error), err error) {
	fromT := fetchTag(from)
	toT := fetchTag(to)

	// A -> A
	if fromT == toT {
		return func(in reflect.Value) (out reflect.Value, err error) {
			return in, nil
		}, nil
	}

	// A -> B
	if castTo, ok := casters[fromT]; ok {
		if caster, ok := castTo[toT]; ok {
			return caster, nil
		}
	}

	return cst, fmt.Errorf("unknown caster: %s to %s", fromT, toT)
}

// cast general logic which helps to find and use available from-to casting function.
func cast(in reflect.Value, to reflect.Type) (out reflect.Value, err error) {
	cst, err := caster(in.Type(), to)
	if err != nil {
		return out, err
	}

	//TODO: unwrap in and wrap out at out
	out, err = cst(deref(in))
	if err != nil {
		return out, err
	}

	return out, nil
}

// register adds new caster.
func register(from reflect.Type, to reflect.Type, caster func(in reflect.Value) (out reflect.Value, err error)) error {
	fromTag := fetchTag(from)
	toTag := fetchTag(to)

	if castTo, ok := casters[fromTag]; ok {
		if _, ok := castTo[toTag]; ok {
			return fmt.Errorf("duplicate caster: %s to %s", from, to)
		}
	} else {
		casters[fromTag] = map[tag]func(in reflect.Value) (out reflect.Value, err error){}
	}
	casters[fromTag][toTag] = caster
	return nil
}

// registerProxy adds new caster from existing proxy custer.
// from -> proxy -> to.
func registerProxy(from reflect.Type, to reflect.Type, proxy reflect.Type) error {
	caster := func(in reflect.Value) (out reflect.Value, err error) {
		cstToProxy, err := caster(from, proxy)
		if err != nil {
			return out, err
		}
		cstToOut, err := caster(proxy, to)
		if err != nil {
			return out, err
		}
		valProxy, err := cstToProxy(in)
		if err != nil {
			return out, nil
		}
		valOut, err := cstToOut(valProxy)
		if err != nil {
			return out, err
		}
		return valOut, nil
	}
	return register(from, to, caster)
}

// deref accepts any referenced value and deep dereferences it.
func deref(val reflect.Value) reflect.Value {
	v := val
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
