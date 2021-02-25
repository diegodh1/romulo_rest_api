package handler

import "gorm.io/gorm"

//SearchClient by name or last name
func SearchClient(id string, db *gorm.DB) Response {
	clients := []ClienteErp{}
	db.Where("nombre_tercero LIKE ? or nit_tercero LIKE ?", "%"+id+"%", id+"%").Find(&clients)
	return Response{Payload: clients, Message: "OK", Status: 200}
}

//GetPersonalInfo func
func GetPersonalInfo(nit string, db *gorm.DB) Response {
	personalInfo := []ClienteContactosErp{}
	db.Where("nit_cc = ?", nit).Find(&personalInfo)
	return Response{Payload: personalInfo, Message: "OK", Status: 200}
}
