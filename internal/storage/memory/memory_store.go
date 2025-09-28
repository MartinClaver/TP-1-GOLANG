// internal/storage/memory/memory_store.go
package memory

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/storage"
	"fmt"
	"sort"
	"sync"
)

type MemoryStore struct {
	mu     sync.RWMutex
	data   map[uint]*app.Contact
	nextID uint
}

func New() storage.Storage {
	return &MemoryStore{
		data:   make(map[uint]*app.Contact),
		nextID: 1,
	}
}

func (s *MemoryStore) Add(c *app.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	c.ID = s.nextID
	s.nextID++
	cp := *c
	s.data[c.ID] = &cp
	return nil
}

func (s *MemoryStore) List() ([]app.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var res []app.Contact
	for _, c := range s.data {
		res = append(res, *c)
	}
	sort.Slice(res, func(i, j int) bool { return res[i].ID < res[j].ID })
	return res, nil
}

func (s *MemoryStore) Update(id uint, updates map[string]interface{}) (*app.Contact, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	c, ok := s.data[id]
	if !ok {
		return nil, fmt.Errorf("contact %d introuvable", id)
	}
	if v, ok := updates["FirstName"]; ok {
		c.FirstName = v.(string)
	}
	if v, ok := updates["LastName"]; ok {
		c.LastName = v.(string)
	}
	if v, ok := updates["Email"]; ok {
		c.Email = v.(string)
	}
	if v, ok := updates["Phone"]; ok {
		c.Phone = v.(string)
	}
	if v, ok := updates["Company"]; ok {
		c.Company = v.(string)
	}
	if v, ok := updates["Notes"]; ok {
		c.Notes = v.(string)
	}
	cp := *c
	return &cp, nil
}

func (s *MemoryStore) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; !ok {
		return fmt.Errorf("contact %d introuvable", id)
	}
	delete(s.data, id)
	return nil
}

func (s *MemoryStore) GetByID(id uint) (*app.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.data[id]
	if !ok {
		return nil, fmt.Errorf("contact %d introuvable", id)
	}
	cp := *c
	return &cp, nil
}

func (s *MemoryStore) GetByEmail(email string) (*app.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, c := range s.data {
		if c.Email == email {
			cp := *c
			return &cp, nil
		}
	}
	return nil, fmt.Errorf("contact avec email %s introuvable", email)
}

func (s *MemoryStore) Close() error { return nil }
