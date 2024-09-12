package data

var allPlayers = []Player{
	{Id: 1, PlayerName: "Peter Forsberg", JerseyNumber: 21, TeamName: "Colorado", Born: 1973},
	{Id: 2, PlayerName: "Mats Sundin", JerseyNumber: 13, TeamName: "Toronto", Born: 1971},
	{Id: 3, PlayerName: "Niklas Lidström", JerseyNumber: 5, TeamName: "Detroit", Born: 1970},
}

func FindAll() []Player {
	return allPlayers
}
