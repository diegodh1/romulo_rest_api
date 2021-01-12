package routes

import (
	handler "romulo/handler"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//LoginUser struct login
type LoginUser struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

//USER ROUTES

//Login route
func Login(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var userLogin LoginUser
		err := c.BindJSON(&userLogin)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
				"status":  400,
			})
		default:
			response := handler.LoginUser(userLogin.UserID, userLogin.Password, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//CreateUser func
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.AppUser
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.CreateUser(&user, db)
			c.JSON(400, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//UpdateUser func
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	var user handler.AppUser
	fn := func(c *gin.Context) {
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "petición mal estructurada",
				"status":  400,
			})
		default:
			response := handler.UpdateUser(&user, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return fn
}

//AssignProfile func
func AssignProfile(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var profile handler.AppUserProfile
		err := c.BindJSON(&profile)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "Petición mal estructurada",
				"status":  400,
			})
		default:
			response := handler.AssignProfile(&profile, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}

	}
	return fn
}

//***************************CLIENT****************************

//SearchClient func
func SearchClient(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		name := c.Param("name")
		name = strings.ToUpper(strings.ReplaceAll(name, "%", ""))
		response := handler.SearchClient(name, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetInfoClient func
func GetInfoClient(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nit := c.Param("nit")
		response := handler.GetPersonalInfo(nit, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}
