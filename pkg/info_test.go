package path

import (
	"testing"
)

func TestDefaultTopInfo(t *testing.T) {
	info := NewInfo(DefaultTop)
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
		PrintDetails(info)
	}
}

func TestTopInfo(t *testing.T) {
	info := NewInfo("\\")
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
		PrintDetails(info)
	}
}

func TestInfoCustomSeperator(t *testing.T) {
	info := NewInfoCustomSeperator('/', "/")
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
		PrintDetails(info)
	}
}

func TestInfoCustomTop(t *testing.T) {
	info := NewInfoCustomTop("\\ROOT", "\\ROOT")
	if !info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
		PrintDetails(info)
	}
}

func TestNonTopInfo(t *testing.T) {

	info := NewInfo("\\ROOT\\")
	if info.IsTop() {
		t.Errorf("%s has IsTop() set to false - should be true", info.String())
		PrintDetails(info)
	}
}
