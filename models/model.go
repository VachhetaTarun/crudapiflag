package model

// Worker represents a worker in the system.
type Worker struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	MongoID string `gorm:"column:mongo_id"` // Store MongoID as a string
	Name    string `json:"name"`
	Role    string `json:"role"`
	SiteID  int    `json:"site_id"`
}

// SiteManagerDetails represents the details of a site manager.
type SiteManagerDetails struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	MongoID string `gorm:"column:mongo_id"` // Store MongoID as a string
	Name    string `json:"name"`
	Email   string `json:"email"`
}
