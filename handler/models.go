package handler

import "time"

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

//ItemsExtension1 Struct
type ItemsExtension1 struct {
	CodigoErp       int
	Referencia      string
	Descripcion     string
	F117Descripcion string
	FechaCreado     time.Time
	F117ID          string
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
