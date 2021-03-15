package main

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, v := range nums {
		j, ok := hashmap[target-v]

		if ok {
			return []int{i, j}
		}

		hashmap[v] = i
	}

	return []int{-1, -1}
}
