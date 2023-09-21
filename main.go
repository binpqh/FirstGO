package main

import (
	"context"
	"firstGO/pkg/config"
	"firstGO/pkg/model"
	"firstGO/pkg/query"
	"fmt"
)

func main() {
	err := config.ConnectDb()
	if err == nil {
		fmt.Printf("Ket noi thanh cong")
	}
	query.SetDefault(config.GetDb())
	ctx := context.Background()
	err1 := query.Q.Class.Create(&model.Class{
		Name: "Lop2",
	})
	if err1 != nil {
		fmt.Printf("Create loi~")
	}
	u := query.Q.Class
	class, err2 := u.WithContext(ctx).Where(u.Name.Eq("Lop2")).Delete()
	if err2 != nil {
		fmt.Printf("Lay du lieu thanh cong", class.Error)
	}
}
