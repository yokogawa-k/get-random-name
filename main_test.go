package main 

import (
	"strings"
	"testing"
)

func TestNameFormat(t *testing.T) {
	name := GetRandomNameWithHyphen(0)
	if !strings.Contains(name, "-") {
		t.Fatalf("Generated name does not contain an hyphen")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func TestNameRetries(t *testing.T) {
	name := GetRandomNameWithHyphen(1)
	if !strings.Contains(name, "-") {
		t.Fatalf("Generated name does not contain an hyphen")
	}
	if !strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name doesn't contain a number")
	}

}
