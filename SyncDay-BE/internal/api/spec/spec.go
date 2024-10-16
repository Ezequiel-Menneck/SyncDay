package spec

import "math/big"

type CreateUserRequest struct {
	Name       string  `json:"name" validate:"min=3,max=20"`
	LastName   string  `json:"last_name" validate:"min=3,max=30"`
	Email      string  `json:"email" validate:"min=5,email"`
	Cellphone  string  `json:"cellphone" validate:"min=1,max=13"`
	BaseSalary big.Int `json:"base_salary"`
}

type UpdateUserBaseSalary struct {
	BaseSalary big.Int `json:"base_salary"`
}
