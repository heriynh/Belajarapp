package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica-rica",
			Price:     41250,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error {
	DrinkMenu := []MenuItem{
		{
			Name:      "Jus Apel",
			OrderCode: "jus_apel",
			Price:     25700,
		},
		{
			Name:      "Jus Mangga",
			OrderCode: "jus_mangga",
			Price:     28250,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": DrinkMenu,
	})
}

func main() {
	e := echo.New()
	//localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
