package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"

	repositories "github.com/smirnoffmg/url-shortener/internal/repositories"
	services "github.com/smirnoffmg/url-shortener/internal/services"
)

func GetGinServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	repo := repositories.NewDBRepository("test.db")
	svc := services.NewUrlService(repo, 4)

	r.GET("/r/:alias", func(ctx *gin.Context) {
		alias := ctx.Param("alias")

		url_record := svc.Get(alias)

		if url_record != nil {
			ctx.Redirect(http.StatusFound, url_record.OriginalUrl)
		} else {
			ctx.String(http.StatusNotFound, "Not found")
		}
	})

	r.POST("/api/v1/create", func(ctx *gin.Context) {
		originalUrl := ctx.PostForm("original")

		record := svc.Create(originalUrl)

		if record != nil {
			ctx.JSON(http.StatusCreated, record)
		} else {
			ctx.String(http.StatusBadRequest, "Bad request")
		}

	})

	return r
}
