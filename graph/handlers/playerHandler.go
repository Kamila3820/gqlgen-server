package handlers

import "gqlgen-server/graph/model"

var players = []*model.Player{
	{
		ID:    "PLAYER-001",
		Name:  "Fang",
		Level: 1,
		Class: &model.AllPlayerClass[1],
		Item: []*model.Item{
			{
				ID:   "ITEM-001",
				Name: "Scepter",
			},
			{
				ID:   "ITEM-002",
				Name: "Potion",
			},
		},
	},
	{
		ID:    "PLAYER-001",
		Name:  "Five",
		Level: 5,
		Class: &model.AllPlayerClass[0],
		Item: []*model.Item{
			{
				ID:   "ITEM-003",
				Name: "Sword",
			},
		},
	},
}

func GetPlayers() []*model.Player {
	return players
}
