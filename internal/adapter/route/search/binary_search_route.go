package search

import (
	"github.com/gin-gonic/gin"

	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/adapter/handler/search"
	searchPort "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/port/search"
)

func BinarySearchRoute(route *gin.Engine, basePath string, port searchPort.BinarySearcher) {
	handler := search.NewBinarySearchHandler(port)

	binarySearchRoute := route.Group(basePath + "/binary")
	binarySearchRoute.POST("", handler.BinarySearch)
}
