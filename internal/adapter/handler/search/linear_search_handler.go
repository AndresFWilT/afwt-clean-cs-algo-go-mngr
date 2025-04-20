package search

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/response"
	searchModel "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/search"
	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/port/search"
)

type LinearSearchHandler struct {
	Service search.LinearSearcher
}

func NewLinearSearchHandler(service search.LinearSearcher) *LinearSearchHandler {
	return &LinearSearchHandler{
		Service: service,
	}
}

func (h LinearSearchHandler) LinearSearch(c *gin.Context) {
	data := searchModel.LinearSearchRequest{}
	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := h.Service.Search(data)
	var statusCode int

	if result == -1 {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusOK
	}

	responseJson := &response.Generic{
		Base: response.Base{
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       "some-uuid",
			StatusCode: statusCode,
		},
		Description: "Request Completed Successfully, used linear search",
		Data:        result,
	}

	c.JSON(statusCode, *responseJson)
}
