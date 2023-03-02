package model

type Regist struct {
	Nama        string `json:"nama" binding:"required"`
	Email       string `json:"email" binding:"email,required"`
	Password    string `json:"password" binding:"required"`
	Passconfirm string `json:"passconfirm" binding:"required"`
	Number      string `json:"number" binding:"required"`
}

type LogIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
