package sort

func quickSort(nums []int, left, right int) {
	if left < right {
		idx := partition(nums, left, right)
		quickSort(nums, left, idx)
		quickSort(nums, idx+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left+(right-left)/2]
	i, j := left-1, right+1
	for {
		i++
		for nums[i] < pivot {
			i++
		}

		j--
		for nums[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}
