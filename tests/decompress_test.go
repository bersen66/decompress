package tests

import (
	"testing"

	dc "github.com/bersen66/decompress"
)

func TestDecompress(t *testing.T) {
	dc.Decompress("It works")
}