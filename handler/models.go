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
