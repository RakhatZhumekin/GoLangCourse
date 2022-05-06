package dto

// type UserCreateDTO struct {
// 	ID       uint64 `json:"id" form:"id" binding:"required"`
// 	Name     string `json:"name" form:"name" binding:"required, min:1"`
// 	Email    string `json:"email" form:"email" binding:"required, email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
// }

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
