package config

import (
	"fmt"

	"github.com/ihulsbus/cookbook/internal/interfaces"
	"github.com/ihulsbus/cookbook/internal/models"
	"github.com/ihulsbus/cookbook/internal/repositories"
	"github.com/ihulsbus/cookbook/internal/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Configuration    Config
	RecipeRepository interfaces.RecipeRepository
	RecipeService    services.RecipeService
)

func initViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/cookbook/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found! Error was: %v", err)
		} else {
			log.Fatalf("Unknown config error occured. Error is: %v", err)
		}
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		log.Fatalf("Error unmarshalling config file: %v", err)
	}

	viper.WatchConfig()

	log.Infof("Using config file found at: %s", viper.GetViper().ConfigFileUsed())
	if Configuration.Global.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debugln("Enabled DEBUG logging level")
	}

	return err
}

func initDatabase(host string, user string, password string, dbname string, port int, sslmode string, timezone string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)
	log.Debugf(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to the database. Exiting..\n%v\n", err)
	}

	err = db.SetupJoinTable(&models.Recipe{}, "Ingredients", &models.Recipe_Ingredient{})
	if err != nil {
		log.Errorf("Error while creating recipe join tables: %s", err.Error())
	}

	err = db.SetupJoinTable(&models.MealPlan{}, "Recipes", &models.MealPlan_Recipe{})
	if err != nil {
		log.Errorf("Error while creating mealplan join tables: %s", err.Error())
	}

	err = db.AutoMigrate(
		&models.Ingredient{},
		&models.UserData{},
		&models.ShoppingList{},
		&models.Recipe{},
		&models.MealPlan{})

	if err != nil {
		log.Errorf("Error while automigrating database: %s", err.Error())
	}

	return db
}

func init() {
	if err := initViper(); err != nil {
		log.Fatalf("Error reading config: %v", err.Error())
	}
	Configuration.DatabaseClient = initDatabase(Configuration.Database.Host, Configuration.Database.Username, Configuration.Database.Password, Configuration.Database.Database, Configuration.Database.Port, Configuration.Database.SSLMode, Configuration.Database.Timezone)

	// Init repositories
	RecipeRepository = repositories.NewGormRecipeRepository(Configuration.DatabaseClient)

	// Init services
	RecipeService = services.NewRecipeService(RecipeRepository)

}
