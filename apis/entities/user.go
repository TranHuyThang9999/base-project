package entities

type CreateUserRequest struct {
	UserName    string `json:"user_name,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
