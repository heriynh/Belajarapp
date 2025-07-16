package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rocksus/go-restaurant-app/internal/database"
	"github.com/rocksus/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/rocksus/go-restaurant-app/internal/repository/menu"
	rUsecase "github.com/rocksus/go-restaurant-app/internal/usecase/resto"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=go_resto_app sslmode=disable"
)

func main() {

	e := echo.New()
	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	restoUsecase := rUsecase.GetUsecase(menuRepo)
	h := rest.NewHandler(restoUsecase)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}
