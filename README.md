# API-APPLICATION-METEO

Une application web développée en Go permettant de rechercher une commune française et de consulter ses prévisions météorologiques sur sept jours.

## Fonctionnalités

- Recherche d'une commune par son nom.
- Affichage des communes correspondant à la recherche.
- Sélection d'une commune à partir de son code INSEE.
- Récupération des coordonnées géographiques de la commune.
- Affichage des prévisions météorologiques sur sept jours.
- Présentation des températures minimales et maximales.
- Visualisation des variations de température à l'aide de barres graphiques.
- Affichage de la localisation sur une carte OpenStreetMap.
- Gestion des erreurs liées aux recherches et aux appels API.

## Technologies utilisées

- **Go** : langage utilisé pour le serveur et la logique applicative.
- **HTML** : structure des pages web.
- **CSS** : mise en forme et interface utilisateur.
- **API Géo de data.gouv.fr** : recherche des communes françaises et récupération de leurs coordonnées.
- **Open-Meteo** : récupération des prévisions météorologiques.
- **OpenStreetMap** : affichage de la carte géographique.

## Architecture du projet

API-APPLICATION-METEO/
│
├── main.go
│
├── api/
│   ├── client.go
│   ├── communes.api.go
│   └── meteo.api.go
│
├── models/
│
├── services/
│   ├── communes.services.go
│   └── meteo.service.go
│
├── controllers/
│
├── routes/
│   └── main.routes.go
│
├── helpers/
│   └── templates.helpers.go
│
├── templates/
│   ├── accueil.html
│   ├── list-villes.html
│   └── result.html
│
└── assets/
    └── css/
        └── style.css

### Rôle des dossiers

| Dossier        | Rôle                                                 |
| -------------- | ---------------------------------------------------- |
| `api/`         | Gestion des requêtes HTTP vers les API externes.     |
| `models/`      | Structures de données utilisées dans l'application.  |
| `services/`    | Traitement des données et logique métier.            |
| `controllers/` | Gestion des requêtes reçues par le serveur.          |
| `routes/`      | Définition des routes HTTP de l'application.         |
| `helpers/`     | Fonctions utilitaires, notamment pour les templates. |
| `templates/`   | Pages HTML affichées à l'utilisateur.                |
| `assets/`      | Ressources statiques, notamment les fichiers CSS.    |

## Installation et lancement

### Prérequis

- Go installé sur votre machine.
- Une connexion Internet pour accéder aux API externes.

### Lancer le projet

1. Cloner le dépôt : git clone URL_DU_DEPOT

2. Accéder au dossier : cd API-APPLICATION-METEO


3. Télécharger les dépendances : go mod tidy

4. Démarrer le serveur : go run .


5. Ouvrir l'application dans un navigateur : http://localhost:8080

## Fonctionnement de l'application

Le fonctionnement repose sur plusieurs étapes :

1. L'utilisateur saisit le nom d'une commune.
2. L'application interroge l'API Géo afin de récupérer les communes correspondantes.
3. L'utilisateur sélectionne une commune dans la liste.
4. Le serveur récupère ses coordonnées géographiques.
5. Ces coordonnées sont utilisées pour interroger l'API Open-Meteo.
6. Les données météorologiques sont traitées puis transmises au template HTML.
7. L'application affiche les prévisions, les températures et la carte de localisation.

## Tests

Pour vérifier que les packages du projet compilent correctement, exécuter : go test ./...


## APIs utilisées

- API Géo : https://geo.api.gouv.fr/
- Open-Meteo : https://open-meteo.com/
- OpenStreetMap : https://www.openstreetmap.org/
