package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateSchemaDb() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	dsn := "sqlserver://sa:b123@localhost:1433?database=SM.V1" //ConnectionStringURL
	gormdb, _ := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	g.UseDB(gormdb)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}

var (
	db *gorm.DB
)

func ConnectDb() (err error) {
	dsn := "sqlserver://sa:b123@localhost:1433?database=SM.V1"

	d, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db = d
	return
}
func GetDb() *gorm.DB {
	return db
}
