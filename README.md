# JUMP BACKEND INTERVIEW

## Introduction
Dans le cadre d'un entretien technique voici un serveur Go capable de communiquer avec une base de données PostgreSQL.
## Librairies
Pour réaliser ce serveur j'ai utiliser les librairies gin, pq et yaml.<br>
La librairies [gin](https://gin-gonic.com) est un framework web HTTP permettant de creer des API simplement et rapidement. Il offre de bonnes performances et de nombreuses fonctionalités comme l'ajout de middleware ou la possibilité de formatter des structures de données en JSON directement dans les réponses des requêtes.
<br>
La librairie [pq](https://github.com/lib/pq) fournit elle les drivers permettant a la lib sql fournit par go de pouvoir fonctionner avec postgreSQL.
<br>
La librairie [yaml](https://pkg.go.dev/gopkg.in/yaml.v2@v2.4.0) permet de transformer une "String" contenant des informations sous la form yaml en structure de données.

## Lancement
Afin de pouvoir lancer le programme, il est nécessaire d'avoir [Golang](https://go.dev) et [Docker](https://www.docker.com) d'installés.<br>
Il faut aussi que le docker fourni par Jump contenant la base de données PostgreSQL soit lancé.<br>
Pour installer les dépendances du projet, utilisez la commande:

```bash
go install
```

Pour lancer le serveur, utilisez la commande:

```bash
go run .
# or
go build . && ./jump-backend-interview
```

## Configuration
Vous pouvez configurer dans le fichier [config.yml](config.yml) les valeurs suivantes:

```YML
# Server
server:
  ip: "0.0.0.0" # Ip sur lequel le serveur va écouter
  port: 8080 # Port sur lequel le serveur va écouter
# Database credentials
database:
  user: "jump" # Nom de l'utilisateur PostgreSQL
  pass: "password" # Mot de passe de l'utilisateur PostgreSQL
  ip: "localhost" # Ip ou se trouve la base de données PostgreSQL
  port: 5432 # Port de la base de données PostgreSQL
  name: "jump" # Noms de la base de données
```

## Routes
Les routes accessibles sont:

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
- Un code 204 si la requête s'est effectuée avec succès
- Un code 500 si un probleme sur le serveur est survenu
- Un code 404 si l'id de l'utilisateur ne correspond à aucun utilisateur de la base de données
- un code 400 si l'un des arguments du json est manquant ou érroné


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
- Un code 204 si la requête s'est effectuée avec succès
- Un code 422 si la facture est déjà payée
- Un code 400 si le montant de la transaction ne correspond pas au montant de la facture
- Un code 404 si l'id de la facture ne correspond a aucune facture de la base de données
- Un code 500 si un problème sur le serveur est survenu
- un code 501 si l'un des arguments du json est manquant ou érroné