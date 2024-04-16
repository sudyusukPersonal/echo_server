package main

import (
	"log"
	"net/http"

	"github.com/sudyusukPersonal/echo_server/database"
	"github.com/sudyusukPersonal/echo_server/models"

	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	database.InitDatabase()

	// ルートルートに対するハンドラを定義
	e.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK,map[string]string {
				"home":"this is hoeme",
			})
	})

	e.POST("/new_user", func(c echo.Context) error {
    email := "newemailthird@gmail.com"
    emailExists, err := models.EmailExists(email)
    if err != nil {
        log.Println(err)
        return c.String(http.StatusInternalServerError, "Internal Server Error")
    }
    if emailExists {
        return c.String(http.StatusBadRequest, "アドレス重複")
    }
    if err := models.NewUser(database.DB, email); err != nil {
        log.Println(err)
        return c.String(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.NoContent(http.StatusNoContent)
})


	e.Logger.Fatal(e.Start(":8080"))
}