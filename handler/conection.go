package handler

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//Config struct
type Config struct {
	Server   string
	User     string
	Pass     string
	Port     string
	Database string
}

//Init the db
func (c *Config) Init() {
	c.Server = "172.16.5.3"
	c.User = "sa"
	c.Pass = "AdminSQL.2019$"
	c.Port = "1433"
	c.Database = "Integrapps_Pruebas"
}

//Connect to DB
func (c *Config) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		c.User,
		c.Pass,
		c.Server,
		c.Port,
		c.Database,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		return nil, err
	}
	fmt.Println("conectado a la base de datos")
	return db, nil
}
