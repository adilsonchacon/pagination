package pagination

import "strconv"

// LinkInterface is
type LinkInterface interface {
	ToString()
}

// Link is
type Link struct {
	Value int
}

// ToString is
func (link *Link) ToString() string {
	if link.Value == 0 {
		return "..."
	}

	return strconv.Itoa(link.Value)
}
