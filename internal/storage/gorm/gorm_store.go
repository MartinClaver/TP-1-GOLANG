// internal/storage/gorm/gorm_store.go
package gorm

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/storage"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type GORMStore struct {
	db *gorm.DB
}

func New(path string) (storage.Storage, error) {
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		_ = os.MkdirAll(dir, 0o755)
	}
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ouverture SQLite: %w", err)
	}
	if err := db.AutoMigrate(&app.Contact{}); err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}
	return &GORMStore{db: db}, nil
}

func (s *GORMStore) Add(c *app.Contact) error {
	return s.db.Create(c).Error
}

func (s *GORMStore) List() ([]app.Contact, error) {
	var res []app.Contact
	if err := s.db.Order("id ASC").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (s *GORMStore) Update(id uint, updates map[string]interface{}) (*app.Contact, error) {
	var c app.Contact
	if err := s.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&c).Updates(updates).Error; err != nil {
		return nil, err
	}
	// recharger pour valeurs finales
	if err := s.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *GORMStore) Delete(id uint) error {
	return s.db.Delete(&app.Contact{}, id).Error
}

func (s *GORMStore) GetByID(id uint) (*app.Contact, error) {
	var c app.Contact
	if err := s.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *GORMStore) GetByEmail(email string) (*app.Contact, error) {
	var c app.Contact
	if err := s.db.Where("email = ?", email).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *GORMStore) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
