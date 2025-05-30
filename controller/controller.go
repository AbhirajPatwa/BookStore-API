
package controller

import (
	"net/http"
	"github.com/AbhirajPatwa/BookStore-API/config"

	"github.com/labstack/echo/v4"
)

type Book struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateBook(c echo.Context) error {
	b := new(Book)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	book := &Book{
		Name:        b.Name,
		Description: b.Description,
	}

	if err := db.Create(&book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	b := new(Book)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_book := new(Book)

	if err := db.First(&existing_book, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_book.Name = b.Name
	existing_book.Description = b.Description
	if err := db.Save(&existing_book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_book,
	}

	return c.JSON(http.StatusOK, response)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var books []*Book

	if res := db.Find(&books, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": books[0],
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	book := new(Book)

	err := db.Delete(&book, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a book has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}