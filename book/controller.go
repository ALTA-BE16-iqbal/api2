package book

import (
	"net/http"
	"strconv"
	"tugas-1-API/helper"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	model BookModel
}

func (bc *BookController) SetModel(m BookModel) {
	bc.model = m
}

// func AllBooks(c echo.Context) error {
// 	return c.JSON(http.StatusOK, listBooks)
// }

func (bc *BookController) AddBooks(c echo.Context) error {
	input := Books{}
	err := c.Bind(&input)

	if err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari Book", nil))
	}
	res, err := bc.model.Insert(input)

	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
	}

	return c.JSON(helper.ReponsFormat(http.StatusCreated, "sukses menambahkan data", res))
}

func (bc *BookController) GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.QueryParam("user_id")
		cnv := 0
		if userID != "" {
			cnvID, err := strconv.Atoi(userID)
			cnv = cnvID
			if err != nil {
				c.Logger().Error("Input error ", err.Error())
				return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan pada input ID", nil))
			}
		}

		res, err := bc.model.GetAllBook(uint(cnv))

		if err != nil {
			c.Logger().Error("Book model error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}
}

func (bc *BookController) GetBookByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputParameter := c.Param("bookId")
		cnv, err := strconv.Atoi(inputParameter)
		if err != nil {
			c.Logger().Error("Input error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan pada input ID", nil))
		}
		res, err := bc.model.GetBookByID(uint(cnv))

		if err != nil {
			c.Logger().Error("Book model error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}
}
func (bc *BookController) DeleteBookByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputParameter := c.Param("bookId")
		cnv, err := strconv.Atoi(inputParameter)
		if err != nil {
			c.Logger().Error("Input error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan pada input ID", nil))
		}
		res, err := bc.model.DeleteBookByID(uint(cnv))

		if err != nil {
			c.Logger().Error("Book model error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses Menghapus data", res))
	}
}
