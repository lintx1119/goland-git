package admin

// test 2
import (
	"github.com/gin-gonic/gin"
	"go-study/models"
	"net/http"
	"time"
)

type UserController struct {
	BaseController
}

func (u UserController) GetUser(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "GetUser Success",
	//})
	//u.success(c)

	// test query ALL
	//studentList := []models.Student{}
	//models.DB.Find(&studentList)
	//c.JSON(http.StatusOK, gin.H{
	//	"studentList": studentList,
	//})

	// test query by where
	studentList := []models.Student{}
	models.DB.Where("username=?", "LTX").Find(&studentList)
	c.JSON(http.StatusOK, gin.H{
		"studentList": studentList,
	})
}

func (u UserController) AddUser(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "AddUser Success",
	//})
	studentList := &models.Student{
		Username: "LTX1",
		Age:      124,
		Email:    "lintx1119@gmail.com",
		AddTime:  int(time.Now().Unix()),
	}
	models.DB.Create(studentList)
	u.success(c)
}
