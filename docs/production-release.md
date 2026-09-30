# Mise à jour de la VM AeroPath existante

Le workflow **CD** se lance explicitement sur `main` (`workflow_dispatch`). Un push GitHub conserve les contrôles CI, mais ne remplace plus immédiatement le site avec un import incomplet. L'image API utilise un tag de commit immuable et contient uniquement le binaire, les migrations et l'interface compilée.

Le script `deployments/aws/deploy-existing-vm.sh` s'exécute en root sur la VM avec deux arguments : le répertoire d'une version exportée par `git archive` sous `/opt/aeropath/releases/`, puis son image API déjà disponible localement. Le fichier `.env` existant et les certificats restent sur la VM.

## Contrôles et bascule

1. Verrou exclusif de déploiement et sauvegarde PostgreSQL vérifiée, dans un répertoire accessible uniquement à root sous `/var/backups/`.
2. Extraction des seuls fichiers publics compilés dans `release/public`, distinct du dépôt source. Vérification des types de fichiers et recherche des secrets d'environnement dans ces fichiers, sans afficher les valeurs.
3. Migrations additives 005–007 et import PPL/LAPL/CPL/ATPL/IR dans une transaction unique. Les tables personnelles sont verrouillées pendant cette opération ; leurs empreintes et la conservation des anciens identifiants de questions sont vérifiées avant validation.
4. Démarrage d'une API candidate sur `127.0.0.1:18080`. Vérification des cours FR/EN, des 15 QCM par leçon, des examens PPL, de corrections invitées et de l'authentification des routes privées.
5. Remplacement des seuls services API et Nginx, puis vérification HTTPS, fichiers publics, SVG et refus des chemins sensibles.

Le fichier `/opt/aeropath/compose.release.yml` désigne la version active. Toute commande manuelle visant ces services doit utiliser **les deux** fichiers Compose : `docker-compose.yml` et `compose.release.yml`. `active-release` contient le chemin de la dernière version vérifiée. Aucun `git clean`, effacement de volumes ou redémarrage global de la stack n'est effectué.

## Retour arrière

Un échec après préparation de la nouvelle configuration restaure automatiquement la configuration précédente de l'application. Les migrations sont additives et les anciens enregistrements ne sont pas supprimés : le code précédent peut continuer à fonctionner. Les contenus importés restent présents si leur transaction avait réussi.

Le répertoire privé de sauvegarde contient `compose.release.yml`, les fichiers de configuration précédents, le dump de la base et les diagnostics. Pour revenir manuellement au code précédent, restaurer ce fichier Compose puis relancer uniquement `api-gateway` et `nginx` avec les deux fichiers Compose. Vérifier ensuite HTTPS et `/health`.

Une restauration intégrale du dump PostgreSQL doit se faire lors d'une maintenance planifiée, après conservation des écritures postérieures à la sauvegarde. Elle n'est pas déclenchée automatiquement : cela pourrait effacer de nouveaux résultats d'utilisateurs.

## Confidentialité

Ne jamais ajouter `.env`, clés, certificats privés, sauvegardes, journaux ou `.local/` à Git. Ne jamais placer le dépôt ou les sauvegardes dans la racine web. Les cookies de session sont `HttpOnly`, `SameSite=Lax` et `Secure` en production. Les tests de déploiement ne nécessitent pas d'exporter un compte réel ou son mot de passe.
