package models

type Character struct {
	ID                uint
	Name              string
	NickName          string
	Origin            Origin
	OriginID          uint
	Race              Race
	RaceID            uint
	Organization      Organization
	OrganizationID    uint
	SubOrganization   SubOrganization
	SubOrganizationID uint
}

type Origin struct {
	ID         uint
	Name       string
	Characters []Character
}

type Race struct {
	ID         uint
	Name       string
	Characters []Character
}

type Organization struct {
	ID               uint
	Name             string
	Characters       []Character
	SubOrganizations []SubOrganization
	ParentID         uint
}
type SubOrganization struct {
	ID             uint
	Name           string
	OrganizationID uint
	Organization   Organization `gorm:"foreignKey:OrganizationID"` // Parent organization
}
