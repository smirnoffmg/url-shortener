package interfaces

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	db "github.com/smirnoffmg/url-shortener/internal/db"
	repositories "github.com/smirnoffmg/url-shortener/internal/repositories"
	services "github.com/smirnoffmg/url-shortener/internal/services"
)

func GetGinServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	urlsDbConn := db.Connect(os.Getenv("URLS_DB_DSN"))
	urlsRepo := repositories.NewUrlRecordsRepository(urlsDbConn)
	urlSvc := services.NewUrlService(urlsRepo, 4)

	visitsDbConn := db.Connect(os.Getenv("VISITS_DB_DSN"))
	visitsRepo := repositories.NewVisitsRepository(visitsDbConn)
	visitsSvc := services.NewVisitsService(visitsRepo)

	r.GET("/r/:alias", func(ctx *gin.Context) {
		alias := ctx.Param("alias")

		url_record := urlSvc.Get(alias)

		if url_record != nil {

			visitsSvc.SaveVisit(alias, ctx.ClientIP(), ctx.GetHeader("User-Agent"))

			ctx.Redirect(http.StatusFound, url_record.OriginalUrl)
		} else {
			ctx.String(http.StatusNotFound, "Not found")
		}
	})

	r.POST("/api/v1/create", func(ctx *gin.Context) {
		originalUrl := ctx.PostForm("original")

		alias, err := urlSvc.Create(originalUrl)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"alias": alias,
		})
	})

	r.GET("/api/v1/stats/:alias", func(ctx *gin.Context) {
		alias := ctx.Param("alias")

		visits, err := visitsSvc.GetVisitsByAlias(alias)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"visits": visits,
		})
	})

	return r
}
