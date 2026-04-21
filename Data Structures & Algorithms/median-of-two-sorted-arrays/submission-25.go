func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	total := n1 + n2
	half := (total + 1) / 2
	low, high := 0, n1

	for low <= high {
		p := low + (high - low) / 2
		j := half - p
		l1, l2 := math.MinInt, math.MinInt
		r1, r2 := math.MaxInt, math.MaxInt

		if p > 0 {
			l1 = nums1[p-1]
		}
		if p < n1 {
			r1 = nums1[p]
		}
		if j > 0 {
			l2 = nums2[j-1]
		}
		if j < n2 {
			r2 = nums2[j]
		}

		if l1 <= r2 && l2 <= r1 {
			if total % 2 == 0 {
				// even
				leftMax := float64(max(l1, l2))
				rightMin := float64(min(r1, r2))

				return leftMax + (rightMin - leftMax) / float64(2)
			} else {
				// odd
				leftMax := float64(max(l1, l2))
				return leftMax
			}
		} else if l1 > r2 {
			high = p - 1
		} else {
			low = p + 1
		}
	}

	return float64(-1)
}
