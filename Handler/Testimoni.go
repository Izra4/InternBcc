package Handler

import (
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Testi(c *gin.Context) {
	var in struct {
		Pesan string `json:"pesan"`
	}
	//id := c.MustGet("user")

	if err := c.ShouldBindJSON(&in); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "failed to read", err)
	}

}
