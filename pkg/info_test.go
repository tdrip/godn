package path

import (
	"testing"
)

func TestTopInfo(t *testing.T) {
	info := NewInfo(DefaultTop)
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
	}
}

func TestNonTopInfo(t *testing.T) {

	info := NewInfo("\\ROOT\\")
	if info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
	}
}
