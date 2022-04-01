package respfmt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(c *gin.Context, err string) {
	fmt.Println("badRequest:", err)
	c.JSON(http.StatusBadRequest, map[string]interface{}{
		"error": err,
		"data":  nil,
	})
}
func InternalServer(c *gin.Context, err string) {
	fmt.Println("err:", err)
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error": err,
		"data":  nil,
	})
}
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  data,
	})
}
