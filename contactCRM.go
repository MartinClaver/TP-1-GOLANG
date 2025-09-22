package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddContact(contactMap *map[int]Contact) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrer le nom du contact : ")
	nom, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture du nom : %v", err))
	}
	nom = strings.TrimSpace(nom)

	fmt.Print("Entrer l'email du contact : ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture de l'email : %v", err))
	}
	email = strings.TrimSpace(email)

	newID := 1
	for id := range *contactMap {
		if id >= newID {
			newID = id + 1
		}
	}

	(*contactMap)[newID] = Contact{ID: newID, Nom: nom, Email: email}
	fmt.Println("Contact ajouté avec succès !")
}

func ShowContactMap(contactMap map[int]Contact) map[int]Contact {
	return contactMap
}

func DeleteContactMap(contactMap *map[int]Contact) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrer l'ID du contact : ")
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture de l'id : %v", err))
	}
	choix, _ := strconv.Atoi(strings.TrimSpace(id))
	delete(*contactMap, choix)
}

func UpdateContactMap(contactMap *map[int]Contact) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrer l'ID du contact : ")
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture de l'id : %v", err))
	}
	choix, _ := strconv.Atoi(strings.TrimSpace(id))

	fmt.Print("Entrer le nom du contact : ")
	nom, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture du nom : %v", err))
	}
	nom = strings.TrimSpace(nom)

	fmt.Print("Entrer l'email du contact : ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture de l'email : %v", err))
	}
	email = strings.TrimSpace(email)

	(*contactMap)[choix] = Contact{ID: choix, Nom: nom, Email: email}
	fmt.Println("Contact modifié avec succès !")

}
