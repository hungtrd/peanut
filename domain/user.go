package domain

type User struct {
	BaseModel
	Username    string `gorm:"unique"`
	DisplayName string
	Email       string
	Password    string
}

type LoginReq struct {
	Username string `json:"username" binding:"required" example:"hungtran" extensions:"x-order=1"`
	Password string `json:"password" binding:"required" example:"thisispassword"`
}

type LoginResp struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3Jpem..."`
}

type SignupReq struct {
	Username    string `json:"username" binding:"required" example:"hungtran" extensions:"x-order=1"`
	DisplayName string `json:"display_name" binding:"required" example:"Hung Tran" extensions:"x-order=2"`
	Email       string `json:"email" binding:"required,email" example:"hung@example.com" extensions:"x-order=3"`
	Password    string `json:"password" binding:"required" example:"thisispassword" extensions:"x-order=4"`
}

type SignupResp struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3Jpem..."`
}
