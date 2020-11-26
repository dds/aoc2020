package lib

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Reduce(f func(int, int) int, nums ...int) (r int) {
	r = f(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		r = f(r, nums[i])
	}
	return
}

func Apply(f func(int) int, nums ...int) (r []int) {
	if len(nums) < 1 {
		return
	}
	r = make([]int, len(nums))
	copy(r, nums)
	for i := 0; i < len(r); i++ {
		r[i] = f(r[i])
	}
	return
}

func Sign(a int) int {
	if a >= 0 {
		return 1
	}
	return -1
}
