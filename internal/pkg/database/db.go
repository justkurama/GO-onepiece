package database

import (
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Kurama_0723"
	dbname   = "onepiece"
)

var DB = GetDB()

func Migrate() {
	err := DB.AutoMigrate(&models.Origin{}, &models.Race{}, &models.Organization{}, &models.Character{})
	if err != nil {
		panic("migration failed")
		return
	}
	addInitialData()
}

func addInitialData() {
	addOrigins()
	addRaces()
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
func GetDB() *gorm.DB {
	dbURL := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + dbname
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("no connection")
	}
	return db
}
