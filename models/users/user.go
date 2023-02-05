package users

import "time"

type Employee struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Email    string    `json:"email" validate:"email"`
	Checkin  time.Time `json:"checkin"`
	Checkout time.Time `json:"checkout"`
	// Activity  []Activity `json:"activity" foreignkey:"userID"`
	// Absensi   []Absensi  `json:"absen" foreignkey:"userID"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
