package handler

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

//SearchItem by desc
func SearchItem(desc string, db *gorm.DB) Response {
	items := []ItemsVenta{}
	db.Where("descripcion_erp LIKE ? or f120_referencia LIKE ?", "%"+desc+"%", desc+"%").Order("descripcion_erp asc").Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetExt1 by cod
func GetExt1(codigo int, idListaPrecio, bodega string, db *gorm.DB) Response {
	items := []ItemsVentaErp{}
	db.Distinct("ext1, ext1_color").Where("codigo_erp = ? and id_lista_precio = ? and f150_id = ?", codigo, idListaPrecio, bodega).Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

//GetExt2 by cod
func GetExt2(codigo int, idListaPrecio string, ext1 string, bodega string, db *gorm.DB) Response {
	items := []ItemsVentaErp{}
	db.Where("codigo_erp = ? and id_lista_precio = ? and ext1 = ? and f150_id = ?", codigo, idListaPrecio, ext1, bodega).Find(&items)
	return Response{Payload: items, Message: "OK", Status: 200}
}

func GetBodegas(userID string, db *gorm.DB) Response {
	bodegas := []appVendedorBodega{}
	db.Where("f200_id like ?", userID+" %").Find(&bodegas)
	return Response{Payload: bodegas, Message: "OK", Status: 200}
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
	return consec.ConsecutivoPedido
}

//SavePedidoErp fuc
func SavePedidoErp(pedido *Pedido, db *gorm.DB) Response {
	consecutivo := getConsecPedido(db)
	pedido.InfoPedido.PvcCotNum = consecutivo
	pedido.InfoPedido.PvcCenOper = "001"
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
	db.Model(ConsecPedido{}).Where("consecutivo_pedido = ?", consecutivo).Omit("Descripcion").Updates(ConsecPedido{ConsecutivoPedido: (consecutivo + 1)})
	return Response{Payload: consecutivo, Message: "Registro Realizado! Pedido Nro: " + strconv.Itoa(consecutivo), Status: 201}
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

//GetCarteraCliente by nit and plac
func GetCarteraCliente(nit string, sucursal string, db *gorm.DB) Response {
	carteras := []CarteraClientesErp{}
	db.Where("nit = ? and id_sucursal = ?", nit, sucursal).Find(&carteras)
	return Response{Payload: carteras, Message: "OK", Status: 200}
}

//GetSaldoCliente by nit and plac
func GetSaldoCliente(nit string, sucursal string, db *gorm.DB) Response {
	var saldo SucursalesErp
	db.Where("nit_tercero = ? and f201_id_sucursal = ?", nit, sucursal).First(&saldo)
	return Response{Payload: saldo, Message: "OK", Status: 200}
}

//GetFolders get all folder
func GetFolders() Response {
	folders := []string{}
	files, err := ioutil.ReadDir("C:/catalogos")
	if err != nil {
		return Response{Payload: folders, Message: "OK", Status: 200}
	}
	for _, f := range files {
		folders = append(folders, f.Name())
	}
	return Response{Payload: folders, Message: "OK", Status: 200}
}

//GetPhotos get all photos
func GetPhotos(folder string) Response {
	folders := []string{}
	files, err := ioutil.ReadDir("C:/catalogos/" + folder)
	if err != nil {
		return Response{Payload: folders, Message: "OK", Status: 200}
	}
	for _, f := range files {
		folders = append(folders, f.Name())
	}
	return Response{Payload: folders, Message: "OK", Status: 200}
}

//GetPhotoBase64 photo base64
func GetPhotoBase64(folder string, photo string, db *gorm.DB) Response {
	// Open file on disk.
	f, _ := os.Open("C:/catalogos/" + folder + "/" + photo)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	data := struct {
		Base64      string
		Descripcion string
		Precio      float32
	}{
		encoded,
		"",
		0,
	}
	return Response{Payload: data, Message: "OK", Status: 200}
}

func getItemVentaERP(referencia string, db *gorm.DB) ItemsVentaErp {
	item := ItemsVentaErp{}
	ref := strings.Split(referencia, ".")
	db.Where("referencia = ?", ref[0]).Limit(1).Find(&item)
	return item
}

//GetPedidoERP get pedido ERP
func GetPedidoERP(num int, db *gorm.DB) Response {
	pedido := PedidoErp{}
	db.Where("pvc_cot_num = ?", num).Limit(1).Find(&pedido)
	detalle := []ResumenPedido{}
	db.Where("pvc_cot_num = ?", num).Find(&detalle)
	cliente := ClienteErp{}
	db.Where("nit_tercero = ?", pedido.PvcDocID).First(&cliente)
	punto := ClientesPuntosEnvioErp{}
	db.Where("f215_rowid = ?", pedido.F215ID).First(&punto)
	data := struct {
		Pedido     PedidoErp
		Detalle    []ResumenPedido
		Cliente    ClienteErp
		PuntoEnvio ClientesPuntosEnvioErp
	}{
		pedido,
		detalle,
		cliente,
		punto,
	}
	return Response{Payload: data, Message: "OK", Status: 200}
}

//GetPuntosEnvios fuc
func GetPedidosUser(vendedorID string, nit string, db *gorm.DB) Response {
	pedidos := []PedidoErp{}
	if nit == "-1" {
		db.Limit(100).Where("pvc_doc_vendedor = ?", vendedorID).Find(&pedidos)
		return Response{Payload: pedidos, Message: "OK", Status: 200}
	}
	db.Limit(100).Where("pvc_doc_vendedor = ? AND pvc_doc_id LIKE ?", vendedorID, nit+"%").Find(&pedidos)
	return Response{Payload: pedidos, Message: "OK", Status: 200}
}
