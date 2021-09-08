package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n
	for m > 0 && n > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
		p--
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}

func main() {

}
