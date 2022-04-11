// package catchers allows to manipulate post-casters handlers.
// It suitable to work with some data if no one casters found for this type pair.
package catchers

import (
	"fmt"
	"reflect"
	"sync"
)

const (
	// NotSuitableError special error to indicate that this catcher isn't suitable
	// for current case and executor must use next catcher from queue.
	NotSuitableError = "not_suitable"
)

// Catcher describes function which implements catcher to work with unknown or incastable types.
type Catcher func(in reflect.Value, to reflect.Type) (out reflect.Value, err error)

// catchers list of handlers.
var catchers = make([]Catcher, 0, 8)

// catchersMutex allows work with methods of this module safely with multiple goroutines.
var catchersMutex = sync.RWMutex{}

// Clear removes all data and all catchers.
func Clear() {
	catchers = make([]Catcher, 0, 8)
}

// Add adds new catcher in the list.
func Add(catch Catcher) {
	catchersMutex.Lock()
	catchers = append(catchers, catch)
	catchersMutex.Unlock()
}

// Process step by step executes catchers from queue to find suitable and return they result.
func Process(in reflect.Value, to reflect.Type) (out reflect.Value, err error) {
	for _, catch := range catchers {
		out, err := catch(in, to)
		if err != nil {
			if err.Error() == NotSuitableError {
				continue
			}
			return out, err
		}
		return out, nil
	}
	return out, fmt.Errorf("no suitable catchers for [%v(%v) -> %v]", in.Type(), in.Interface(), to)
}
