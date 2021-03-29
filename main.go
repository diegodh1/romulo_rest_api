package main

import (
	"log"
	handler "romulo/handler"
	routes "romulo/routes"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func decodeJWT(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, nil)
	if token == nil {
		return true
	}
	return false
}

//TokenMiddleware func
func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "Tokén no válido", "status": 401})
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		if decodeJWT(reqToken) {
			c.AbortWithStatusJSON(401, gin.H{"payload": nil, "message": "Tokén no válido", "status": 401})
			return
		}
		decodeJWT(reqToken)
		c.Next()
	}

}

//Main func
func main() {
	r := gin.New()

	//connect to db
	var connection *handler.Config
	connection = new(handler.Config)
	connection.Init()
	db, err := connection.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.POST("/user/login", routes.Login(db))
	r.Use(TokenMiddleware())
	r.POST("/user/create", routes.CreateUser(db))
	r.POST("/user/update", routes.UpdateUser(db))
	r.GET("/user/search/:userID", routes.SearchUser(db))
	r.GET("/client/search/:id/:userID", routes.SearchClient(db))
	r.GET("/client/info/:nit", routes.GetInfoClient(db))
	r.GET("/client/cartera/:nit/:sucursal", routes.GetCarteraCliente(db))
	r.GET("/client/saldo/:nit/:sucursal", routes.GetSaldoCliente(db))
	r.POST("/client/create", routes.CreateClient(db))
	r.POST("/pedido/search/item", routes.SearchItem(db))
	r.GET("/pedido/get/color/:code/:list/:userID", routes.GetExt1(db))
	r.GET("/pedido/get/talla/:code/:list/:ext1/:userID", routes.GetExt2(db))
	r.GET("/pedido/get/puntos/:nit", routes.GetPuntosDeEnvio(db))
	r.GET("/pedido/get/cupo/:nit", routes.GetBloqueoCupo(db))
	r.POST("/pedido/crear", routes.SavePedido(db))
	r.GET("/catalogo/get/colecciones", routes.GetFolders(db))
	r.GET("/catalogo/get/fotos/:folder", routes.GetPhotos(db))
	r.GET("/catalogo/get/fotos/:folder/:photo", routes.GetPhotoBase64(db))
	r.GET("/pedido/search/:pedido", routes.GetPedidoERP(db))
	r.GET("/pedido/getAll/:vendedorID/:nit", routes.GetPedidosUser(db))
	r.GET("/pedido/realizarSolicitud/:nombre/:nit", routes.RealizarSolicitudCupo(db))
	r.Run(":5000")
}
