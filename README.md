# Jump API

**Questionnement :**
- retour HTTP 201 vs 204

**Améliorations :**
- middleware d'authentification
- gestion de la configuration via config yaml/var d'env
- récupérer le mot de passe de db depuis un gestionnaire de secret
- système de lock sur les factures pour éviter les doubles paiements
- quelques tests supplémentaires notamment dans endpoint_test.go, il faudrait setup un meilleur environnement pour les tests afin d'avoir une db ephemere et faire des tests en cascade (creation de facture puis validation de celle-ci)
- gérer un versioning via /v1 notamment dans l'url

**Principale problème rencontré :**
- gérer les décimales avec un nombre flotant

**Pour démarrer les tests :** make test

**pour builder le projet  :** make builds

j'ai volontairement laissé gin en mode debug pour que ce soit plus agréable lors des tests.