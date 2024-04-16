package main

import (
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"

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

		//本当はフロントのjsonをバインドする
		email := "sample44@gmail.com"
		var user models.User
		err := database.DB.Where("email = ?", email).First(&user).Error
		if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
						// レコードが見つからなかった場合の処理
						err := models.NewUser(database.DB)
						if err != nil {
							log.Println(err)
							return c.String(http.StatusInternalServerError, "Internal Server Error")
					}
				} else {
						// その他のエラー処理
						return c.String(http.StatusInternalServerError, "Internal Server Error")
				}
		} else {
				// レコードが見つかった場合の処理
				return c.String(http.StatusBadRequest, "アドレス重複")
		}
    return c.NoContent(http.StatusOK)
})

	// サーバーをポート8080で起動
	e.Logger.Fatal(e.Start(":8080"))
  

}