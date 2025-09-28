package storage

import (
	"awesomeProject/internal/app"
	"fmt"
)

// Contact est notre structure de données centrale
type Contact struct {
	ID    int
	Name  string
	Email string
}

// Storer est un CONTRAT de stockage
// Il définit un ensemble de comportements (méthodes) que tout type
// de stockage doit respecter. On ne se soucie par du comment c'est fait
// (en mémoire, fichier, BDD...) seulement de ce qui peut être fait

type Storage interface {
	Add(contact *app.Contact) error
	List() ([]app.Contact, error)
	Update(id uint, updates map[string]interface{}) (*app.Contact, error)
	Delete(id uint) error
	GetByID(id uint) (*app.Contact, error)
	GetByEmail(email string) (*app.Contact, error)
	Close() error
}

var ErrContactNotFound = func(id int) error { return fmt.Errorf("Contact ave l'ID %d non trouvé", id) }
