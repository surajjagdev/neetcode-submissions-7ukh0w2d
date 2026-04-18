func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// need to find the partition from nums1 array and pick remaining
	// element from the second array

	// using this partition we have a left idx, before p, and then p 
	// can be before or after the idx found in nums2

	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	total := (n1 + n2)
	// max number of elements if both arrays are combined
	half := (total) / 2

	low, high := 0, n1

	for low <= high {
		partition := low + (high - low) / 2
		j := half - partition


		l1, l2 := math.MinInt, math.MinInt
		r1, r2 := math.MaxInt, math.MaxInt


		if partition > 0 {
			l1 = nums1[partition-1]
		}
		if partition < n1 {
			r1 = nums1[partition]
		}
		if j > 0 {
			l2 = nums2[j-1]
		}
		if j < n2 {
			r2 = nums2[j]
		}

		// is this a correct partition?
		if l1 <= r2 && l2 <= r1 {
			rightMin := float64(min(r1, r2))

			// is this even or odd?
			if total % 2 != 0 {
				// odd
				return rightMin
			} else {
				// even
				leftMax := max(l1, l2)

				return (float64(leftMax) + rightMin) / float64(2)
			}
		} else if l1 > r2 {
			high = partition - 1
		} else {
			low = partition + 1
		}
	}

	return float64(-1)
}
