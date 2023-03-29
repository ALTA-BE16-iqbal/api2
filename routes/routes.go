package routes

import (
	"tugas-1-API/book"
	"tugas-1-API/user"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, uc user.UserController, bc book.BookController) {
	e.POST("/register", uc.RegisterUser)
	e.POST("/login", uc.Login())
	e.GET("/users", uc.GetUser())

	e.GET("/books/:bookId", bc.GetBookByID())
	e.GET("/deletebook/:bookId", bc.DeleteBookByID())
	e.GET("/books", bc.GetBook())
	e.POST("/books", bc.AddBooks)
}
