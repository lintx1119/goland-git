package admin

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (bc BaseController) success(c *gin.Context) {
	if v, ok := c.Get("user"); ok {
		c.String(200, "success"+v.(string))
	} else {
		c.String(200, "failed")
	}
}
