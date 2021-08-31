package model

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Age      int    `form:"age" binding:"gte=18"` // >= 18
}
