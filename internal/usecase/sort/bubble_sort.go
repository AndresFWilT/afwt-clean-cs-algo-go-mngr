package sort

type BubbleSort struct{}

func NewBubbleSortService() BubbleSort {
	return BubbleSort{}
}

func (s *BubbleSort) Sort(arrayToSort []int) (sorted []int, err error) {
	length := len(arrayToSort)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arrayToSort[i] > arrayToSort[j] {
				aux := arrayToSort[i]
				arrayToSort[i] = arrayToSort[j]
				arrayToSort[j] = aux
			}
		}
	}
	return arrayToSort, nil
}
