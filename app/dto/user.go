package dto

type UserDto struct {
	ID        int    `json:"id"`
	CompanyID int    `json:"company_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserDto struct {
	CompanyID *int    `json:"company_id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
}
