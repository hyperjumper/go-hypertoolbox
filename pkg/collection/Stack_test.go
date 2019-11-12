package collection

import (
	assert "github.com/hyperjumper/go-hypertoolbox/pkg/testing"
	"testing"
)

func TestStack_Clear(t *testing.T) {
	s := Stack{}
	assert.AssertTrue(t, s.Size() == 0)
	s.Push("One")
	s.Push("Two")
	s.Push("Three")
	s.Push("Four")
	assert.AssertTrue(t, s.Size() == 4)
	s.Clear()
	assert.AssertTrue(t, s.Size() == 0)
}

func TestStack_Size(t *testing.T) {
	s := Stack{}
	assert.AssertTrue(t, s.Size() == 0)
	s.Push("One")
	assert.AssertTrue(t, s.Size() == 1)
	s.Push("Two")
	assert.AssertTrue(t, s.Size() == 2)
	s.Push("Three")
	assert.AssertTrue(t, s.Size() == 3)
	s.Push("Four")
	assert.AssertTrue(t, s.Size() == 4)
}

func TestStack_Peek(t *testing.T) {
	s := Stack{}
	assert.AssertNil(t, s.Peek())
	s.Push("One")
	assert.AssertEquals(t, "One", s.Peek())
	s.Push("Two")
	assert.AssertEquals(t, "Two", s.Peek())
	s.Push("Three")
	assert.AssertEquals(t, "Three", s.Peek())
	s.Push("Four")
	assert.AssertEquals(t, "Four", s.Peek())
}

func TestStack_Pop(t *testing.T) {
	s := Stack{}
	assert.AssertNil(t, s.Pop())
	s.Push("One")
	s.Push("Two")
	s.Push("Three")
	s.Push("Four")
	assert.AssertEquals(t, "Four", s.Pop())
	assert.AssertEquals(t, "Three", s.Pop())
	assert.AssertEquals(t, "Two", s.Pop())
	assert.AssertEquals(t, "One", s.Pop())
	assert.AssertNil(t, s.Pop())
}
