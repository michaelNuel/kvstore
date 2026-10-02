package main;  


import (
	"sync"
)

type Testing struct {
	mu sync.Mutex
	store map[string]string
}


func (t *Testing) Set(key, value  string ) {
   t.mu.Lock()
   defer t.mu.Unlock()

   if t.store == nil {
	t.store = make(map[string]string)
   }

   t.store[key] = value
}


func (t *Testing) Get(key string) string {
  t.mu.Lock()
  defer t.mu.Unlock()

  return t.store[key]
}
