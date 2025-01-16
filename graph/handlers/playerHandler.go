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

func CreatePlayer(req *model.NewPlayer) *model.Player {
	new := &model.Player{
		ID:    "PLAYER-003",
		Name:  req.Name,
		Level: 1,
		Class: &req.Class,
		Item: []*model.Item{
			{
				ID:   "ITEM-000",
				Name: "Stick",
			},
		},
	}

	players = append(players, new)

	return new
}
