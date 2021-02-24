package pagination

import "testing"

func TestLink(t *testing.T) {
	link1 := Link{Value: 0}
	if link1.ToString() == "..." {
		t.Log("ToString() PASSED, returns an ellipsis if values is 0")
	} else {
		t.Errorf("ToString() FAILED, expected an ellipsis, got %s", link1.ToString())
	}

	link2 := Link{Value: 10}
	if link2.ToString() == "10" {
		t.Log("ToString() PASSED, returns 10")
	} else {
		t.Errorf("ToString() FAILED, expected \"10\", got \"%s\"", link2.ToString())
	}
}
