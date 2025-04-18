package sort

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

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

	response, err := s.Service.Sort(data.Array)

	c.JSON(200, gin.H{
		"data": response,
	})
}
