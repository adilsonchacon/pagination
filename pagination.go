package pagination

import (
	"errors"
	"fmt"
	"strings"
)

// PageInterface is the interface to interact with th pagination package
type PageInterface interface {
	Generate()
	ToString()
}

// PageInfo is the attributes used as input for pagination proccess
type PageInfo struct {
	CurrentPage int
	TotalPages  int
	Boundaries  int
	Around      int
	links       []Link
}

// Generate generates the pagination, however returns an error if an invalid value is given
func (page *PageInfo) Generate() error {
	if page.CurrentPage < 1 {
		return errors.New("Current page should be great than zero")
	} else if page.CurrentPage > page.TotalPages {
		return errors.New("Current page should be less than or equal to Total pages")
	}

	page.addCurrentPage()
	page.addAround()
	page.addBoundaries()

	return nil
}

// ToString returns a string representating the pagination
func (page *PageInfo) ToString() string {
	var stringLinks []string
	for _, value := range page.links {
		stringLinks = append(stringLinks, value.ToString())
	}

	return strings.Join(stringLinks, " ")
}

func (page *PageInfo) addCurrentPage() {
	page.links = nil
	page.push(page.CurrentPage)
}

func (page *PageInfo) addAround() {
	for i := 0; i < page.Around; i++ {
		if page.TotalPages > len(page.links) {
			page.unshiftAround()
			page.pushAround()
		}
	}
}

func (page *PageInfo) unshiftAround() {
	firstValue := page.links[0].Value

	if firstValue > 1 {
		page.unshift(firstValue - 1)
	}
}

func (page *PageInfo) pushAround() {
	lastValue := page.links[len(page.links)-1].Value

	if lastValue < page.TotalPages {
		page.push(lastValue + 1)
	}
}

func (page *PageInfo) addBoundaries() {
	for i := 0; i < page.Boundaries; i++ {
		if page.TotalPages > len(page.links) {
			page.unshiftBoundary(page.Boundaries - i)
			page.pushBoundary(page.TotalPages - (page.Boundaries - (i + 1)))
		}
	}
}

func (page *PageInfo) unshiftBoundary(value int) {
	firstValue := page.links[0].Value
	gap := firstValue - value

	// if gap between numbers is greater than or equal to 2 then add ellipsis
	if gap >= 2 {
		page.unshift(0)
	}

	if gap > 0 {
		page.unshift(value)
	}

}

func (page *PageInfo) pushBoundary(value int) {
	lastValue := page.links[len(page.links)-1].Value
	gap := value - lastValue

	// if gap between numbers is greater than or equal to 2 then add ellipsis
	if gap >= 2 {
		page.push(0)
	}

	if gap > 0 {
		page.push(value)
	}
}

func (page *PageInfo) unshift(value int) {
	page.insertAt(0, value)
}

func (page *PageInfo) push(value int) {
	page.links = append(page.links, Link{Value: value})
}

func (page *PageInfo) insertAt(index, value int) error {
	if len(page.links) == 0 || len(page.links) == index {
		page.push(value)
	} else if index < len(page.links) {
		page.links = append(page.links[:index+1], page.links[index:]...)
		page.links[index] = Link{Value: value}
	} else {
		return errors.New(fmt.Sprintf("index out of range [%d] length is %d", index, len(page.links)))
	}

	return nil
}
