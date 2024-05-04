package util

func InArray[K comparable](str K, array []K) bool {
	for _, elem := range array {
		if str == elem {
			return true
		}
	}

	return false
}

func InSlice[T comparable](check T, arr []T) bool {
	for _, item := range arr {
		if item == check {
			return true
		}
	}
	return false
}

func RemoveDuplicate[T comparable](array []T) []T {
	m := map[T]bool{}

	for _, element := range array {
		m[element] = true
	}

	results := []T{}
	for element := range m {
		results = append(results, element)
	}

	return results
}

func IsElementDuplicate[T comparable](array []T) bool {
	m := map[T]bool{}

	for _, element := range array {
		m[element] = true
	}

	return len(array) != len(m)
}

func ArrayIntersect(array1 []string, array2 []string) bool {
	mapper := map[string]int{}
	array1 = append(array1, array2...)
	for _, item := range array1 {
		if _, ok := mapper[item]; ok {
			return true
		} else {
			mapper[item] = 1
		}
	}

	return false
}
