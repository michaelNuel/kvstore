package main;  


import (
	"sync"
   "time"
)

type Version struct {
   data string 
   date time.Time
}

type Testing struct {
	mu sync.Mutex
	store map[string][]Version
}


func (t *Testing) Set(key, value  string ) {
   t.mu.Lock()
   defer t.mu.Unlock()

   if t.store == nil {
	t.store = make(map[string][]Version) 
   }

   t.store[key] = append(t.store[key], Version{data: value, date: time.Now()})  
}


func (t *Testing) Get(key string) string {
  t.mu.Lock()
  defer t.mu.Unlock()

  versions := t.store[key]
  if len(versions) == 0 {
	return ""
  }
  return versions[len(versions)-1].data
}
