package doorkeeper

// SortEnum represents sort enum
type SortEnum struct {
	value string
}

// GetValue returns value for SortEnum option
func (s SortEnum) GetValue() string {
	return s.value
}

// SortByPublishedAt sort by published_at
func SortByPublishedAt() SortEnum {
	return SortEnum{value: "published_at"}
}

// SortByStartsAt sort by starts_at
func SortByStartsAt() SortEnum {
	return SortEnum{value: "starts_at"}
}

// SortByUpdatedAt sort by updated_at
func SortByUpdatedAt() SortEnum {
	return SortEnum{value: "updated_at"}
}
