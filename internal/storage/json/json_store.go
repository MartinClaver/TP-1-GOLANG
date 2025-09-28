// internal/storage/json/json_store.go
package json

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type JSONStore struct {
	path   string
	mu     sync.RWMutex
	data   map[uint]*app.Contact
	nextID uint
}

type fileData struct {
	NextID  uint          `json:"next_id"`
	Records []app.Contact `json:"records"`
}

func New(path string) (storage.Storage, error) {
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		_ = os.MkdirAll(dir, 0o755)
	}
	js := &JSONStore{
		path: path,
		data: make(map[uint]*app.Contact),
	}
	_ = js.load() // si pas de fichier => démarrage vide
	return js, nil
}

func (s *JSONStore) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	f, err := os.Open(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s.nextID = 1
			return nil
		}
		return err
	}
	defer f.Close()

	var fd fileData
	if err := json.NewDecoder(f).Decode(&fd); err != nil {
		return err
	}
	s.nextID = fd.NextID
	if s.nextID == 0 {
		s.nextID = 1
	}
	for i := range fd.Records {
		c := fd.Records[i]
		cp := c
		s.data[c.ID] = &cp
	}
	return nil
}

func (s *JSONStore) persist() error {
	var fd fileData
	fd.NextID = s.nextID
	for _, c := range s.data {
		fd.Records = append(fd.Records, *c)
	}
	sort.Slice(fd.Records, func(i, j int) bool { return fd.Records[i].ID < fd.Records[j].ID })
	tmp := s.path + ".tmp"

	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(&fd); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}

func (s *JSONStore) Add(c *app.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	c.ID = s.nextID
	s.nextID++
	cp := *c
	s.data[c.ID] = &cp
	return s.persist()
}

func (s *JSONStore) List() ([]app.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var res []app.Contact
	for _, c := range s.data {
		res = append(res, *c)
	}
	sort.Slice(res, func(i, j int) bool { return res[i].ID < res[j].ID })
	return res, nil
}

func (s *JSONStore) Update(id uint, updates map[string]interface{}) (*app.Contact, error) {
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
	if err := s.persist(); err != nil {
		return nil, err
	}
	cp := *c
	return &cp, nil
}

func (s *JSONStore) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; !ok {
		return fmt.Errorf("contact %d introuvable", id)
	}
	delete(s.data, id)
	return s.persist()
}

func (s *JSONStore) GetByID(id uint) (*app.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.data[id]
	if !ok {
		return nil, fmt.Errorf("contact %d introuvable", id)
	}
	cp := *c
	return &cp, nil
}

func (s *JSONStore) GetByEmail(email string) (*app.Contact, error) {
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

func (s *JSONStore) Close() error { return nil }
