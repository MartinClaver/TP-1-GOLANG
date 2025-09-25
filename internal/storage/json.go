package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonStore struct {
	filePath string
	contacts map[int]*Contact
}

func (j JsonStore) Add(c *Contact) {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
		return
	}
	err = os.WriteFile(j.filePath, data, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
		return
	}
	fmt.Println("Ajout effectué")
}

func (j JsonStore) GetAll() []Contact {
	var contacts []Contact
	err := json.Unmarshal([]byte(j.filePath), &contacts)
	if err != nil {
		fmt.Println("Erreur lors du décodage :", err)
		return nil
	}
	fmt.Println(contacts)
	return contacts
}

func (j JsonStore) GetById(id int) {
	var contact Contact
	err := json.Unmarshal([]byte(j.filePath), &contact)
	if err != nil {
		fmt.Println("Erreur lors du décodage :", err)
		return
	}
	fmt.Println(contact)
}
func (j JsonStore) Update() {

}

func (j JsonStore) Delete() {}
