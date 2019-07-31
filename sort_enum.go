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
var SortByPublishedAt = SortEnum{value: "published_at"}

// SortByStartsAt sort by starts_at
var SortByStartsAt = SortEnum{value: "starts_at"}

// SortByUpdatedAt sort by updated_at
var SortByUpdatedAt = SortEnum{value: "updated_at"}
