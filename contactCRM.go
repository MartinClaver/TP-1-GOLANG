package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
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

	contact, err := NewContact(nom, email)
	if err != nil {
		fmt.Println("Erreur lors de la création du contact :", err)
		return
	}
	(*contactMap)[newID] = contact
	fmt.Println("Contact ajouté avec succès ! \n")
}

func ShowContactMap(contactMap map[int]Contact) {
	for id, contact := range contactMap {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", id, contact.Nom, contact.Email)
	}
}

func DeleteContactMap(contactMap *map[int]Contact) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrer l'ID du contact : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture de l'id : %v", err))
	}
	id, _ := strconv.Atoi(strings.TrimSpace(input))

	if _, ok := (*contactMap)[id]; ok {
		delete(*contactMap, id)
		fmt.Println("Contact supprimé avec succès ! \n")
	} else {
		fmt.Println("L'ID n'existe pas ! \n")
	}
}

func UpdateContactMap(contactMap *map[int]Contact) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrer l'ID du contact : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("Erreur de lecture de l'id : %v", err))
	}
	id, _ := strconv.Atoi(strings.TrimSpace(input))

	if _, ok := (*contactMap)[id]; ok {
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

		(*contactMap)[id] = Contact{Nom: nom, Email: email}
		fmt.Println("Contact modifié avec succès ! \n")
	} else {
		fmt.Println("L'ID n'existe pas ! \n")
	}
}

func NewContact(nom, email string) (Contact, error) {
	nom = strings.TrimSpace(nom)
	email = strings.TrimSpace(email)

	if nom == "" {
		return Contact{}, errors.New("le nom ne peut pas être vide")
	}

	// Validation simple de l'email (format basique)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return Contact{}, errors.New("email invalide")
	}

	return Contact{Nom: nom, Email: email}, nil
}
