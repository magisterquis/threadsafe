// tsmap implements threadsafe maps.
//
// Example:
//  m := Map(string, int)
//  m.Put("foo", 1)
//  fmt.Printf("%v: %v\n", "foo", m.Get("foo")) /* "foo: 1" */
package tsmap

import (
	"sync"
)

/* Map implements a threadsafe map. */
type Map struct {
	sync.RWMutex
	m map[string]interface{}
}

/* New makes a new map */
func New() *Map {
        m := &Map{}
        m.m = map[string]interface{}{}
        return m
}


/* Get returns the value pointed to by k */
func (m *Map) Get(k string) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.m[k]
	return v, ok

}

/* Put puts value v at key k in the Map. */
func (m *Map) Put(k string, v interface{}) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

/* PutUnique puts value v at key k only if k does not exist already, and
returns whether the put was successful or not. */
func (m *Map) PutUnique(k string, v interface{}) bool {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.m[k]; ok {
		return false
	}
	m.m[k] = v
	return true
}

/* Delete the key/value pair with key k. */
func (m *Map) Delete(k string) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

/* TODO: Function to apply function to value */
