package api

type mugiwaras struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	NickName    string `json:"nickname"`
	Bounty      int    `json:"bounty"`
	DevilFruit  string `json:"devil_fruit"`
	Origin      string `json:"origin"`
	Occupations string `json:"occupations"`
}

var Mugiwaras = []mugiwaras{
	{1, "Monkey D. Luffy", "Straw Hat", 3000000000, "Gomu Gomu no Mi", "East Blue", "Captain"},
	{2, "Roronoa Zoro", "Pirate Hunter", 1111000000, "None", "East Blue", "Swordsman"},
	{3, "Nami", "Cat Burglar", 366000000, "None", "East Blue", "Navigator"},
	{4, "Usopp", "God", 500000000, "None", "East Blue", "Sniper"},
	{5, "Sanji", "Black Leg", 1032000000, "None", "North Blue", "Cook"},
	{6, "Tony Tony Chopper", "Cotton Candy Lover", 1000, "Hito Hito no Mi", "Grand Line", "Doctor"},
	{7, "Nico Robin", "Devil Child", 930000000, "Hana Hana no Mi", "West Blue", "Archaeologist"},
	{8, "Franky", "Iron Man", 394000000, "None", "South Blue", "Shipwright"},
	{9, "Brook", "Soul-King", 383000000, "Yomi Yomi no Mi", "West Blue", "Musician"},
	{10, "Jinbe", "Knight of the Sea", 1100000000, "None", "Grand Line", "Helmsman"},
}
