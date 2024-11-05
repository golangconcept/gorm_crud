// models/user.go
package models

// User represents a user in the system
type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:100"`
	LastName  string `gorm:"size:100"`
	Email     string `gorm:"uniqueIndex"`
	Age       int
}
