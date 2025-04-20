package search

import "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/search"

type LinearSearcher interface {
	Search(data search.LinearSearchRequest) int64
}
