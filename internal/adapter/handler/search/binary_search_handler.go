package search

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	responseModel "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/response"
	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/search"
	searchPort "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/port/search"
)

type BinarySearchHandler struct {
	Service searchPort.BinarySearcher
}

func NewBinarySearchHandler(port searchPort.BinarySearcher) BinarySearchHandler {
	return BinarySearchHandler{
		Service: port,
	}
}

func (h BinarySearchHandler) BinarySearch(c *gin.Context) {
	data := search.BinarySearchRequest{}
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := h.Service.SearchSorted(data)
	var statusCode int
	if result == -1 {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusOK
	}

	response := &responseModel.Generic{
		Base: responseModel.Base{
			StatusCode: statusCode,
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       "some-uuid",
		},
		Description: "Request processed successfully, used binary search",
		Data:        result,
	}

	c.JSON(statusCode, response)

}
