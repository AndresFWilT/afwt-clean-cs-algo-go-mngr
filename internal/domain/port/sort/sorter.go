package sort

type Sorter interface {
	Sort(arrayToSort []int) (sorted []int, err error)
}
