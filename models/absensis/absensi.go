package absensis

import "time"

type Absensi struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"userID"`
	Name     string    `json:"name"`
	Checkin  time.Time `json:"checkin"`
	Checkout time.Time `json:"checkout"`
}
