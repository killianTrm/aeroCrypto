# Blockchain - Médicare

	Authors : Christophe VIGUIER, Killian TERROM, Clément QUENNESSON, Audric PANDELE, Pierre BRETHES
	Version : 1.0

Ce document spécifie le projet "blockchain PièceAvion", appelé "AeroCrypto". 

## Problème


> Aucun suivi des pièces aéronautique sur les chaines de production

> Acteurs et fournisseurs peu important, beaucoup d'import export de pièces.

> 300% d’augmentation du trafic de Aeronautique entre 2007 et 2010

> Le marché de la recupération de pièce aéronautique est important, la fabrication de celles-ci demandant un fort cout energétique


## Pourquoi une blockchain ?
> Vérification de l'authenticité de la pièce 

> Traçabilité des différents fournisseurs

> Gain de temps

> Registre immuable et durable dans le temps

> Suivi de la non-récupération de pièces

## User stories 

### Pour le client 

> Permet de vérifier que la localisation actuelle de la pièce ainsi que le lieu d'achat

### Pour l'état 

> Permet de tracer toutes les pièces existantes dans sa compagnie voir dans le monde (import / export de pièces)

## Technologies utilisées 

Architecture basée sur le projet : [https://github.com/hyperledger/education/tree/master/LFS171x/fabric-material/tuna-app](https://github.com/hyperledger/education/tree/master/LFS171x/fabric-material/tuna-app)

## Modélisation d'un médicament 

| Nom du champ | Description  |  Format  |
|--|--|--|
| refPiece| Référence unique à chaque produit (présent dans un datamatrix) | Chaîne de caractères | 
| nomPiece | Nom du produit | Chaîne de caractères |
| nomFabricant | Nom du fabriquant du produit | Chaîne de caractères |
| numeroLot | Numéro de lot du produit | Chaîne de caractères |
| dateExp | Date d'expiration | Date |
| localisationVille | Ville actuelle du produit | Chaîne de caractères |
| locationEtablissement | Etablissement de réception du produit | Chaîne de caractères |
| dateCreation | Date de création du médicament | Date |
| dateReception | Date de réception associée au lieu | Date |
| venteClient | Booléen permettant de savoir si le produit a été vendu | Booléen |


## CRUD 

### CREATE 

La partie CREATE du client web permet l’ajout d’une pièce dans la blockchain. Les différents champs sur la page CREATE sont tous les champs présentés ci-dessus (Partie Modélisation d’une pièce aéronautique).

### READ 

La partie READ du client web permet de visualiser les différentes pièces présentes dans la blockchain. Tous les champs de la blockchain sont visibles dans la partie READ.

### UPDATE 

La partie UPDATE permet de mettre à jour la position actuelle de la pièce, ainsi que sa date de vente à partir de la référence du produit.

| Nom du champ | Description  |  Format  |
|--|--|--|
| refPice | Référence unique à chaque pièce (présent dans un datamatrix) | Chaîne de caractères | 
| localisationVille | Ville actuelle de la pièce | Chaîne de caractères |
| locationEtablissement | Etablissement de réception de la pièce | Chaîne de caractères |
| dateReception | Date de réception associée au lieu | Date |
| venteClient | Booléen permettant de savoir si la pièce a été vendu | Booléen |

### DELETE 

Pas de possibilité de supprimer des pièces. 


