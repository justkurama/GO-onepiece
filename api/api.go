package api

type mugiwaras struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	NickName    string `json:"nickname"`
	Bounty      int    `json:"bounty"`
	DateOfBirth string `json:"date_of_birth"`
	Origin      string `json:"origin"`
	Occupations string `json:"occupations"`
}

var Players = []Player{}
