package controllers

import (
	"net/http"
	"strconv"

	"open-contribute/pkg/db/models"
	"open-contribute/pkg/db/repositories"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	repo *repositories.BookRepository
}

func NewBookController(repo *repositories.BookRepository) *BookController {
	return &BookController{repo: repo}
}

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (c *BookController) AddBook(ctx *gin.Context) {
	var body AddBookRequestBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title:       body.Title,
		Author:      body.Author,
		Description: body.Description,
	}

	if err := c.repo.CreateBook(&book); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := c.repo.GetBookByID(uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	if err := c.repo.DeleteBook(book); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *BookController) GetBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := c.repo.GetBookByID(uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) GetBooks(ctx *gin.Context) {
	books, err := c.repo.GetAllBooks()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var body UpdateBookRequestBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := c.repo.GetBookByID(uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if err := c.repo.UpdateBook(book); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}
