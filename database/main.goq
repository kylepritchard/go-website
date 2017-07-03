package main

import (
	"fmt"
	"sync"

	uuid "github.com/nu7hatch/gouuid"

	"./packages"
)

type Store struct {
	posts map[string]string //post - a map of strings with string as a key
	mu    sync.RWMutex      //mu - mutex lock
}

// Make a new Store
// ==========================================================

func NewStore() *Store {
	return &Store{
		posts: make(map[string]string),
	}
}

// Get data from the Store
// ==========================================================
func (s *Store) Get(key string) string {

	// Put a lock on the store whilst reading from it
	s.mu.RLock()
	// Unlock the store after the get is done - defer until everything is finished
	defer s.mu.RUnlock()
	// Lookup the post in the map and set to a variable 'p'
	p := s.posts[key]
	// Return the post data
	return p

}

// Set data on the Store
// ==========================================================
func (s *Store) Set(key, post string) bool {

	s.mu.RLock()
	defer s.mu.RUnlock()

	// Check if the key already exists
	_, exists := s.posts[key]

	if exists {
		return false
	}
	s.posts[key] = post
	return true
}

func main() {

	s := NewStore()

	if s.Set("Kyle", "Pritchard") {
		fmt.Println("Saved to the store")
	}

	get := s.Get("Kyle")
	fmt.Println("The value returned is:", get)

	u3, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(u3.String())

	fmt.Println(random.String(8))
}
