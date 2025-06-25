package handler

import (
	"io"
	"library_api/contract"
	"library_api/dto"
	"library_api/middleware"
	"library_api/pkg/errs"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	service contract.BookService
}

func (b *bookController) getPrefix() string {
	return "/books"
}

func (b *bookController) initService(service *contract.Service) {
	b.service = service.Book
}

func (b *bookController) initRoute(app *gin.RouterGroup) {
	app.GET("/", b.GetAllBooks)
	app.GET("/:id", b.GetBookByID)
	app.POST("/insertBook", middleware.AdminCheck, b.InsertBook)
	app.GET("/assets/pdf/:filename", b.ServePDF)
	app.GET("/assets/thumbnails/:filename", b.ServeThumbnail)
}

func (b *bookController) ServePDF(ctx *gin.Context) {
	filename := ctx.Param("filename")
	filepath := "assets/pdf/" + filename

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "PDF not found"})
		return
	}

	ctx.File(filepath)
}

func (b *bookController) ServeThumbnail(ctx *gin.Context) {
	filename := ctx.Param("filename")
	filepath := "assets/thumbnails/" + filename

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Thumbnail not found"})
		return
	}

	ctx.File(filepath)
}

func (b *bookController) GetAllBooks(ctx *gin.Context) {
	genre := ctx.Query("genre")
	search := ctx.Query("search")
	var result *dto.BookListResponse
	var err error

	if genre != "" {
		result, err = b.service.GetBooksByGenre(genre)
	} else if search != "" {
		result, err = b.service.SearchBooks(search)
	} else {
		result, err = b.service.GetAllBooks()
	}

	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (b *bookController) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.NewBadRequest(err.Error()))
		return
	}

	result, err := b.service.GetBookByID(idUint)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (b *bookController) InsertBook(ctx *gin.Context) {
	var payload dto.BookRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	pdfFile, err := ctx.FormFile("file_path")
	if err != nil {
		ctx.JSON(400, gin.H{"error": "PDF file is required"})
		return
	}

	pdfFileReader, err := pdfFile.Open()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to open PDF file"})
		return
	}
	defer pdfFileReader.Close()

	pdfBytes, err := io.ReadAll(pdfFileReader)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to read PDF file"})
		return
	}

	result, err := b.service.InsertBook(&payload, pdfBytes)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
