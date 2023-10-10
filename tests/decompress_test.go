package tests

import (
	"testing"

	dc "github.com/bersen66/decompress"
	"github.com/stretchr/testify/assert"
)

func TestDecompressFromCondition1(t *testing.T) {
	result, err := dc.Decompress("a4bc2d5e")
	assert.Equal(t, err, nil)
	assert.Equal(t, "aaaabccddddde", result)
}

func TestDecompressFromCondition2(t *testing.T) {
	result, err := dc.Decompress("abcd")
	assert.Equal(t, err, nil)
	assert.Equal(t, "abcd", result)
}

func TestDecompressFromCondition3(t *testing.T) {
	_, err := dc.Decompress("45")
	assert.NotEqual(t, nil, err)
}

func TestDecompressFromCondition4(t *testing.T) {
	result, err := dc.Decompress("")
	assert.Equal(t, nil, err)
	assert.Equal(t, "", result)
}

func TestEscape1(t *testing.T) {
	result, err := dc.Decompress("qwe\\4\\5")
	assert.Equal(t, err, nil)
	assert.Equal(t, "qwe45", result)
}

func TestEscape2(t *testing.T) {
	result, err := dc.Decompress("qwe\\45")
	assert.Equal(t, err, nil)
	assert.Equal(t, "qwe44444", result)
}

func TestEscape3(t *testing.T) {
	result, err := dc.Decompress("qwe\\\\5")
	assert.Equal(t, err, nil)
	assert.Equal(t, "qwe\\\\\\\\\\", result)
}
