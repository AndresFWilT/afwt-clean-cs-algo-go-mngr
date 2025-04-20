package sort

import (
	"github.com/gin-gonic/gin"

	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/adapter/handler/sort"
	sorterPort "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/port/sort"
)

func BubbleRoute(router *gin.Engine, basePath string, sorter sorterPort.Sorter) {
	bubbleSortHandler := sort.NewBubbleSortHandler(sorter)

	bubbleRoute := router.Group(basePath + "/bubble")
	bubbleRoute.POST("", bubbleSortHandler.SortBubbleAlgorithm)
}
