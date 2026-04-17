func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// find correct partition
	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	half := (n1 + n2 + 1) / 2

	low, high := 0, n1


	for low <= high {
		mid := low + ((high - low) / 2)
		remainder := half - mid

		// elements to take from left and right
		l1, l2 := math.MinInt, math.MinInt
		r1, r2 := math.MaxInt, math.MaxInt

		if mid > 0 {
			l1 = nums1[mid - 1]
		}
		if mid < n1 {
			r1 = nums1[mid]
		}
		if remainder > 0 {
			l2 = nums2[remainder - 1]
		}
		if remainder < n2 {
			r2 = nums2[remainder]
		}


		// lets pick the 2 numbers to create median
		if l1 <= r2 && l2 <= r1{
			leftMax := max(l1, l2)

			// if its odd, return max between the lefts
			if (n1 + n2) % 2 != 0 {
				return float64(leftMax)
			} else {
				// need left + right / 2
				rightMin := min(r1, r2)

				return float64(leftMax + rightMin) / float64(2)
			}
		} else if l1 > r2 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

