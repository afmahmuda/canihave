package canihave

func ValueOfOrDefault[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

func ValueOfOrEmpty[T any](ptr *T) T {
	var defVal T
	return ValueOfOrDefault(ptr, defVal)
}

func PointerOf[T any](v T) *T {
	return &v
}
