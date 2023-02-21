# JUMP BACKEND INTERVIEW

## Introduction
Dans le cadre d'un entretient technique voici un server Go capable de communiquer avec une base de donnée PostgreSQL.

## Lancement
Afin de pouvoir lancer le programme il est nécéssaire d'avoir [Golang](https://go.dev) et [Docker](https://www.docker.com) d'installés.<br>
Il faut aussi que le docker fourni par Jump contenant la base de données PostgreSQL soit lancé.<br>
Pour installer les dépendances du projet, utilisez la commande:

```bash
go install
```

Pour lancer le server utilisez la commande:

```bash
go run .
# or
go build . && ./jump-backend-interview
```

## Configuration
Vous pouvez configurer dans le fichier [config.yml](config.yml) les valeurs suivante:

```YML
# Server
server:
  ip: "0.0.0.0" # Ip sur lequel le server va écouter
  port: 8080 # Port sur lequel le server va écouter
# Database credentials
database:
  user: "jump" # Nom de l'utilisateur PostgreSQL
  pass: "password" # Mot de passe de l'utilisateur PostgreSQL
  ip: "localhost" # Ip ou se trouve la base de donnée PostgreSQL
  port: 5432 # Port de la base de donnée PostgreSQL
  name: "jump" # Noms de la base de donnée
```

## Routes
Les routes accessible sont:

### GET /users
la routes GET /users n'attend pas d'argument et renvoie en cas de réussite un json au format suivant avec le code 200:
```json
[
    {
        "user_id": 2,
        "first_name": "Kevin",
        "last_name": "Findus",
        "balance": 492.97
    },
    {
        "user_id": 3,
        "first_name": "Lynne",
        "last_name": "Gwafranca",
        "balance": 825.4
    },
]
```
Si une erreur se produit, le code 500 est renvoyé. 
### POST /invoice
La route POST /invoice attend en entrée un JSON au format:
```json
{
    "user_id": 21,
    "amount": 113.45,
    "label": "Work for April"
}
```
Et renvoie:
- Un code 204 si la requetes c'est effectuée avec succés
- Un code 500 si un probleme sur le server est survenue
- Un code 404 si l'id de l'utilisateur ne correspond a aucun utilisateur de la base de donnée
- un code 400 si l'un des argument du json est manquant ou érronée

### POST /transaction
La route POST /transaction attend en entrée un JSON au format:
```json
{
   "invoice_id": 42,
   "amount": 956.32,
   "reference": "JMPINV200220117"
}
```

Et renvoie:
- Un code 204 si la requetes c'est effectuée avec succés
- Un code 422 si la facture est déjà payée
- Un code 400 si le montant de la transaction ne correspond pas au montant de la facture
- Un code 404 si l'id de la facture ne correspond a aucune facture de la base de donnée
- Un code 500 si un probleme sur le server est survenue
- un code 501 si l'un des argument du json est manquant ou érronée