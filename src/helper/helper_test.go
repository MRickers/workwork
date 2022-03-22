package helper

import (
	"testing"
)

func TestStripExeName1(t *testing.T) {
	var path = "C:\\Path\\to\\test.exe"

	stripped := StripExeName(path)

	if stripped != "C:/Path/to/" {
		t.Fatalf("wrong exe path %s", stripped)
	}
}

func TestStripExeName2(t *testing.T) {
	var path = "C:/Path/to/your/exe/test.exe"

	stripped := StripExeName(path)

	if stripped != "C:/Path/to/your/exe/" {
		t.Fatalf("wrong exe path %s", stripped)
	}
}

func TestStripExeName3(t *testing.T) {
	var path = "test.exe"

	stripped := StripExeName(path)

	if stripped != "" {
		t.Fatalf("wrong exe path %s", stripped)
	}
}
