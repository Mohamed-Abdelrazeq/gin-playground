package entity

type User struct {
	UserData  UserData  `json:"userData"`
	UserToken UserToken `json:"userToken"`
}

type UserData struct {
	Email     string `json:"email" binding:"required,email"`
	Passoword string `json:"password" binding:"required,min=6,max=16"`
	Name      string `json:"name" binding:"required,min=3,max=16"`
}

type LoginUser struct {
	Email     string `json:"email" binding:"required,email"`
	Passoword string `json:"password" binding:"required,min=6,max=16"`
}

type UserToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
