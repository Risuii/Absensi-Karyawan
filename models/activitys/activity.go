package activitys

import "time"

type Activity struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userID"`
	Description string    `json:"deskripsi" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}
