package path

import (
	"testing"
)

func TestDefaultTopInfo(t *testing.T) {
	info := NewInfo(DefaultTop)
	shouldBeTop(info, t)
	hasName(info, t)
	shouldNotHaveParent(info, t)
}

func TestTopInfo(t *testing.T) {
	info := NewInfo("\\")
	shouldBeTop(info, t)
	hasNoName(info, t)
	shouldNotHaveParent(info, t)
}

func TestInfoCustomSeperator(t *testing.T) {
	info := NewInfoCustomSeperator('/', "/")
	shouldBeTop(info, t)
	hasNoName(info, t)
	shouldNotHaveParent(info, t)
}

func TestInfoCustomTop(t *testing.T) {
	info := NewInfoCustomTop("\\TOP", "\\TOP")
	shouldBeTop(info, t)
	hasName(info, t)
	shouldNotHaveParent(info, t)
}

func TestInfoCustomTopWithSlash(t *testing.T) {
	info := NewInfoCustomTop("\\TOP", "\\TOP\\")
	shouldBeTop(info, t)
	hasName(info, t)
	shouldNotHaveParent(info, t)
}

/// CHILD TESTS FROM HERE FORWARD ////

func TestFirstChildInfo(t *testing.T) {

	info := NewInfo("\\TOP\\Child")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)
}

func TestFirstChildInfoWithSlash(t *testing.T) {
	info := NewInfo("\\TOP\\Child\\")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)

}

func TestSecondChildInfoWithSlash(t *testing.T) {
	info := NewInfo("\\TOP\\Child\\NewChild\\")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)

}
func TestSecondChildInfo(t *testing.T) {
	info := NewInfo("\\TOP\\Child\\NewChild")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)

}

// common functions
func shouldNotBeTop(info *Info, t *testing.T) {
	if info.IsTop() {
		t.Errorf("%s has IsTop() set to true - should be false", info.String())
		PrintDetails(info)
	}
}

func shouldBeTop(info *Info, t *testing.T) {
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
		PrintDetails(info)
	}
}

func hasName(info *Info, t *testing.T) {
	if len(info.Name) == 0 {
		t.Errorf("%s does not have a name set", info.String())
		PrintDetails(info)
	}
}

func hasNoName(info *Info, t *testing.T) {
	if len(info.Name) > 0 {
		t.Errorf("%s does  have a name set when it should not", info.String())
		PrintDetails(info)
	}
}

func shouldHaveParent(info *Info, t *testing.T) {
	if info.Parent == nil {
		t.Errorf("%s does not have a parent when it should", info.String())
		PrintDetails(info)
	}
}

func shouldNotHaveParent(info *Info, t *testing.T) {
	if info.Parent != nil {
		t.Errorf("%s does have a parent when it should not", info.String())
		PrintDetails(info)
	}
}
