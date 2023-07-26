package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	m := M{}
	result := m.Add([]int{1, 2, 3}) // slice (array)
	assert.Equal(t, 6, result)

	a := M{}
	result = a.Add([]int{1, 2}) // slice (array)
	assert.Equal(t, 3, result)

}

func TestSub(t *testing.T) {
	m := M{}
	result := m.Sub(3, 1)
	assert.Equal(t, 2, result)
}

func TestDupes(t*testing.T) {
	m := M{}
	result := m.Dupes([]int{1, 2, 3, 1, 2})
	assert.Equal(t, []int{1, 2,}, result)
}
// todo: add test for Concat

func TestConcat(t *testing.T) {

	m := M{}
	result := m.Concat([]string{"hello", "world"})
	assert.Equal(t, "hello world", result)
}
