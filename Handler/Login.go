package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

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
		"data":    user,
		"error":   nil,
		"message": "logged in",
	})
}
