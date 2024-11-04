# Scribble.io Clone

Ce projet est une réplique du jeu **Scribble.io**, un jeu en ligne multijoueur où les joueurs dessinent et devinent des mots. Ce projet est réalisé en groupe avec **Go** pour le backend et **React** pour le frontend.

## Objectif du projet

L'objectif est de reproduire l'essentiel des fonctionnalités du jeu Scribble.io :
- Un joueur dessine un mot tandis que les autres joueurs devinent ce que c’est.
- Les points sont attribués en fonction de la rapidité et de l'exactitude des réponses.
- Le jeu se joue en plusieurs rounds, chaque joueur ayant l'occasion de dessiner à tour de rôle.

## Fonctionnalités principales

### Backend (Go)
- **Gestion des sessions de jeu** : Créer et rejoindre des salles de jeu.
- **Logique de jeu** : Gérer les rounds, les tours de dessin et le calcul des points.
- **Sockets pour la communication en temps réel** : Utilisation de WebSockets pour des échanges rapides et bidirectionnels.
- **Authentification des utilisateurs** : Gestion des sessions et des identifiants pour que chaque utilisateur ait un pseudonyme unique.
- **Gestion des scores** : Calcul et suivi des scores par joueur.

### Frontend (React)
- **Interface de jeu interactive** : Zone de dessin, champ de chat pour deviner les mots et affichage des scores.
- **Zone de dessin en temps réel** : Utilisation de canvas pour permettre au joueur de dessiner et aux autres de voir en direct.
- **Chat en temps réel** : Les joueurs peuvent discuter et envoyer leurs suppositions.
- **Affichage dynamique des scores et des rounds** : Visualisation des scores de chaque joueur et des informations sur les rounds en cours.

## Technologies Utilisées

### Backend
- **Go** : Langage principal pour le serveur backend.
- **WebSockets** : Pour la communication en temps réel entre le serveur et les clients.
- **Gin** : Framework Go pour faciliter la gestion des requêtes HTTP et des WebSockets.
- **Redis** : Stockage en mémoire pour gérer les sessions de jeu et les données en temps réel.

### Frontend
- **React** : Framework JavaScript pour construire une interface utilisateur interactive.
- **Canvas API** : Pour dessiner sur le tableau de jeu.
- **Socket.IO** : Utilisé pour les WebSockets côté client, permettant la communication en temps réel avec le serveur.
- **CSS Modules** : Pour un style modulaire et maintenable.

## Installation et Exécution

### Prérequis

- **Go** (version 1.16 ou supérieure)
- **Node.js** (version 14 ou supérieure) avec **npm**
- **Redis** pour la gestion des sessions

  
1. **Cloner le dépôt**  
 ```bash
   git clone git@github.com:NaheH/Scribble.git
```

2. **Lancer le backend**  
   Dans le répertoire `server`, lancez le serveur Go :
   
 ```bash
   go run main.go
```
3. **Lancer le front**  
   Dans le répertoire `client`, lancez le serveur Go :
   
 ```bash
  npm run start
```

4. **Vous pouvez jouer!**

### Défis et Solutions

Synchronisation en temps réel : Utiliser WebSockets pour maintenir une communication fluide entre les clients et le serveur.
Gestion des sessions de jeu : Redis a été utilisé pour suivre les sessions en temps réel de manière rapide et efficace.
Zone de dessin en temps réel : Utilisation de la Canvas API en combinaison avec les WebSockets pour actualiser les dessins en direct.

### Améliorations Futures
Ajout d'un mode spectateur pour permettre aux utilisateurs de regarder une partie sans y participer.
Optimisation des performances pour gérer un plus grand nombre de joueurs par salle.
Animations et effets visuels pour améliorer l’expérience utilisateur.


### Projet réalisé par l'équipe :

- Valentin Neff
- Elias Ouissi
- Nahé Hutin 
