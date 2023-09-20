package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	dsn := "sqlserver://sa:b123@localhost:1433?database=SM.V1"
	gormdb, _ := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	g.UseDB(gormdb)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
