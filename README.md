
# Jump API

**Questionnement :**

- retour HTTP 201 vs 204

**Améliorations :**

- middleware d'authentification
- gestion de la configuration via config yaml/var d'env
- récupérer le mot de passe de db depuis un gestionnaire de secret
- système de lock sur les factures pour éviter les doubles paiements
- quelques tests supplémentaires notamment dans endpoint_test.go, il faudrait setup un meilleur environnement pour les tests afin d'avoir une db éphémère et faire des tests en cascade (creation de facture puis validation de celle-ci)
- gérer un versioning via /v1 notamment dans l'url

  

**Principale problème rencontré :**

- gérer les décimales avec un nombre flotant

# Démarrer le projet
  
## Quelques commandes 

**Pour démarrer les tests :** make test
**pour builder le projet :** make builds

  ## Pour le déployer via docker
- buildez l'image docker de la base de données transmise en exemple sous le nom "jump-database"
- faites un docker compose up

l'API est ensuite accessible sur l'adresse http://127.0.0.1:8080/

## Elements de configuration

Comme indiqué dans le point numéro 2 des améliorations, il faudrait mettre en place un système de configuration.
A la place de ça, pour le moment il y a 2 endroits où l'accès à la base de donnée est indiqué :

- cmd/api/setup/setup.go : c'est la configuration utilisée pour docker, d'où l'url indiqué à "db". Elle peut être modifié si le projet est buildé en local.
- internal/transport/endpoint_test.go : on y retrouve cette fois-ci l'url en 127.0.0.1 pour permettre de lancer les tests en dehors d'un container.
  

j'ai volontairement laissé gin en mode debug pour que ce soit plus agréable lors des tests.
