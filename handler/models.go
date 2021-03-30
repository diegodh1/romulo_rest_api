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
	NitCC         string
	Direccion     string
	Email         string
	Telefono      string
	Celular       string
	IDListaPrecio string
}

//ClientesVendedoresErp struct
type ClientesVendedoresErp struct {
	Nit           string
	Cliente       string
	RowIdVendedor int
	IDSucursal    string
	IDListaPrecio string
	DocVendedor   string
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
	CodigoErp     int
	Referencia    string
	Descripcion   string
	Ext1          string
	Ext2          string
	PrecioUnt     float32
	UndPrecio     string
	Ext1Color     string
	Existencia    float32
	F150ID        string
	IDListaPrecio string
	F200ID        int
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
	IDListaPrecio  string
	F215ID         string
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

//ClientesEstadoCupo Struct
type ClientesEstadoCupo struct {
	Nit         string
	BloqueoCupo int
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
	PvcDetCant         float32
	PvcDetNota         string
	PvcDetReferencia   string
	PvcDetExt1         string
	PvcDetExt2         string
	PvcDetPrecioUnt    float32
	PvcDetListaPrecio  string
}

//ConsecPedido Struct
type ConsecPedido struct {
	ConsecutivoPedido int
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

//CarteraClientesErp Struct
type CarteraClientesErp struct {
	Nit        string
	IDSucursal string
	NoDoc      int
	FechaVence *time.Time
	Saldo      float32
}

//SucursalesErp struct
type SucursalesErp struct {
	NitTercero             string
	F201IDSucursal         string
	F201IndEstadoBloqueado int
	CupoCredito            float32
}

//ResumenPedido struct
type ResumenPedido struct {
	Referencia     string
	Color          string
	Talla          string
	Cantidad       float32
	PrecioUnitario float32
	Total          float32
	Fecha          string
	PvcCotNum      int
}

//ResumenPedido struct
type appVendedorBodega struct {
	F200ID string
	F150ID string
}
