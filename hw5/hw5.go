package main

func Map(data []int, action func(int) int) []int {
	if data == nil || action == nil {
		return []int(nil)
	}
	if len(data) == 0 {
		return []int{}
	}

	result := make([]int, len(data))
	for i, v := range data {
		result[i] = action(v)
	}

	return result
}

func Filter(data []int, action func(int) bool) []int {
	if data == nil || action == nil {
		return []int(nil)
	}
	if len(data) == 0 {
		return []int{}
	}

	var result []int
	for _, v := range data {
		if action(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	if data == nil || action == nil {
		return initial
	}

	result := initial

	for _, v := range data {
		result = action(result, v)
	}

	return result
}
