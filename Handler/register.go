package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
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
