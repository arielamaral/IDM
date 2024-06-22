package models

import "time"

func init() {
	// Create the table in the database
	// db.AutoMigrate(&User{})
}

type User struct {
	EmployeeID      string    `json:"employee_id"`
	Username        string    `json:"username"`
	Name            string    `json:"name"`
	LastName        string    `json:"last_name"`
	DisplayName     string    `json:"display_name"`
	Email           string    `json:"email"`
	ManagerUsername string    `json:"manager_username"`
	ManagerEmail    string    `json:"manager_email"`
	Position        string    `json:"position"`
	Division        string    `json:"division"`
	Team            string    `json:"team"`
	Subteam         string    `json:"subteam"`
	Status          string    `json:"status"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
}
