package main

import (
	"tugas-1-API/book"
	"tugas-1-API/config"
	"tugas-1-API/routes"
	"tugas-1-API/user"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	cfg := config.InitSQL()
	cfg.AutoMigrate(user.User{})
	cfg.AutoMigrate(book.Books{})

	mdl := user.UserModel{}
	mdl.SetDB(cfg)
	ctl := user.UserController{}
	ctl.SetModel(mdl)

	bookMdl := book.BookModel{}
	bookMdl.SetDB(cfg)
	bookCtl := book.BookController{}
	bookCtl.SetModel(bookMdl)

	routes.Route(e, ctl, bookCtl)

	e.Start(":8000")
}
