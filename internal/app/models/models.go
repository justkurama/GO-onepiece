package models

type Character struct {
	ID             uint
	Name           string
	NickName       string
	Origin         Origin
	OriginID       uint `json:"-"`
	Race           Race
	RaceID         uint `json:"-"`
	Organization   Organization
	OrganizationID uint `json:"-"`
}

type Origin struct {
	ID         uint
	Name       string
	Characters []Character `json:"-"`
}

type Race struct {
	ID         uint
	Name       string
	Characters []Character `json:"-"`
}

type Organization struct {
	ID         uint
	Name       string
	Characters []Character `json:"-"`
}
