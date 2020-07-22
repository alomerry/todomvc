package user

type LoginRequest struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterRequest struct {
	Name           string `form:"name" binding:"required"`
	Password       string `form:"password" binding:"required"`
	RepeatPassword string `form:"repeatPassword" binding:"required"`
}
