package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, i := len(nums1), 0
	len2, j := len(nums2), 0
	lenR, result := len1+len2, make([]int, len1+len2)

	for k := 0; k < len1+len2; k++ {
		if i == len1 || (i < len1 && j < len2 && nums1[i] > nums2[j]) {
			result[k] = nums2[j]
			j++
		} else {
			result[k] = nums1[i]
			i++
		}
	}

	if lenR%2 == 0 {
		return float64(result[lenR/2]+result[lenR/2-1]) / 2
	} else {
		return float64(result[lenR/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{78, 79, 91, 103, 115, 127, 139, 151, 163, 175, 187, 199, 211, 223, 235, 247, 259, 271, 283, 295, 307, 319, 331, 343, 355, 367, 379, 391, 403, 415, 427, 439, 451, 463, 475, 487, 499, 511, 523, 535, 547, 559, 571, 583, 595, 607, 619, 631, 643, 655, 667, 679, 691, 703, 715, 727, 739, 751, 763, 775, 787, 799, 811, 823, 835, 847, 859, 871, 883, 895, 907, 919, 931, 943, 955, 967, 979, 991, 1003, 1015, 1027, 1039, 1051, 1063, 1075, 1087, 1099, 1111}, []int{140, 161, 182, 203, 224, 245, 266, 287, 308, 329, 350, 371, 392, 413, 434, 455, 476, 497, 518, 539, 560, 581, 602, 623, 644, 665, 686, 707, 728, 749, 770, 791, 812, 833, 854, 875, 896, 917, 938, 959, 980, 1001, 1022, 1043, 1064, 1085, 1106, 1127, 1148, 1169, 1190, 1211, 1232, 1253, 1274, 1295, 1316, 1337, 1358, 1379, 1400, 1421, 1442, 1463, 1484, 1505, 1526, 1547, 1562, 1568, 1589, 1610, 1631, 1652, 1673, 1694, 1715, 1736, 1757, 1778, 1799, 1820, 1841, 1862, 1883, 1904, 1925, 1946, 19670}))
}
