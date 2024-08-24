package twittergraphql

import "time"


type UserRepo interface{
	
}
// User represents a system user with relevant details.
type User struct {
	ID        string    // Unique identifier for the user
	Username  string    // Username chosen by the user
	Email     string    // User's email address
	Password  string    // User's password (should be stored securely)
	CreatedAt time.Time // Timestamp of when the user was created
	UpdatedAt time.Time // Timestamp of the last update to the user's information
}
