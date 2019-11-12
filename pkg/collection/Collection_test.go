package collection


import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue(0)
	if q.Length() != 0 || !q.IsEmpty() {
		t.Errorf("length should be 0")
	}
	val, err := q.Peek()
	if err == nil {
		t.Errorf("should be an error")
	}
	if val != nil {
		t.Errorf("value should be nil")
	}

	toAdd := []int{1, 2, 3, 4}

	for i, v := range toAdd {
		if q.Push(v) != nil {
			t.Errorf("should not return error")
		}
		if q.Length() != i+1 || q.IsEmpty() {
			t.Errorf("length should not be empty")
		}
		val, err = q.Peek()
		if err != nil {
			t.Errorf("should be no error")
		}
		if val == nil {
			t.Errorf("value should not nil")
		}
		if val.(int) != 1 {
			t.Errorf("value should should be 1 but %d", val.(int))
		}
	}

	for i := 0; i < 4; i++ {
		val, err = q.Pop()
		if err != nil {
			t.Errorf("should be no error")
		}
		if val == nil {
			t.Errorf("value should not nil")
		}
		if val.(int) != toAdd[i] {
			t.Errorf("value should should be 1 but %d", val.(int))
		}
	}

	if q.Length() != 0 || !q.IsEmpty() {
		t.Errorf("length should be 0")
	}
	val, err = q.Peek()
	if err == nil {
		t.Errorf("should be an error")
	}
	if val != nil {
		t.Errorf("value should be nil")
	}
}

func TestNewQueueWithLimit(t *testing.T) {
	q := NewQueue(2)
	if q.Length() != 0 || !q.IsEmpty() {
		t.Errorf("length should be 0")
	}
	val, err := q.Peek()
	if err == nil {
		t.Errorf("should be an error")
	}
	if val != nil {
		t.Errorf("value should be nil")
	}

	if q.Push(1) != nil {
		t.Errorf("queue should not be full")
	}
	if q.Push(2) != nil {
		t.Errorf("queue should not be full")
	}
	if q.Push(3) == nil {
		t.Errorf("queue should be full")
	}
	_, _ = q.Pop()
	if q.Push(3) != nil {
		t.Errorf("queue should not be full")
	}
}

func TestNewStack(t *testing.T) {
	s := NewStack(0)
	if s.Length() != 0 || !s.IsEmpty() {
		t.Errorf("length should be 0")
	}
	val, err := s.Peek()
	if err == nil {
		t.Errorf("should be an error")
	}
	if val != nil {
		t.Errorf("value should be nil")
	}

	toAdd := []int{1, 2, 3, 4}

	for i, v := range toAdd {
		if s.Push(v) != nil {
			t.Errorf("should not return error")
		}
		if s.Length() != i+1 || s.IsEmpty() {
			t.Errorf("length should not be empty")
		}
		val, err = s.Peek()
		if err != nil {
			t.Errorf("should be no error")
		}
		if val == nil {
			t.Errorf("value should not nil")
		}
		if val.(int) != toAdd[i] {
			t.Errorf("value should should be 1 but %d", val.(int))
		}
	}

	for i := 3; i >= 0; i-- {
		val, err = s.Pop()
		if err != nil {
			t.Errorf("should be no error")
		}
		if val == nil {
			t.Errorf("value should not nil")
		}
		if val.(int) != toAdd[i] {
			t.Errorf("value should should be %d but %d", toAdd[i], val.(int))
		}
	}

	if s.Length() != 0 || !s.IsEmpty() {
		t.Errorf("length should be 0")
	}
	val, err = s.Peek()
	if err == nil {
		t.Errorf("should be an error")
	}
	if val != nil {
		t.Errorf("value should be nil")
	}
}

