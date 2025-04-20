package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	searchRoute "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/adapter/route/search"
	sortRoute "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/adapter/route/sort"
	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/usecase/search"
	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/usecase/sort"
)

func InitializeServer() (*gin.Engine, string) {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}

	router := gin.Default()
	configureRoutes(router)
	return router, os.Getenv("API_PORT")
}

func configureRoutes(router *gin.Engine) {
	configureSorterAlgorithms(router, os.Getenv("API_BASE_PATH"))
	configureSearchAlgorithms(router, os.Getenv("API_BASE_PATH"))
}

func configureSorterAlgorithms(router *gin.Engine, routerBasePath string) {
	routerBasePath = routerBasePath + "/sort"
	bubbleSortService := sort.NewBubbleSortService()
	sortRoute.BubbleRoute(router, routerBasePath, &bubbleSortService)
}

func configureSearchAlgorithms(router *gin.Engine, routerBasePath string) {
	routerBasePath = routerBasePath + "/search"

	linearSearchService := search.NewLinearSearchService()
	searchRoute.LinearSearchRoute(router, routerBasePath, &linearSearchService)

	binarySearchService := search.NewBinarySearch()
	searchRoute.BinarySearchRoute(router, routerBasePath, &binarySearchService)
}
