package search

import "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/search"

type BinarySearcher interface {
	SearchSorted(data search.BinarySearchRequest) int64
}
