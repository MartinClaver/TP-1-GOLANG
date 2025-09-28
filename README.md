# TP-4-GOLANG


Un gestionnaire de contacts simple en ligne de commande, écrit en Go.

## Fonctionnalités
- CRUD complet des contacts (add, list, update, delete)
- CLI professionnelle avec Cobra
- Configuration externe avec Viper (`configs/config.yaml`)
- Persistance interchangeable : GORM/SQLite, JSON, mémoire

## Installation

Installer les dépendances :
`go mod tidy`

Générer le binaire :
`go build -o bin/crm .`

## Execution

Ajout d'un contact :
`./bin/crm add -f "Ada" -l "Lovelace" -e "ada@analytical.com" -p "0600000000" -c "Babbage" -n "Pionnière"`

Lister un contact : 
`./bin/crm list`

ou en JSON `./bin/crm list -o json`

Mettre à jour un contact : `./bin/crm update 1 --phone "0612345678" --company "Math Society"`

Supprimer un contact : `./bin/crm delete 1`