package twosum

func TwoSum(nums []int, target int) []int {

	output := make(map[int]int)

	for k, v := range nums {
		if curindex, isFound := output[target-v]; isFound {
			return []int{curindex, k}
		}
		output[v] = k
	}
	return []int{}
}
