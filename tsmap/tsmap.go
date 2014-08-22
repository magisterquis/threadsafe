// tsmap implements threadsafe maps.
//
// Example:
//  m := Map(string, int)
//  m.Put("foo", 1)
//  fmt.Printf("%v: %v\n", "foo", m.Get("foo")) /* "foo: 1" */
package threadsafe/tsmap

import (
        "errors"
        "sync"
)

/* Map implements a threadsafe map. */
type Map {
        sync.RWMutex
        m map[interface{}]interface{}
        k type /* Key type */
        v type /* Value type */
}

/* New creates a new Map, with keys of type k and values of type v. */
func New(k, v type) *Map {
        m := &Map{k: k, v: v}
        return m
}

/* Get returns the value pointed to by k, which must be the type specified
when the Map was created. */
func (m *map) Get(k interface{}) (interface{}, bool, error) {
        m.RLock()
        defer m.RUnlock()
        /* Typecast */
        if v, ok := k.(m.k); !ok {
                return nil, errors.New("Invalid type %T for key %v; "+
                "expected %v", k, k, m.k)
        }
        v, ok := m[k]
        return v, ok, nil
}

/* Put puts value v at key k in the Map.  Both values are type-checked before
insertion. */
func (m *Map) Put(k, v interface{}) error {
        m.Lock()
        defer m.Unlock()
        if _, ok := k.(m.k); !ok {
                return errors.New("Invalid type %T for key %v; expected %v",
                k, k, m.k)
        }
        if _, ok := v.(m.v); !ok {
                return errors.New("Invalid type %T for value %v; expected %v",
                v, v, m.v)
        }
        m.m[k] = v
        return nil
}

/* Delete the key/value pair with key k. */
func (m *Map)Delete(k) error {
m.Lock()
defer m.Unlock()
        if _, ok := k.(m.k); !ok {
                return errors.New("Invalid type %T for key %v; expected %v",
                k, k, m.k)
        }
        delete(m, k)
        return nil
}



