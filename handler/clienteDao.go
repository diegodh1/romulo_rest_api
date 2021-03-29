package handler

import "gorm.io/gorm"

//SearchClient by name or last name
func SearchClient(id string, userID int, db *gorm.DB) Response {
	clients := []ClientesVendedoresErp{}
	db.Where("row_id_vendedor = ? and (cliente LIKE ? or nit LIKE ?)", userID, "%"+id+"%", id+"%").Find(&clients)
	return Response{Payload: clients, Message: "OK", Status: 200}
}

//GetPersonalInfo func
func GetPersonalInfo(nit string, db *gorm.DB) Response {
	personalInfo := []ClienteContactosErp{}
	db.Where("nit_cc = ?", nit).Find(&personalInfo)
	return Response{Payload: personalInfo, Message: "OK", Status: 200}
}
