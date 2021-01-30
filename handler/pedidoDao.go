package handler

import "gorm.io/gorm"

//SearchItem by desc
func SearchItem(desc string, db *gorm.DB) Response {
	items := []ItemsVenta{}
	db.Where("descripcion_erp LIKE ?", "%"+desc+"%").Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetExt1 by cod
func GetExt1(codigo int, db *gorm.DB) Response {
	items := []ItemsExtension1{}
	db.Where("codigo_erp = ?", codigo).Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetExt2 by cod
func GetExt2(codigo int, db *gorm.DB) Response {
	items := []ItemsExtension2{}
	db.Where("codigo_erp = ?", codigo).Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}
