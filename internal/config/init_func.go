package config

import (
	"fmt"

	"github.com/ihulsbus/cookbook/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDatabase(host string, user string, password string, dbname string, port int, sslmode string, timezone string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to the database. Exiting..\n%v\n", err)
	}

	err = db.SetupJoinTable(&models.Recipe{}, "Ingredients", &models.Recipe_Ingredient{})
	if err != nil {
		log.Errorf("Error while creating recipe join tables: %s", err.Error())
	}

	err = db.SetupJoinTable(&models.Recipe{}, "Sections", &models.Recipe_Ingredient{})
	if err != nil {
		log.Errorf("Error while creating recipe join tables: %s", err.Error())
	}

	// err = db.SetupJoinTable(&models.MealPlan{}, "Recipes", &models.MealPlan_Recipe{})
	// if err != nil {
	// 	log.Errorf("Error while creating mealplan join tables: %s", err.Error())
	// }

	err = db.AutoMigrate(
		&models.Ingredient{},
		&models.Recipe{},
		&models.Section{},
		// &models.ShoppingList{},
		// &models.MealPlan{},
		// &models.UserData{},
	)

	if err != nil {
		log.Errorf("Error while automigrating database: %s", err.Error())
	}

	return db
}
