package handler

import (
	"io"
	"library_api/contract"
	"library_api/dto"
	"library_api/middleware"
	"library_api/pkg/errs"
	"library_api/pkg/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lendingController struct {
	service contract.LendingService
}

func (l *lendingController) getPrefix() string {
	return "/lendings"
}

func (l *lendingController) initService(service *contract.Service) {
	l.service = service.Lending
}

func (l *lendingController) initRoute(app *gin.RouterGroup) {
	app.GET("/", middleware.AdminCheck, l.getAllLendings)
	app.POST("/makeLending", middleware.MemberCheck, l.makeLending)
	app.PUT("/return/:lendingID", middleware.MemberCheck, l.returnBook)
}

func (l *lendingController) getAllLendings(ctx *gin.Context) {
	genre := ctx.Query("status")
	keyword := ctx.Query("keyword")
	var result *dto.LendingResponse
	var err error

	switch {
	case keyword != "":
		result, err = l.service.SearchLendings(keyword)
	case genre != "":
		result, err = l.service.GetLendingsByStatus(genre)
	default:
		result, err = l.service.GetAllLendings()
	}

	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (l *lendingController) makeLending(ctx *gin.Context) {
	user, valid := ctx.Get("users")

	if !valid {
		ctx.JSON(http.StatusBadRequest, errs.ErrValid)
		return
	}

	userData := user.(*token.UserAuthToken)

	var payload dto.LendingRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		handlerError(ctx, err)
		return
	}

	result, err := l.service.MakeLending(userData.ID, &payload)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (l *lendingController) returnBook(ctx *gin.Context) {
	lendingIDStr := ctx.Param("lendingID")
	lendingID, err := strconv.ParseUint(lendingIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Lending ID"})
		return
	}

	user, valid := ctx.Get("users")
	if !valid {
		ctx.JSON(http.StatusBadRequest, errs.ErrValid)
		return
	}
	userData := user.(*token.UserAuthToken)

	// Get the uploaded file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Open the file for reading
	fileReader, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer fileReader.Close()

	// Read the file bytes (optional: save to disk or process as needed)
	fileBytes, err := io.ReadAll(fileReader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read uploaded file"})
		return
	}

	// Pass fileBytes to your service layer
	result, err := l.service.ReturnBook(lendingID, userData.ID, fileBytes)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
