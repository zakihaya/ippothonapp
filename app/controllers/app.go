package controllers


import (
    "fmt"
    "github.com/robfig/revel"
    "ippothonapp/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    InitDB()
    rows, _ := DbMap.Select(models.Stock{}, "select * from stock")
    for _, row := range rows {
        stock := row.(*models.Stock)
        fmt.Printf("%d, %s\n", stock.Id, stock.CompanyName)
    }

    return c.Render(rows)
}
