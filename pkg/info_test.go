package path

import (
	"testing"
)

func TestDefaultTopInfo(t *testing.T) {
	info := NewInfo(DefaultTop)
	PrintDetails(info)
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
	}
}

func TestTopInfo(t *testing.T) {
	info := NewInfo("\\")
	PrintDetails(info)
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
	}
}

func TestNonTopInfo(t *testing.T) {

	info := NewInfo("\\ROOT\\")
	PrintDetails(info)
	if info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
	}
}
