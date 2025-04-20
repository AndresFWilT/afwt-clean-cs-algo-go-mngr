package search

type BinarySearchRequest struct {
	SortedArray []int64 `json:"sorted_array"`
	Target      int64   `json:"target"`
}
