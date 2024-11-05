package main

import (
	"fmt"
	"gorm_crud/migrations"
	"gorm_crud/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize a connection to the SQLite database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Successfully connected to the database")

	// Run the migrations (create or update tables)
	if err := migrations.RunMigrations(db); err != nil {
		fmt.Println("Error running migrations:", err)
		return
	}

	fmt.Println("Migrations applied successfully")

	// Create a new user
	user := models.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Age: 30}
	db.Create(&user)
	fmt.Printf("User created with ID: %d\n", user.ID)

	// Read a user by email
	var fetchedUser models.User
	db.Where("email = ?", "john.doe@example.com").First(&fetchedUser)
	fmt.Printf("User fetched: %s %s, Age: %d\n", fetchedUser.FirstName, fetchedUser.LastName, fetchedUser.Age)

	// Update the user's age
	fetchedUser.Age = 31
	db.Save(&fetchedUser)
	fmt.Printf("User updated: %s %s, Age: %d\n", fetchedUser.FirstName, fetchedUser.LastName, fetchedUser.Age)

	// Delete the user
	db.Delete(&fetchedUser)
	fmt.Println("User deleted")
}
