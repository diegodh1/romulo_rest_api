package handler

import (
	"time"
)

//Response struct
type Response struct {
	Payload interface{}
	Message string
	Status  int
}

//AppUser struct
type AppUser struct {
	UserID        string
	Password      string
	Name          string
	LastName      string
	Email         string
	Phone         string
	Photo         string
	Status        *bool
	CreactionDate time.Time
}

//User struct
type User struct {
	User     AppUser
	Profiles []AppUserProfile
}

//AppUserProfile struct
type AppUserProfile struct {
	UserID       string
	ProfileID    string
	Status       *bool
	CreationDate time.Time
}

//ClientApp struct
type ClientApp struct {
	TipoDoc  string
	NroDoc   string
	Nombre   string
	Apellido string
	Correo   string
	Telefono string
	Ciudad   string
}

//ClienteContactosErp struct
type ClienteContactosErp struct {
	NitCC     string
	Direccion string
	Email     string
	Telefono  string
	Celular   string
}

//ClienteErp struct
type ClienteErp struct {
	NombreTercero string
	NitTercero    string
}

//ItemsVenta Struct
type ItemsVenta struct {
	CodigoErp      int
	DescripcionErp string
	F120Referencia string
}

//ItemsVentaErp Struct
type ItemsVentaErp struct {
	CodigoErp   int
	Referencia  string
	Descripcion string
	Ext1        string
	Ext2        string
	PrecioUnt   string
	UndPrecio   string
	Ext1Color   string
}

//ItemsExtension2 Struct
type ItemsExtension2 struct {
	CodigoErp       int
	Referencia      string
	Descripcion     string
	F119Descripcion string
	FechaCreado     time.Time
	IDExt2          string
}

//ClientesPuntosEnvioErp Struct
type ClientesPuntosEnvioErp struct {
	F201IDSucursal string
	Nit            string
	F215Rowid      int
	IDSucursal     string
	IDVendedor     string
	PuntoEnvio     string
}

//PedidoErp Struct
type PedidoErp struct {
	PvcCotNum       int
	PvcDocID        string
	PvcFechaEntrega *time.Time
	PvcDocVendedor  string
	F215ID          string
	PvcNotas        string
	PvcCenOper      string
}

//Pedido Struct
type Pedido struct {
	InfoPedido    *PedidoErp
	DetallePedido *[]PedidoErpDet
}

//PedidoErpDet Struct
type PedidoErpDet struct {
	PvcCotNum          int
	PvcDetFechaEntrega string
	PvcDetCant         int
	PvcDetNota         string
	PvcDetReferencia   string
	PvcDetExt1         string
	PvcDetExt2         string
	PvcDetPrecioUnt    int
}

//ConsecPedido Struct
type ConsecPedido struct {
	ConsecutivoPedido *int
	Descripcion       string
}

//EventosErp Struct
type EventosErp struct {
	EventoTipo    string
	EventoParam1  string
	EventoParam2  string
	EventoParam3  string
	EventoPruebas bool
}

//ViewCategoriaFotos Struct
type ViewCategoriaFotos struct {
	CodigoErp   int
	Referencia  string
	Ruta        string
	CategoriaID string
	Estado      bool
	Descripcion string
	PrecioUnt   float32
	UndPrecio   string
}

//CategoriaItem Struct
type CategoriaItem struct {
	CategoriaID string
	Descripcion string
	Year        int
	Activo      bool
}
