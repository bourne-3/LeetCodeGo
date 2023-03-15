package sort

func merSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := merSort(nums[:mid])
	right := merSort(nums[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	ret := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ret[k] = left[i]
			i++
		} else {
			ret[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		ret[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		ret[k] = right[j]
		j++
		k++
	}

	return ret
}
