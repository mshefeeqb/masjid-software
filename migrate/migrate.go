package main

import (
	"fmt"
	"log"

	"github.com/mshefeeqb/masjid-software/initializers"
	"github.com/mshefeeqb/masjid-software/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Member{}, &models.Ward{}, &models.FeePackage{})
	fmt.Println("? Migration complete")
}
