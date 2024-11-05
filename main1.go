package main

import (
	"fmt"
	"gorm_crud/migrations"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// set the name of the configuration file
	viper.SetConfigFile("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	host := viper.GetString("server.host")

	fmt.Printf("server: %s:%d\n", host, 8080)
	// Initialize a connection to the SQLite database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Successfully connected to the database")

	// Run the migrations (create or update tables)
	// if err := runMigrations(db); err != nil {
	// 	fmt.Println("Error running migrations:", err)
	// 	return
	// }
	runMigrations(db)
	fmt.Println("Migrations applied successfully")

	// Create a new user
	// user := models.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Age: 30}
	// db.Create(&user)
	// fmt.Printf("User created with ID: %d\n", user.ID)

	// // Read a user by email
	// var fetchedUser models.User
	// db.Where("email = ?", "john.doe@example.com").First(&fetchedUser)
	// fmt.Printf("User fetched: %s %s, Age: %d\n", fetchedUser.FirstName, fetchedUser.LastName, fetchedUser.Age)

	// // Update the user's age
	// fetchedUser.Age = 31
	// db.Save(&fetchedUser)
	// fmt.Printf("User updated: %s %s, Age: %d\n", fetchedUser.FirstName, fetchedUser.LastName, fetchedUser.Age)

	// // Delete the user
	// db.Delete(&fetchedUser)
	// fmt.Println("User deleted")
}

func runMigrations(db *gorm.DB) {
	// Run migrations from the migrations package
	if err := migrations.RunMigrations(db); err != nil {
		fmt.Println("Error applying user migration:", err)
		return
	}

	if err := migrations.MigrateCreateProducts(db); err != nil {
		fmt.Println("Error applying product migration:", err)
		return
	}
}
