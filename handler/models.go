package handler

import "time"

//AppUser struct
type AppUser struct {
	UserID       string
	Password     string
	Name         string
	LastName     string
	Email        string
	Phone        string
	Photo        string
	Status       *bool
	CreationDate *time.Time
}

//AppUserProfile struct
type AppUserProfile struct {
	UserID    string
	ProfileID string
}

//ClienteContactoErp struct
type ClienteContactoErp struct {
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
	Contacto      *[]ClienteContactoErp
}
