// package casters allows to manipulate from-to types casters.
package casters

import (
	"reflect"
	"sync"
)

// Caster describes function which implements casting way.
type Caster func(in reflect.Value) (out reflect.Value, err error)

// way describes 2key-value storage to store casters.
type way map[reflect.Type]map[reflect.Type]Caster

// casters is casters' storage which allows find specific caster with from-to-caster model.
var casters = way{}

// castersMutex allows work with methods of this module safely with multiple goroutines.
var castersMutex = sync.RWMutex{}

// Clear removes all data and all casting ways.
func Clear() {
	casters = way{}
}

// Set creates or replaces caster for specified from-to way.
func Set(from, to reflect.Type, cst Caster) {
	castersMutex.Lock()
	if castersTo, ok := casters[from]; ok {
		if _, ok := castersTo[to]; ok {
			castersTo[to] = cst
		}
	} else {
		casters[from] = map[reflect.Type]Caster{}
	}
	casters[from][to] = cst
	castersMutex.Unlock()
}

// Get returns caster or nil if not exists.
func Get(from, to reflect.Type) Caster {
	castersMutex.RLock()
	if castersTo, ok := casters[from]; ok {
		if cst, ok := castersTo[to]; ok {
			return cst
		}
	}
	castersMutex.RUnlock()
	return nil
}

// Exists checks that current from-to way is presented.
func Exist(from, to reflect.Type) bool {
	if cst := Get(from, to); cst != nil {
		return true
	}
	return false
}

// Remove deletes specified-way caster.
func Remove(from, to reflect.Type) {
	Set(from, to, nil)
}
