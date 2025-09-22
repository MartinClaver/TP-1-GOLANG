package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type Contact struct {
	ID    int
	Nom   string
	Email string
}

func main() {
	contactMap := make(map[int]Contact)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello ! Bienvenue dans un Mini-CRM pour gérer tes contacts. Tape un des chiffres suivants pour effectuer une action")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Voir tes contacts")
	fmt.Println("3. Supprimer un contact à l'aide de son ID")
	fmt.Println("4. Mettre à jour un contact à l'aide de son ID")
	fmt.Println("5. Quitter le CRM")
	input, _ := reader.ReadString('\n')
	choix, _ := strconv.Atoi(strings.TrimSpace(input))
	for {
		switch choix {
		case 1:
			AddContact(&contactMap)
		case 2:
			fmt.Println(ShowContactMap(contactMap))
		case 3:
			DeleteContactMap(&contactMap)
		case 4:
			UpdateContactMap(&contactMap)
		case 5:
			break
		}
	}
}
