package user_ser

type StructUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
