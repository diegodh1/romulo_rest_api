package handler

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

//SearchItem by desc
func SearchItem(desc string, db *gorm.DB) Response {
	items := []ItemsVenta{}
	db.Where("descripcion_erp LIKE ?", "%"+desc+"%").Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetExt1 by cod
func GetExt1(codigo int, db *gorm.DB) Response {
	items := []ItemsVentaErp{}
	db.Where("codigo_erp = ?", codigo).Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetExt2 by cod
func GetExt2(codigo int, db *gorm.DB) Response {
	items := []ItemsExtension2{}
	db.Where("codigo_erp = ?", codigo).Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetPuntosEnvios fuc
func GetPuntosEnvios(nit string, db *gorm.DB) Response {
	puntos := []ClientesPuntosEnvioErp{}
	db.Where("nit = ?", nit).Find(&puntos)
	return Response{Payload: puntos, Message: "OK", Status: 200}
}

//GetPuntosEnvios fuc
func getConsecPedido(db *gorm.DB) int {
	consec := ConsecPedido{}
	db.Last(&consec)
	return *consec.ConsecutivoPedido
}

//SavePedidoErp fuc
func SavePedidoErp(pedido *Pedido, db *gorm.DB) Response {
	consecutivo := getConsecPedido(db)
	pedido.InfoPedido.PvcCotNum = consecutivo
	pedido.InfoPedido.PvcCenOper = "01"
	//Evento ERP
	evento := EventosErp{EventoTipo: "PV", EventoParam1: strconv.Itoa(consecutivo), EventoParam2: pedido.InfoPedido.PvcDocID, EventoPruebas: true}
	if err := db.Create(evento).Error; err != nil {
		fmt.Println(err.Error())
		return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
	}
	//Pedido ERP
	if err := db.Create(pedido.InfoPedido).Error; err != nil {
		fmt.Println(err.Error())
		return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
	}
	//Detalle ERP
	for _, v := range *pedido.DetallePedido {
		saveDetallePedido(&v, consecutivo, db)
	}
	return Response{Payload: nil, Message: "Registro Realizado!", Status: 201}
}

//saveDetallePedido func
func saveDetallePedido(detalle *PedidoErpDet, consec int, db *gorm.DB) {
	detalle.PvcCotNum = consec
	if err := db.Create(detalle).Error; err != nil {
		fmt.Println(err.Error())
	}
}

//GetItemsFotos by desc
func GetItemsFotos(temporada string, db *gorm.DB) Response {
	items := []ViewCategoriaFotos{}
	db.Where("categoria_id LIKE ?", "%"+temporada+"%").Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetColecciones by cod
func GetColecciones(db *gorm.DB) Response {
	colecciones := []CategoriaItem{}
	db.Where("activo = ?", true).Find(&colecciones)
	return Response{Payload: colecciones, Message: "OK", Status: 200}
}
