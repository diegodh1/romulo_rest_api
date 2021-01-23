package routes

import (
	"fmt"
	"os"
	handler "romulo/handler"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
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
		var user handler.User
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.CreateUser(&user.User, &user.Profiles, db)
			c.JSON(400, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func enviarCorreo(correo string, tipoDoc string, nroDoc string, name string, lastName string, cellphone string, phone string, dir string, routesF []string) bool {
	from := "noreply-ventas@calzadoromulo.com.co"
	pass := "Temporal.2021@"
	to := "diegodiazh1994@gmail.com"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Creación de Cliente")
	m.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	<style>
		body {
		   font-family: "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif; 
		   font-weight: 300;
		}
	</style>
	<html>
		<body><p>Hola buen d&iacute;a,</p>
	<p>El siguiente correo es para hacer una <strong>solicitud de creaci&oacute;n de cliente</strong> con los siguientes datos.</p>
	<table table style="border-collapse: collapse; background-color: #fF6FE49; border-style: solid;" border="1">
	<tbody>
	<tr>
	<td>Correo</td>
	<td>Tipo Doc</td>
	<td>Nro Doc</td>
	<td>Nombre</td>
	<td>Apellido</td>
	<td>Celular</td>
	<td>Tel&eacute;fono</td>
	<td>Direccion</td>
	</tr>
	<tr>
	<td>%s</td>
	<td>%s</td>
	<td>%s</td>
	<td>%s</td>
	<td>%s</td>
	<td>%s</td>
	<td>%s</td>
	<td>%s</td>
	</tr>
	</tbody>
	</table>
	<p>Muchas Gracias,</p>
	<p>PD: ajunto documentos</p></body>
	</html>`, correo, tipoDoc, nroDoc, name, lastName, cellphone, phone, dir))

	for i := 0; i < len(routesF); i++ {
		_, err := os.Open(routesF[i])
		if err == nil {
			m.Attach(routesF[i])
		}
	}

	// Send the email to Bob
	d := gomail.NewPlainDialer("smtpout.secureserver.net", 80, from, pass)
	if err := d.DialAndSend(m); err != nil {
		return false
	}
	return true
}

//CreateClient func
func CreateClient(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		nroDoc := c.PostForm("nroDoc")
		routesF := []string{}
		for _, file := range files {
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, "C:/profilePhotos/"+nroDoc+"-"+file.Filename+".pdf")
			routesF = append(routesF, "C:/profilePhotos/"+nroDoc+"-"+file.Filename+".pdf")
		}
		tipoDoc := c.PostForm("tipoDoc")
		name := c.PostForm("name")
		lastName := c.PostForm("lastName")
		email := c.PostForm("email")
		celphone := c.PostForm("cellphone")
		phone := c.PostForm("phone")
		dir := c.PostForm("dir")
		enviado := enviarCorreo(email, tipoDoc, nroDoc, name, lastName, celphone, phone, dir, routesF)
		switch {
		case enviado:
			c.JSON(200, gin.H{
				"payload": nil,
				"message": "Petición enviada con éxito",
				"status":  200,
			})
		default:
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "No se pudo enviar la petición",
				"status":  400,
			})
		}
	}
	return fn
}

//UpdateUser func
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	var user handler.User
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
			response := handler.UpdateUser(&user.User, &user.Profiles, db)
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

//SearchUser func
func SearchUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userID := c.Param("userID")
		userID = strings.ToUpper(strings.ReplaceAll(userID, "%", ""))
		response := handler.SearchUser(userID, db)
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
