package encode

import (
	"bytes"
	"image"
	"image/png"
	"testing"
)

func TestEscape(t *testing.T) {
	esc := `"foo;bar\baz"`
	exp := `\"foo\;bar\\baz\"`
	got := Escape(esc)
	if got != exp {
		t.Error("Expected:", exp, "got:", got)
	}
}

func TestEncodeWpaReturnsSomething(t *testing.T) {
	data := EncodeWpa("foo", "bar")
	if len(data) == 0 {
		t.Error("Didn't get any encoded data")
	}
}

func TestEncodeWpaIsPng(t *testing.T) {
	data := EncodeWpa("foo", "bar")
	buf := bytes.NewBuffer(data)
	img, err := png.Decode(buf)
	if err != nil {
		t.Error("Could not decode png:", err)
	}
	expBounds := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{256, 256},
	}
	if img.Bounds() != expBounds {
		t.Error("Unexpected size, got:", img.Bounds(), "expected:", expBounds)
	}

}
