package sort

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/response"
	modelSort "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/sort"
	port "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/port/sort"
)

type BubbleSortHandler struct {
	Service port.Sorter
}

func NewBubbleSortHandler(sort port.Sorter) BubbleSortHandler {
	return BubbleSortHandler{
		Service: sort,
	}
}

func (s *BubbleSortHandler) SortBubbleAlgorithm(c *gin.Context) {
	fmt.Println("BubbleSortHandler")

	data := modelSort.Array{}
	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := s.Service.Sort(data.Array)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	bubbleSortedResponse := &response.Generic{
		Base: response.Base{
			StatusCode: http.StatusOK,
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       "some-uuid",
		},
		Description: "Request Completed Successfully, array sorted via bubble sort algorithm",
		Data:        result,
	}

	c.JSON(200, *bubbleSortedResponse)
}
