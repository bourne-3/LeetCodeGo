package string

import (
	"math"
	"sort"
)

func strStr(haystack, needle string) int {
	var j int
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := -1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}
	var ret int
	dvd, dvs := abs(dividend), abs(divisor)
	for dvd >= dvs {
		tDvd, tRet := dvs, 1
		for tDvd<<1 <= dvd {
			tDvd <<= 1
			tRet <<= 1
		}
		dvd -= tDvd
		ret += tRet
	}

	return sign * ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var ret []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			ret = append(ret, nums1[i])
			i++
			j++
		}
	}
	return ret
}
