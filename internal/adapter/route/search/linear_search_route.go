package search

import (
	"github.com/gin-gonic/gin"

	searchHandler "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/adapter/handler/search"
	searchPort "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/port/search"
)

func LinearSearchRoute(g *gin.Engine, basePath string, port searchPort.LinearSearcher) {
	handler := searchHandler.NewLinearSearchHandler(port)

	linearSearchRoute := g.Group(basePath + "/linear")
	linearSearchRoute.POST("", handler.LinearSearch)
}
