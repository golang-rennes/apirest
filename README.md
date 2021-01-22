# Démo API REST

Ce dépôt contient les exemples de code et les [slides](slides.pdf) de la présentation du 15 juin 2017 sur la réalisation d'une **API REST** en Go, donnée par Benoît Masson dans le cadre du [Golang Rennes](https://www.meetup.com/fr-FR/Golang-Rennes/events/240584618/).

[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE.txt)

## Organisation

À chacune des 4 parties correspond un dossier dans `src/`.

Les exemples sont indiqués dans la présentation par une icône de terminal. Pour reproduire ces exemples, vous pouvez repartir des tags suivants :

-   `1-http` : serveur HTTP initialisé, reste à ajouter le(s) handler(s) souhaité(s) ;
-   `2-mux` : routeur `mux` initialisé, route `/users` définie et handler correspondant implémenté, reste à implémenter les autres routes ;
-   `3-beego` : serveur `beego` initialisé et intégralement développé ;
-   `4-restful` : conteneur initialisé, ressources `/users` définie, reste à définir les routes et ajouter les contrôleurs correspondants ;
-   `4-swagger` : serveur `go-restful` complet, génération minimale de `/apidocs.json` et affichage de Swagger-UI sur `docs`, reste à compléter les métadonnées de documentation ;
-   `4-filters` : serveur `go-restful` avec documentation Swagger-UI complète et ajout de filtres au conteneur.

## Exécution

Le code de chaque partie peut être compilé en exécutant à la racine du dépôt `GOPATH=$PWD go build <dossier>` (`<dossier>` parmi `1-http`, `2-mux`, `3-beego`, `4-restful` donc).
Le binaire ainsi créé peut être lancé directement depuis la racine (l'exécution du code Beego nécessite le dossier `conf`, celle du code Restful avec Swagger-UI a besoin du dossier `swagger-ui-2`).

Par défaut, tous les serveurs s'exécutent sur `localhost:8888`, ce qui peut être modifié dans chaque `main.go`.

Pour tester :

```sh
curl -i -X GET  localhost:8888/users && echo
curl -i -X GET  localhost:8888/users/1789 && echo                                          # à partir de 2-mux
curl -i -X POST localhost:8888/users -d '{"firstname":"toto", "lastname":"titi"}' && echo  # à partir de 2-mux
curl -i -X GET  localhost:8888/users?firstname=John && echo                                # à partir de 4-restful
```

## License

MIT License

See [LICENSE](LICENSE.txt) to see the full text.
