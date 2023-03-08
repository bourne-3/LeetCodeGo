package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	inSample := []int{5, 31, 5, 100, 1}
	quickSort(inSample, 0, len(inSample)-1)
	fmt.Println(inSample)
}
