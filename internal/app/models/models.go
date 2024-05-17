package models

type Character struct {
	ID                uint `gorm:"primaryKey"`
	Name              string
	NickName          string
	Origin            Origin `gorm:"foreignKey:OriginID"`
	OriginID          uint
	Race              Race `gorm:"foreignKey:RaceID"`
	RaceID            uint
	Organization      Organization `gorm:"foreignKey:OrganizationID"`
	OrganizationID    uint
	SubOrganization   SubOrganization `gorm:"foreignKey:SubOrganizationID"`
	SubOrganizationID uint
}

type Origin struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Characters []Character `gorm:"foreignKey:OriginID"`
}

type Race struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Characters []Character `gorm:"foreignKey:RaceID"`
}

type Organization struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	Characters       []Character       `gorm:"foreignKey:OrganizationID"`
	SubOrganizations []SubOrganization `gorm:"foreignKey:ParentID"`
}
type SubOrganization struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	ParentID uint
	Parent   Organization `gorm:"foreignKey:ParentID;references:ID"`
}
