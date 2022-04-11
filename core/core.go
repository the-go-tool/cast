package core

import (
	"fmt"
	"reflect"

	"github.com/the-go-tool/cast/core/casters"
	"github.com/the-go-tool/cast/core/catchers"
)

// Cast general logic which helps to find and use available from-to casting function.
func Cast(in reflect.Value, to reflect.Type) (out reflect.Value, err error) {
	// Regular casting way search
	cst := casters.Get(in.Type(), to)
	if cst != nil {
		return cst(in)
	}

	// Find and execute suitable catcher
	out, err = catchers.Process(in, to)
	if err == nil {
		return out, nil
	}

	if cst == nil {
		return out, fmt.Errorf("unknown caster [%s -> %s]", in.Type(), to)
	}

	return out, nil
}

// Register adds new caster.
func Register(from reflect.Type, to reflect.Type, cst casters.Caster) error {
	if !casters.Exist(from, to) {
		casters.Set(from, to, cst)
		return nil
	}
	return fmt.Errorf("duplicate caster [%s -> %s]", from, to)
}

// RegisterProxy adds new caster from existing proxy custer.
// from -> proxy -> to.
func RegisterProxy(from reflect.Type, to reflect.Type, proxy reflect.Type) error {
	caster := func(in reflect.Value) (out reflect.Value, err error) {
		proxyV, err := Cast(in, proxy)
		if err != nil {
			return out, err
		}
		outV, err := Cast(proxyV, to)
		if err != nil {
			return out, err
		}
		return outV, nil
	}
	return Register(from, to, caster)
}
