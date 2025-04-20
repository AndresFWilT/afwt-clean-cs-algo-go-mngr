package search

import "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/search"

type LinearSearch struct{}

func NewLinearSearchService() LinearSearch {
	return LinearSearch{}
}

func (l LinearSearch) Search(data search.LinearSearchRequest) int64 {
	array, target := data.Array, data.Target
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return int64(i)
		}
	}
	return -1
}
