package path

import (
	"testing"
)

func TestDefaultTopInfo(t *testing.T) {
	info := NewInfo(DefaultTop)
	shouldBeTop(info, t)
	checkName(info, t)
}

func TestTopInfo(t *testing.T) {
	info := NewInfo("\\")
	shouldBeTop(info, t)
	checkName(info, t)
}

func TestInfoCustomSeperator(t *testing.T) {
	info := NewInfoCustomSeperator('/', "/")
	shouldBeTop(info, t)
	checkName(info, t)
}

func TestInfoCustomTop(t *testing.T) {
	info := NewInfoCustomTop("\\TOP", "\\TOP")
	shouldBeTop(info, t)
	checkName(info, t)
}

func TestInfoCustomTopWithSlash(t *testing.T) {
	info := NewInfoCustomTop("\\TOP", "\\TOP\\")
	shouldBeTop(info, t)
	checkName(info, t)
}

func TestFirstChildInfo(t *testing.T) {

	info := NewInfo("\\TOP\\Child")
	shouldNotBeTop(info, t)
	checkName(info, t)

}
func TestFirstChildInfoWithSlash(t *testing.T) {

	info := NewInfo("\\TOP\\Child\\")

	shouldNotBeTop(info, t)
	checkName(info, t)

}

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

func checkName(info *Info, t *testing.T) {
	if len(info.Name) == 0 {
		t.Errorf("%s does not have a name set", info.String())
		PrintDetails(info)
	}
}
