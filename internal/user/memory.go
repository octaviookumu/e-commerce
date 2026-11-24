package user

// implementation of storage. interface. in memory impl

import (
	"context"
	"sort"
	"sync"
)

// MemoryRepository: one implementation
type MemoryRepository struct {
	mu    sync.RWMutex // to add concurrency because maps in go are not concurrent-safe
	users map[string]User
}

// will be used as repo := user.NewMemoryRepository()
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{users: make(map[string]User)}
}

// takes Lock because it changes the map.
func (r *MemoryRepository) Create(_ context.Context, u User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[u.ID] = u
	return u, nil
}

// take RLock because they only read it.
func (r *MemoryRepository) ByID(_ context.Context, id string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.users[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}

// take RLock because they only read it.
func (r *MemoryRepository) List(_ context.Context) ([]User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]User, 0, len(r.users))
	for _, u := range r.users {
		out = append(out, u)
	}
	// sort the out slice in place, by each User’s ID, in ascending lexicographic order:
	sort.Slice(out, func(i, j int) bool {
		return out[i].ID < out[j].ID
	})
	return out, nil
}
