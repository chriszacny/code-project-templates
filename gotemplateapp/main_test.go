package main

import "testing"

func TestEdgeOneChar(t *testing.T) {
	in := "A"
	expect := "A"
	out := convert(in, 1)
	if (out != expect) {
		t.Fatalf("expected: %s, got: %s", expect, out)
	}
}

func TestEdgeOneRow(t *testing.T) {
	in := "ABCD"
	expect := "ABCD"
	out := convert(in, 1)
	if (out != expect) {
		t.Fatalf("expected: %s, got: %s", expect, out)
	}
}

func TestBasicA(t *testing.T) {
	in := "PAYPALISHIRING"
	expect := "PAHNAPLSIIGYIR"
	out := convert(in, 3)
	if (out != expect) {
		t.Fatalf("expected: %s, got: %s", expect, out)
	}
}

func TestBasicB(t *testing.T) {
	in := "PAYPALISHIRING"
	expect := "PINALSIGYAHRPI"
	out := convert(in, 4)
	if (out != expect) {
		t.Fatalf("expected: %s, got: %s", expect, out)
	}
}