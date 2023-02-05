package users

type EmployeeLogin struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"email"`
}
