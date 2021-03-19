package solution

// assume the nums is not zero length
func partition(nums []int, p, r int) int {
	pivot := nums[r-1]
	k := p - 1
	i := p
	for i < r-1 {
		if nums[i] >= pivot {
			k = k + 1
			tmp := nums[i]
			nums[i] = nums[k]
			nums[k] = tmp
		}
		i = i + 1
	}
	tmp := nums[r-1]
	nums[r-1] = nums[k+1]
	nums[k+1] = tmp
	return k + 1
}

func FindKthLargest(nums []int, k int) int {
	length := len(nums)
	if length == 0 && k == 1 {
		return nums[0]
	}
	i := partition(nums, 0, length)
	//i + 1, len - i - 1
	if k < i+1 {
		return FindKthLargest(nums[0:i], k)
	} else if k == i+1 {
		return nums[i]
	} else {
		return FindKthLargest(nums[i+1:length], k-i-1)
	}
}
