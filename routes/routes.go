package routes

import (
	"fmt"
	"os"
	handler "romulo/handler"
	"strconv"
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

func enviarCorreo(correo string, tipoDoc string, nroDoc string, name string, lastName string, cellphone string, phone string, dir string, ciudad string, pais string, routesF []string) bool {
	from := "noreply-ventas@calzadoromulo.com.co"
	pass := "Temporal.2021@"
	to := "hernando.gaitan@integrapps.com "

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
	<td>Ciudad</td>
	<td>Pais</td>
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
	<td>%s</td>
	<td>%s</td>
	</tr>
	</tbody>
	</table>
	<p>Muchas Gracias,</p>
	<p>PD: ajunto documentos</p></body>
	</html>`, correo, tipoDoc, nroDoc, name, lastName, cellphone, phone, dir, ciudad, pais))

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

func enviarSolicitud(nit string, nombre string) bool {
	from := "noreply-ventas@calzadoromulo.com.co"
	pass := "Temporal.2021@"
	to := "hernando.gaitan@integrapps.com "

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
	<p>El siguiente correo es para hacer una solicitud de ampliacion de cupo para el cliente %s identificado con el nit Nro. %s</p>
	<p>Muchas Gracias,</p>
	<p>PD: ajunto documentos</p></body>
	</html>`, nombre, nit))

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
			c.SaveUploadedFile(file, "C:/profilePhotos/"+nroDoc+"-"+file.Filename)
			routesF = append(routesF, "C:/profilePhotos/"+nroDoc+"-"+file.Filename)
		}
		tipoDoc := c.PostForm("tipoDoc")
		name := c.PostForm("name")
		lastName := c.PostForm("lastName")
		email := c.PostForm("email")
		celphone := c.PostForm("cellphone")
		phone := c.PostForm("phone")
		dir := c.PostForm("dir")
		ciudad := c.PostForm("ciudad")
		pais := c.PostForm("pais")
		enviado := enviarCorreo(email, tipoDoc, nroDoc, name, lastName, celphone, phone, dir, ciudad, pais, routesF)
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
		id := c.Param("id")
		id = strings.ToUpper(strings.ReplaceAll(id, "%", " "))
		response := handler.SearchClient(id, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//RealizarSolicitudCupo realizar solicitud
func RealizarSolicitudCupo(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nit := c.Param("nit")
		nombre := c.Param("nombre")
		enviarSolicitud(nit, nombre)
		c.JSON(200, gin.H{
			"payload": "correo enviado",
			"message": "correo enviado",
			"status":  200,
		})
	}
	return fn
}

//GetPedidosUser get all pedidos by seller
func GetPedidosUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nit := c.Param("nit")
		vendedorID := c.Param("vendedorID")
		response := handler.GetPedidosUser(vendedorID, nit, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//SearchItem func
func SearchItem(db *gorm.DB) gin.HandlerFunc {
	var item handler.ItemsVenta
	fn := func(c *gin.Context) {
		err := c.BindJSON(&item)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "petición mal estructurada",
				"status":  400,
			})
		default:
			item.DescripcionErp = strings.ToUpper(item.DescripcionErp)
			response := handler.SearchItem(item.DescripcionErp, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return fn
}

//SearchUser func
func SearchUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userID := c.Param("userID")
		response := handler.SearchUser(userID, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetColecciones func
func GetColecciones(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		response := handler.GetColecciones(db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetExt1 func
func GetExt1(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		ext1 := c.Param("code")
		v, err := strconv.Atoi(ext1)
		if err != nil {
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "petición mal estructurada",
				"status":  400,
			})
		}
		list := c.Param("list")
		bodega := c.Param("bodega")
		response := handler.GetExt1(v, list, bodega, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetExt1 func
func GetBodegas(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userID := c.Param("userID")
		response := handler.GetBodegas(userID, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetExt2 func
func GetExt2(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		code := c.Param("code")
		v, err := strconv.Atoi(code)
		if err != nil {
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "petición mal estructurada",
				"status":  400,
			})
		}
		list := c.Param("list")
		ext1 := c.Param("ext1")
		bodega := c.Param("bodega")
		response := handler.GetExt2(v, list, ext1, bodega, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetFolders func
func GetFolders(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		response := handler.GetFolders()
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetPhotos func
func GetPhotos(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		folder := c.Param("folder")
		folder = strings.ReplaceAll(folder, "%", " ")
		response := handler.GetPhotos(folder)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetPhotoBase64 func
func GetPhotoBase64(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		folder := c.Param("folder")
		folder = strings.ReplaceAll(folder, "%", " ")
		photo := c.Param("photo")
		photo = strings.ReplaceAll(photo, "%", " ")
		response := handler.GetPhotoBase64(folder, photo, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetPuntosDeEnvio func
func GetPuntosDeEnvio(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nit := c.Param("nit")
		response := handler.GetPuntosEnvios(nit, db)
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

//SavePedido func
func SavePedido(db *gorm.DB) gin.HandlerFunc {
	var pedido handler.Pedido
	fn := func(c *gin.Context) {
		err := c.BindJSON(&pedido)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "petición mal estructurada",
				"status":  400,
			})
		default:
			response := handler.SavePedidoErp(&pedido, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return fn
}

//GetItemsFotos func
func GetItemsFotos(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		tempo := c.Param("temporada")
		response := handler.GetItemsFotos(tempo, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetCarteraCliente func
func GetCarteraCliente(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nit := c.Param("nit")
		sucursal := c.Param("sucursal")
		response := handler.GetCarteraCliente(nit, sucursal, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetSaldoCliente func
func GetSaldoCliente(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nit := c.Param("nit")
		sucursal := c.Param("sucursal")
		response := handler.GetSaldoCliente(nit, sucursal, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetPedidoERP func
func GetPedidoERP(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		pedido := c.Param("pedido")
		v, err := strconv.Atoi(pedido)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "petición mal estructurada",
				"status":  400,
			})
		default:
			response := handler.GetPedidoERP(v, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return fn
}
