package pagination

import "testing"

func TestGenerate(t *testing.T) {
	page := &PageInfo{CurrentPage: 0, TotalPages: 5, Around: 3, Boundaries: 3}
	err := page.Generate()
	if err != nil {
		t.Log("Generate() PASSED, return an error if CurrentPage is equal to 0")
	} else {
		t.Error("Generate() FAILED, expected return an error if CurrentPage is equal to 0, got nil")
	}

	if err.Error() == "Current page should be great than zero" {
		t.Log("Generate() PASSED, should return the error \"Current page should be great than zero\"")
	} else {
		t.Errorf("Generate() FAILED, expected error \"Current page should be great than zero\", got %s", err.Error())
	}

	page1 := &PageInfo{CurrentPage: 10, TotalPages: 5, Around: 3, Boundaries: 3}
	err1 := page1.Generate()
	if err1 != nil {
		t.Log("Generate() PASSED, return an error if CurrentPage is greater than TotalPages")
	} else {
		t.Error("Generate() FAILED, expected return an error if CurrentPage is greater than TotalPages, got nil")
	}

	if err1.Error() == "Current page should be less than or equal to Total pages" {
		t.Log("Generate() PASSED, should return the error \"Current page should be less than or equal to Total pages\"")
	} else {
		t.Errorf("Generate() FAILED, expected error \"Current page should be less than or equal to Total pages\", got %s", err1.Error())
	}

	page2 := &PageInfo{CurrentPage: 2, TotalPages: 5, Around: 3, Boundaries: 3}
	err2 := page2.Generate()
	if err2 == nil {
		t.Log("Generate() PASSED, does not return an error if CurrentPage is less than or equal to TotalPages")
	} else {
		t.Error("Generate() FAILED, expected not to return an error if CurrentPage is less than or equal to TotalPages. got not nil")
	}

}

func TestToString(t *testing.T) {
	page := &PageInfo{CurrentPage: 11, TotalPages: 20, Around: 3, Boundaries: 3}
	page.Generate()
	expected := "1 2 3 ... 8 9 10 11 12 13 14 ... 18 19 20"
	if expected == page.ToString() {
		t.Logf("ToString() PASSED, length of links is equal to \"%s\"", expected)
	} else {
		t.Errorf("ToString() FAILED, expected \"%s\", got %s", expected, page.ToString())
	}

	page1 := &PageInfo{CurrentPage: 4, TotalPages: 5, Around: 0, Boundaries: 1}
	page1.Generate()
	expected1 := "1 ... 4 5"
	if expected1 == page1.ToString() {
		t.Logf("ToString() PASSED, length of links is equal to \"%s\"", expected1)
	} else {
		t.Errorf("ToString() FAILED, expected \"%s\", got %s", expected1, page1.ToString())
	}

	page2 := &PageInfo{CurrentPage: 4, TotalPages: 10, Around: 2, Boundaries: 2}
	page2.Generate()
	expected2 := "1 2 3 4 5 6 ... 9 10"
	if expected2 == page2.ToString() {
		t.Logf("ToString() PASSED, length of links is equal to \"%s\"", expected2)
	} else {
		t.Errorf("ToString() FAILED, expected \"%s\", got %s", expected2, page2.ToString())
	}
}

func TestAddCurrentPage(t *testing.T) {
	page := &PageInfo{CurrentPage: 3, TotalPages: 6, Around: 1, Boundaries: 1}
	page.push(1)
	page.push(2)
	page.push(3)
	page.addCurrentPage()

	if len(page.links) == 1 {
		t.Log("addCurrentPage() PASSED, length of links is equal to 1")
	} else {
		t.Errorf("addCurrentPage() FAILED, length of links expected be equal to 1, got %d", len(page.links))
	}

	if page.links[0].Value == 3 {
		t.Log("addCurrentPage() PASSED, elements are equal to [{3}]")
	} else {
		t.Errorf("addCurrentPage() FAILED, elements expected be equal to [{3}]], got %d", page.links)
	}

}

func TestAddAround(t *testing.T) {
	page := &PageInfo{CurrentPage: 3, TotalPages: 5, Around: 1, Boundaries: 0}
	page.addCurrentPage()
	page.addAround()

	if len(page.links) == 3 {
		t.Log("addAround() PASSED, length of links is equal to 3")
	} else {
		t.Errorf("addAround() FAILED, length of links expected be equal to 3, got %d", len(page.links))
	}

	if page.links[0].Value == 2 && page.links[1].Value == 3 && page.links[2].Value == 4 {
		t.Log("addAround() PASSED, elements are equal to [{2} {3} {4}]")
	} else {
		t.Errorf("addAround() FAILED, elements expected be equal to [{2} {3} {4}]], got %d", page.links)
	}

}

func TestUnshiftAround(t *testing.T) {
	page := &PageInfo{CurrentPage: 2, TotalPages: 3, Around: 1, Boundaries: 0}
	page.addCurrentPage()
	page.unshiftAround()

	if len(page.links) == 2 {
		t.Log("unshiftAround() PASSED, length of links is equal to 2")
	} else {
		t.Errorf("unshiftAround() FAILED, length of links expected be equal to 2, got %d", len(page.links))
	}

	if page.links[0].Value == 1 && page.links[1].Value == 2 {
		t.Log("unshiftAround() PASSED, elements are equal to [{1} {2}]")
	} else {
		t.Errorf("unshiftAround() FAILED, elements expected be equal to [{1} {2}], got %d", page.links)
	}
}

func TestPushAround(t *testing.T) {
	page := &PageInfo{CurrentPage: 2, TotalPages: 3, Around: 1, Boundaries: 0}
	page.addCurrentPage()
	page.pushAround()

	if len(page.links) == 2 {
		t.Log("pushAround() PASSED, length of links is equal to 2")
	} else {
		t.Errorf("pushAround() FAILED, length of links expected be equal to 2, got %d", len(page.links))
	}

	if page.links[0].Value == 2 && page.links[1].Value == 3 {
		t.Log("pushAround() PASSED, elements are equal to [2, 3]")
	} else {
		t.Errorf("pushAround() FAILED, elements expected be equal to [{2} {3}], got %d", page.links)
	}
}

func TestAddBoundary(t *testing.T) {
	page := &PageInfo{CurrentPage: 2, TotalPages: 3, Around: 0, Boundaries: 1}
	page.addCurrentPage()
	page.addBoundaries()

	if len(page.links) == 3 {
		t.Log("addBoundaries() PASSED, length of links is equal to 3")
	} else {
		t.Errorf("addBoundaries() FAILED, length of links expected be equal to 3, got %d", len(page.links))
	}

	if page.links[0].Value == 1 && page.links[1].Value == 2 && page.links[2].Value == 3 {
		t.Log("addBoundaries() PASSED, elements are equal to [1, 2, 3]")
	} else {
		t.Errorf("addBoundaries() FAILED, elements expected be equal to [1, 2, 3], got %d", page.links)
	}

	page1 := &PageInfo{CurrentPage: 10, TotalPages: 20, Around: 0, Boundaries: 2}
	page1.addCurrentPage()
	page1.addBoundaries()

	if len(page1.links) == 7 {
		t.Log("addBoundaries() PASSED, length of links is equal to 7")
	} else {
		t.Errorf("addBoundaries() FAILED, length of links expected be equal to 7, got %d", len(page1.links))
	}

	if page1.links[0].Value == 1 && page1.links[1].Value == 2 && page1.links[2].Value == 0 &&
		page1.links[3].Value == 10 && page1.links[4].Value == 0 && page1.links[5].Value == 19 &&
		page1.links[6].Value == 20 {
		t.Log("addBoundaries() PASSED, elements are equal to [1, 2, 0, 10, 0, 19, 20]")
	} else {
		t.Errorf("addBoundaries() FAILED, elements expected be equal to [1, 2, 0, 10, 0, 19, 20], got %d", page1.links)
	}

}

func TestUnshiftBoundary(t *testing.T) {
	page := &PageInfo{CurrentPage: 5, TotalPages: 10, Around: 0, Boundaries: 0}
	page.addCurrentPage()

	page.unshiftBoundary(4)
	if len(page.links) == 2 {
		t.Log("unshiftBoundary(4) PASSED, length of links is equal to 2")
	} else {
		t.Errorf("unshiftBoundary(4) FAILED, length of links expected be equal to 2, got %d", len(page.links))
	}

	if page.links[0].Value == 4 {
		t.Log("unshiftBoundary(4) PASSED, value at index 1 is equal to 4")
	} else {
		t.Errorf("unshiftBoundary(4) FAILED, value at index 1 expected be equal to 4, got %d", page.links[0].Value)
	}

	page.unshiftBoundary(2)
	if len(page.links) == 4 {
		t.Log("unshiftBoundary(2) PASSED, length of links is equal to 4")
	} else {
		t.Errorf("unshiftBoundary(2) FAILED, length of links expected be equal to 4, got %d", len(page.links))
	}

	if page.links[1].Value == 0 {
		t.Log("unshiftBoundary(2) PASSED, value at index 1 is equal to 0")
	} else {
		t.Errorf("unshiftBoundary(2) FAILED, value at index 1 expected be equal to 0, got %d", page.links[1].Value)
	}

	if page.links[0].Value == 2 {
		t.Log("unshiftBoundary(2) PASSED, value at index 0 is equal to 2")
	} else {
		t.Errorf("unshiftBoundary(2) FAILED, value at index 0 expected be equal to 8, got %d", page.links[0].Value)
	}
}

func TestPushBoundary(t *testing.T) {
	page := &PageInfo{CurrentPage: 5, TotalPages: 10, Around: 0, Boundaries: 0}
	page.addCurrentPage()

	page.pushBoundary(6)
	if len(page.links) == 2 {
		t.Log("pushBoundary(6) PASSED, length of links is equal to 2")
	} else {
		t.Errorf("pushBoundary(6) FAILED, length of links expected be equal to 2, got %d", len(page.links))
	}

	if page.links[1].Value == 6 {
		t.Log("pushBoundary(6) PASSED, value at index 1 is equal to 6")
	} else {
		t.Errorf("pushBoundary(6) FAILED, value at index 1 expected be equal to 6, got %d", page.links[1].Value)
	}

	page.pushBoundary(8)
	if len(page.links) == 4 {
		t.Log("pushBoundary(8) PASSED, length of links is equal to 4")
	} else {
		t.Errorf("pushBoundary(8) FAILED, length of links expected be equal to 4, got %d", len(page.links))
	}

	if page.links[2].Value == 0 {
		t.Log("pushBoundary(8) PASSED, value at index 2 is equal to 0")
	} else {
		t.Errorf("pushBoundary(8) FAILED, value at index 2 expected be equal to 0, got %d", page.links[2].Value)
	}

	if page.links[3].Value == 8 {
		t.Log("pushBoundary(8) PASSED, value at index 3 is equal to 8")
	} else {
		t.Errorf("pushBoundary(8) FAILED, value at index 3 expected be equal to 8, got %d", page.links[3].Value)
	}
}

func TestPush(t *testing.T) {
	page := &PageInfo{CurrentPage: 4, TotalPages: 10, Around: 0, Boundaries: 0}
	page.push(4)

	if page.links[0].Value == 4 {
		t.Log("push(4) PASSED, value at index 0 is equal to 4")
	} else {
		t.Errorf("push(4) FAILED, value at index 0 expected be equal to 4, got %d", page.links[0].Value)
	}

	if len(page.links) == 1 {
		t.Log("push(4) PASSED, length of links is equal to 1")
	} else {
		t.Errorf("push(4) FAILED, length of links expected be equal to 1, got %d", len(page.links))
	}
}

func TestUnshift(t *testing.T) {
	page := &PageInfo{CurrentPage: 0, TotalPages: 10, Around: 0, Boundaries: 0}
	page.unshift(4)

	if page.links[0].Value == 4 {
		t.Log("push(4) PASSED, value at index 0 is equal to 4")
	} else {
		t.Errorf("push(4) FAILED, value at index 0 expected be equal to 4, got %d", page.links[0].Value)
	}

	if len(page.links) == 1 {
		t.Log("push(4) PASSED, length of links is equal to 1")
	} else {
		t.Errorf("push(4) FAILED, length of links expected be equal to 1, got %d", len(page.links))
	}

	page.unshift(1)
	if page.links[0].Value == 1 {
		t.Log("push(4) PASSED, value at index 0 is equal to 1")
	} else {
		t.Errorf("push(4) FAILED, value at index 0 expected be equal to 1, got %d", page.links[0].Value)
	}

	if len(page.links) == 2 {
		t.Log("push(4) PASSED, length of links is equal to 2")
	} else {
		t.Errorf("push(4) FAILED, length of links expected be equal to 2, got %d", len(page.links))
	}
}

func TestInsertAt(t *testing.T) {
	page := &PageInfo{CurrentPage: 0, TotalPages: 10, Around: 0, Boundaries: 0}
	page.insertAt(0, 1)

	if page.links[0].Value == 1 {
		t.Log("insertAt(0, 1) PASSED, element at index 0 is equal to 1")
	} else {
		t.Errorf("insertAt(0, 1) FAILED, expected be equal to 1, got %d", page.links[0].Value)
	}

	page.insertAt(1, 2)
	if page.links[1].Value == 2 {
		t.Log("insertAt(1, 2) PASSED, element at index 1 is equal to 2")
	} else {
		t.Errorf("insertAt(1, 2) FAILED, expected be equal to 2, got %d", page.links[1].Value)
	}

	page.insertAt(2, 3)
	if page.links[2].Value == 3 {
		t.Log("insertAt(2, 3) PASSED, element at index 2 is equal to 3")
	} else {
		t.Errorf("insertAt(2, 3) FAILED, expected be equal to 3, got %d", page.links[1].Value)
	}

	page.insertAt(1, 4)
	if page.links[1].Value == 4 {
		t.Log("insertAt(1, 4) PASSED, element at index 1 is equal to 4")
	} else {
		t.Errorf("insertAt(1, 4) FAILED, expected be equal to 4, got %d", page.links[1].Value)
	}

	err := page.insertAt(10, 5)
	if err != nil {
		t.Log("insertAt(10, 5) PASSED, could not insert element out of bounds")
	} else {
		t.Error("insertAt(10, 5) FAILED, expected returned error be not nil, got nil")
	}

}
