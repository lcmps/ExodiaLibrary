
# ExodiaLibrary

ExodiaLibrary is a WIP core RESTful API housing info on Yu-Gi-Oh! cards pulled from the YGOPROdeck API.
Data is stored on a Postgres container and can be viewed in Brazillian Portuguese, English, and french. API is made in pure Go :)

### Setup
First, you will need to create the containers to import the actual data from YGOPROdeck API before running the migrate command.
This can be done by either using the `init.sh` script or `docker compose -f docker-compose.yaml up`
followed by the actual import with:
`go run main.go migrate`
It will take a few minutes since there are at least 10.000 cards as of early 2022.

### Running
Now to serve the actual API simply type in a terminal
`go run main.go web`
the default address is `localhost:9001`

### Endpoints
There's currently 3 endpoints
* `/cards`
* `/random`
* `/card-img/`

### `/cards`
The cards endpoint returns a list of cards based on the **optional** parameters that were present in the query, these being:

* Limit ¹
* Offset ¹
* Archetype ²
* Race ²
* Atk ¹ 
* Def ¹
* Level ¹
* Type ²

*¹ Numerical fields*
*² Text fields that are using LIKE when queried on the database*
##### Curl example
`curl --request GET \
--url 'http://localhost:9001/card?limit=1&offset=0&archetype=frog&race=aqua&atk=1900&def=0&level=5&type=Eff'`

##### Response
```
{
	"total": 1,
	"cards": [
		{
			"id": 84451804,
			"name": "Des Frog",
			"name_pt": "Sapo Des",
			"name_fr": "Grenouille Des",
			"type": "Effect Monster",
			"description": "When this card is Tribute Summoned, you can Special Summon \"Des Frog\"(s) from your hand or Deck up to the number of \"T.A.D.P.O.L.E.\"(s) in your Graveyard.",
			"description_pt": "Quando este card é Invocado, por Invocação-Tributo, com sucesso, você pode Invocar, por Invocação-Especial, \"Sapo Des\" da sua mão ou Deck num numéro igual ao de \"T.A.D.P.O.L.E.\"(s) no seu Cemitério.\n",
			"description_fr": "Lorsque cette carte est Invoquée avec succès par Invocation Sacrifice, vous pouvez faire l'Invocation Spéciale d'un nombre de \"Grenouille Des\" de votre main ou votore Deck allant jusqu'au nombre de \"T.A.D.P.O.L.E.\" dans votre Cimetière.\n",
			"image": [
				84451804
			],
			"attribute": "WATER",
			"race": "Aqua",
			"archetype": "Frog",
			"price": "0.98",
			"atk": 1900,
			"def": 0,
			"level": 5
		}
	]
}
```
### `/random`

The `/random` endpoint simply returns random cards from the database based on the limit set on the query, if a limit is not determined a single card will be returned

##### Curl example

`curl --request GET \
  --url 'http://localhost:9001/random?limit=5'`

##### Response
```
[
	{
		"id": 26674724,
		"name": "Nekroz of Brionac",
		"name_pt": "Necroz de Brionac",
		"name_fr": "Nékroz de Brionac",
		"type": "Ritual Effect Monster",
		"description": "You can Ritual Summon this card with any \"Nekroz\" Ritual Spell. Must be Ritual Summoned, without using \"Nekroz of Brionac\". You can only use each of these effects of \"Nekroz of Brionac\" once per turn.\r\n● You can discard this card; add 1 \"Nekroz\" monster from your Deck to your hand, except \"Nekroz of Brionac\".\r\n● You can target up to 2 face-up monsters on the field that were Special Summoned from the Extra Deck; shuffle them into the Deck.",
		"description_pt": "Você pode Invocar este card por Invocação-Ritual com qualquer Magia de Ritual \"Necroz\". Deve ser Invocado por Invocação-Ritual sem usar \"Necroz de Brionac\". Você só pode usar cada um desses efeitos de \"Necroz de Brionac\" uma vez por turno.\r\n● Você pode descartar este card; adicione 1 monstro \"Necroz\" do seu Deck à sua mão, exceto \"Necroz de Brionac\".\r\n● Você pode escolher até 2 monstros com a face para cima no campo que foram Invocados por Invocação-Especial do Deck Adicional; embaralhe-os no Deck.\n",
		"description_fr": "Vous pouvez Invoquer Rituellement cette carte avec une Magie Rituelle \"Nékroz\". Uniquement Invocable Rituellement, sans utiliser \"Nékroz de Brionac\". Vous ne pouvez utiliser chacun de ces effets de \"Nékroz de Brionac\" qu'une fois par tour.\r\n● Vous pouvez défausser cette carte ; ajoutez 1 monstre \"Nékroz\" (\"Nékroz de Brionac\" exclu) depuis votre Deck à votre main.\r\n● Vous pouvez cibler max. 2 monstres face recto sur le Terrain qui ont été Invoqués Spécialement depuis l'Extra Deck ; mélangez-les dans le Deck.\n",
		"image": [
			26674724
		],
		"attribute": "WATER",
		"race": "Warrior",
		"archetype": "Nekroz",
		"price": "3.37",
		"atk": 2300,
		"def": 1400,
		"level": 6
	}
]
```

### `/card-img/`
Quite simply, it's the actual card in `.jpg` format, you have to provide ID+.jpg, no queries needed.

##### Curl example
`curl --request GET \
  --url http://localhost:9001/card-img/84451804.jpg`

![Des Frog card!](/pages/card-img/84451804.jpg "Des Frog")