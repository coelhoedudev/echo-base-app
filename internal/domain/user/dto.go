package user

type CreateUserDTO struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  *string `json:"lastName,omitempty" `
	Password  string  `json:"password" validate:"required"`
	Email     string  `json:"email" validate:"required email"`
}

type UpdateUserDTO struct {
	ID string `json:"id" validate:"required"`
	CreateUserDTO
}
