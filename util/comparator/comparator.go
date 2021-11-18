package comparator

type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	first := a.(int)
	second := b.(int)

	if first < second {
		return -1
	} else if first > second {
		return 1
	}
	return 0
}

func StringComparator(a, b interface{}) int {
	first := a.(string)
	second := b.(string)

	if first < second {
		return -1
	} else if first > second {
		return 1
	}
	return 0
}
