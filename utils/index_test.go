package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	idx := make(Index)

	assert.Nil(t, idx.Search("foo"))
	assert.Nil(t, idx.Search("donut"))

	idx.Add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	assert.Nil(t, idx.Search("a"))
	assert.Equal(t, idx.Search("donut"), []int{1})
	assert.Equal(t, idx.Search("DoNuts"), []int{1})
	assert.Equal(t, idx.Search("glass"), []int{1})

	idx.Add([]document{{ID: 2, Text: "donut is a donut"}})
	assert.Nil(t, idx.Search("a"))
	assert.Equal(t, idx.Search("donut"), []int{1, 2})
	assert.Equal(t, idx.Search("DoNuts"), []int{1, 2})
	assert.Equal(t, idx.Search("glass"), []int{1})
}
