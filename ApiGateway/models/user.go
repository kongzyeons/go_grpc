package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type RegisterReq struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
}

type LoginRes struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	AccToken string `json:"accToken"`
	RefToken string `json:"refToken"`
}

type GetAllUserRes struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type GetByIDRes struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
