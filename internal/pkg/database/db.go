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
}

func addAdmin() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}
	admin := models.User{
		Login:    "admin",
		Password: "admin",
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
		{Name: "Pirate Crews", ParentID: 1},
		{Name: "Marines", ParentID: 2},
		{Name: "Seven Warlords", ParentID: 3},
		{Name: "World Government", ParentID: 4},
		{Name: "Revolutionary Army", ParentID: 5},
	}
	for _, organization := range organizations {
		DB.Create(&organization)
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
