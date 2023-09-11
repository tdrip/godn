package path

import (
	"testing"
)

// TEST TOP
func TestTopDefault(t *testing.T) {
	info := NewInfo(DefaultTop)
	shouldBeTop(info, t)
	hasNoName(info, t)
	shouldNotHaveParent(info, t)
}

func TestTop(t *testing.T) {
	info := NewInfo("\\")
	shouldBeTop(info, t)
	hasNoName(info, t)
	shouldNotHaveParent(info, t)
}

func TestTopCustomSeperator(t *testing.T) {
	info := NewInfoCustomSeperator('/', "/")
	shouldBeTop(info, t)
	hasNoName(info, t)
	shouldNotHaveParent(info, t)
}

func TestTopCustom(t *testing.T) {
	info := NewInfoCustomTop("\\TOP", "\\TOP")
	shouldBeTop(info, t)
	hasName(info, t)
	shouldNotHaveParent(info, t)
}

func TestTopCustomWithSlash(t *testing.T) {
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

/// TEST SEPERATORS

func TestFirstChildCustomSeperator(t *testing.T) {

	info := NewInfoCustomSeperator('/', "/TOP/Child")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)
}

func TestFirstChildCustomSeperatorExtra(t *testing.T) {
	info := NewInfoCustomSeperator('/', "/TOP/Child/")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)

}

func TestSecondChildCustomSeperatorExtra(t *testing.T) {
	info := NewInfoCustomSeperator('/', "/TOP/Child/NewChild/")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)

}
func TestSecondChildCustomSeperator(t *testing.T) {
	info := NewInfoCustomSeperator('/', "/TOP/Child/NewChild")
	shouldNotBeTop(info, t)
	hasName(info, t)
	shouldHaveParent(info, t)

}

func TestChildEqualsChild(t *testing.T) {

	child := NewInfo("\\TOP\\Child\\")
	child1 := NewInfo("\\TOP\\Child\\")

	if !child.Equals(child1) {
		t.Errorf("%s should equal %s", child.String(), child1.String())
	}
}

func TestChildEqualsChildNoSlash(t *testing.T) {

	child := NewInfo("\\TOP\\Child\\")
	child1 := NewInfo("\\TOP\\Child")

	if !child.Equals(child1) {
		t.Errorf("%s should equal %s", child.String(), child1.String())
	}
}

func TestChildEqualsChildString(t *testing.T) {

	child := NewInfo("\\TOP\\Child\\")
	child1 := "\\TOP\\Child\\"

	if !child.StringEquals(child1) {
		t.Errorf("%s should equal %s", child.String(), child1)
	}
}

func TestChildEqualsChildStringNoEnding(t *testing.T) {

	child := NewInfo("\\TOP\\Child\\")
	child1 := "\\TOP\\Child"

	if !child.StringEquals(child1) {
		t.Errorf("%s should equal %s", child.String(), child1)
	}
}

func TestNavigateUp(t *testing.T) {
	child := NewInfoCustomSeperator('/', "/parent/child")
	parent := child.Parent
	pwd := NewInfoCustom("/", '/', parent.ParsedPath)

	if !parent.Equals(pwd) {
		t.Errorf("%s should equal %s", parent.String(), pwd)
	}
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
