package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/book"
)

func RootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "Syakil Ahmad Ginanjar",
		"bio":  "Software Engineer",
	})
}

func BooksHandler(context *gin.Context) {
	id := context.Param("id")
	title := context.Param("title")

	context.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func PostBookHandler(context *gin.Context) {

	var bookInput book.BookInput

	err := context.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		context.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}

func QueryHandler(context *gin.Context) {

	title := context.Query("title")
	price := context.Query("price")

	context.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
