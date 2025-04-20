package search

import (
	"github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/domain/model/search"
)

type BinarySearch struct{}

func NewBinarySearch() BinarySearch {
	return BinarySearch{}
}

func (b *BinarySearch) SearchSorted(data search.BinarySearchRequest) int64 {
	target, nums := data.Target, data.SortedArray
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] == target {
			return int64(mid)
		}
	}
	return -1
}
