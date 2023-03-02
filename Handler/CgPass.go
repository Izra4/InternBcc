package Handler

//
//import (
//	"InternBCC/database"
//	"InternBCC/entity"
//	"github.com/gin-gonic/gin"
//	"gopkg.in/gomail.v2"
//	"os"
//)
//
//func ChangePass(c *gin.Context) {
//	var get struct {
//		Email string `json:"email" binding:"required"`
//	}
//	if err := c.Bind(&get); err != nil {
//		c.JSON(400, gin.H{
//			"error": "failed to read",
//		})
//	}
//	var req entity.User
//	database.DB.First(&req, "email = ?", get.Email)
//	if req.ID == 0 {
//		c.JSON(400, gin.H{
//			"Error": "Invalid Email",
//		})
//		return
//	}
//	m := gomail.NewMessage()
//	m.SetHeader("From", "Grent@gmail.com")
//	m.SetHeader("To", get.Email)
//	m.SetHeader("Subject", "Hello!")
//	m.SetBody("text/html", "<h1>Hello<h1> <p>To change your email please click this link!<p>")
//
//	d := gomail.NewDialer("smtp.gmail.com", 587, "aryaizra2@gmail.com", os.Getenv("PASS"))
//
//	// Send the email to Bob, Cora and Dan.
//	if err := d.DialAndSend(m); err != nil {
//		panic(err)
//	}
//}
