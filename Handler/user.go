package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func Register(c *gin.Context) {
	//get name, email, number, password
	var get model.Regist
	if err := c.ShouldBindJSON(&get); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon lengkapi input Anda", err)
		return
	}
	if get.Password != get.Passconfirm {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Password tidak sama",
		})
		return
	}
	//Hashing
	hash, err := bcrypt.GenerateFromPassword([]byte(get.Password), bcrypt.DefaultCost)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to Hash", err)
		return
	}
	//Create
	user := entity.User{
		Model:    gorm.Model{},
		Nama:     get.Nama,
		Email:    get.Email,
		Password: string(hash),
		Number:   get.Number,
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to create", err)
		return
	}

	//Respond
	sdk.Success(c, http.StatusOK, "Success to Register", user)
}

func LogIn(c *gin.Context) {
	var body model.LogIn
	if err := c.ShouldBindJSON(&body); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Error to read", err)

		return
	}

	//cari data
	var req entity.User
	database.DB.First(&req, "email = ?", body.Email)
	if req.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Invalid Email / Password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(body.Password))
	if err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Failed to compare the password", err)
		return
	}
	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": req.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	//Send back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	sdk.Success(c, http.StatusOK, "User berhasil log in", map[string]string{"token": tokenString})
}

func Logoout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	sdk.Success(c, http.StatusOK, "Anda telah log out", "")
}

func Validate(c *gin.Context) {
	id := c.MustGet("user")

	var user entity.User

	err := database.DB.First(&user, id)
	if err.Error != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not found", err.Error)
		return
	}

	//if err.RowsAffected == 0 {
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"error": err.Error.Error(),
	//	})
	//	return
	//}

	c.JSON(200, gin.H{
		"data":    user.Nama,
		"error":   nil,
		"message": "logged in",
	})
}
