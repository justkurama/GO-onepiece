package database

import (
	"strconv"

	"github.com/justkurama/GO-onepiece/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	//host = "host.docker.internal"
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "kurama_0723"
	dbname   = "onepiece"
)

var DB = GetDB()

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},

		&models.Origin{},
		&models.Race{},
		&models.Organization{},
		&models.Character{},
		&models.SubOrganization{})
	if err != nil {
		panic("migration failed")
	}
	addInitialData()
}

func addInitialData() {
	addAdmin()
	addOrigins()
	addRaces()
	addOrganization()
	addSubOrganizations()
}

func addAdmin() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}
	admin := models.User{
		Login:    "admin",
		Password: "admin	",
	}
	DB.Create(&admin)
}
func addOrigins() {
	var count int64
	DB.Model(&models.Origin{}).Count(&count)
	if count > 0 {
		return
	}
	origins := []models.Origin{
		{Name: "East Blue"},
		{Name: "North Blue"},
		{Name: "South Blue"},
		{Name: "West Blue"},
		{Name: "Grand Line"},
		{Name: "New World"},
	}
	for _, origin := range origins {
		DB.Create(&origin)
	}
}
func addRaces() {
	var count int64
	DB.Model(&models.Race{}).Count(&count)
	if count > 0 {
		return
	}
	races := []models.Race{
		{Name: "Human"},
		{Name: "Giant"},
		{Name: "Fishman"},
		{Name: "Mink"},
		{Name: "Skypiean"},
		{Name: "Longleg Tribe"},
		{Name: "Longarm Tribe"},
	}
	for _, race := range races {
		DB.Create(&race)
	}
}
func addOrganization() {
	var count int64
	DB.Model(&models.Organization{}).Count(&count)
	if count > 0 {
		return
	}
	organizations := []models.Organization{
		{Name: "Pirate Crews"},
		{Name: "Marines"},
		{Name: "Seven Warlords"},
		{Name: "World Government"},
		{Name: "Revolutionary Army"},
	}
	for _, organization := range organizations {
		DB.Create(&organization)
	}
}
func addSubOrganizations() {
	var count int64
	DB.Model(&models.SubOrganization{}).Count(&count)
	if count > 0 {
		return
	}
	subOrganizations := []models.SubOrganization{
		{Name: "Straw Hat Pirates", ParentID: 1},        // Pirate Crews
		{Name: "Red Hair Pirates", ParentID: 1},         // Pirate Crews
		{Name: "Whitebeard Pirates", ParentID: 1},       // Pirate Crews
		{Name: "Blackbeard Pirates", ParentID: 1},       // Pirate Crews
		{Name: "Big Mom Pirates", ParentID: 1},          // Pirate Crews
		{Name: "Marines HQ", ParentID: 2},               // Marines
		{Name: "G-5", ParentID: 2},                      // Marines
		{Name: "Cipher Pol", ParentID: 2},               // Marines
		{Name: "Warlords Council", ParentID: 3},         // Seven Warlords
		{Name: "Revolutionary Commanders", ParentID: 5}, // Revolutionary Army
	}
	for _, subOrganization := range subOrganizations {
		DB.Create(&subOrganization)
	}
}

func GetDB() *gorm.DB {
	dbURL := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + dbname
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("no connection")
	}
	return db
}
