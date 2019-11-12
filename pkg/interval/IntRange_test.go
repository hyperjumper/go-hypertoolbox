package interval

import "testing"

func TestIntRange_CanMerge(t *testing.T) {
	ts := make(map[*IntRange]bool)
	ts[NewIntRange(2,5)] = false
	ts[NewIntRange(2,8)] = false
	ts[NewIntRange(2,9)] = true
	ts[NewIntRange(2,10)] = true
	ts[NewIntRange(8,15)] = true
	ts[NewIntRange(9,15)] = true
	ts[NewIntRange(10,15)] = true
	ts[NewIntRange(15,19)] = true
	ts[NewIntRange(15,20)] = true
	ts[NewIntRange(15,21)] = true
	ts[NewIntRange(15,22)] = true
	ts[NewIntRange(20,20)] = true
	ts[NewIntRange(21,22)] = true
	ts[NewIntRange(22,25)] = false

	r := NewIntRange(10,20)

	for v, b := range ts {
		if r.CanMerge(v) != b {
			t.Errorf("%s can merge %s expect %v : got %v", r.String(), v.String(), b, r.CanMerge(v))
			t.Fail()
		}
	}
}

func TestIntRange_Size(t *testing.T) {
	ts := make(map[*IntRange]int)
	ts[NewIntRange(2,5)] = 4
	ts[NewIntRange(2,8)] = 7
	ts[NewIntRange(2,9)] = 8
	ts[NewIntRange(2,10)] = 9
	ts[NewIntRange(8,15)] = 8
	ts[NewIntRange(9,15)] = 7
	ts[NewIntRange(10,15)] = 6
	ts[NewIntRange(15,19)] = 5
	ts[NewIntRange(15,20)] = 6
	ts[NewIntRange(15,21)] = 7
	ts[NewIntRange(15,22)] = 8
	ts[NewIntRange(20,20)] = 1
	ts[NewIntRange(21,22)] = 2
	ts[NewIntRange(22,25)] = 4

	for v, i := range ts {
		if v.Size() != i {
			t.Errorf("%s size expect %d : got %d", v.String(), i, v.Size())
			t.Fail()
		}
	}
}

func TestIntRange_IsCovering(t *testing.T) {
	ts := make(map[*IntRange]bool)
	ts[NewIntRange(2,5)] = false
	ts[NewIntRange(2,8)] = false
	ts[NewIntRange(2,9)] = false
	ts[NewIntRange(2,10)] = false
	ts[NewIntRange(8,15)] = false
	ts[NewIntRange(9,15)] = false
	ts[NewIntRange(10,15)] = true
	ts[NewIntRange(15,19)] = true
	ts[NewIntRange(15,20)] = true
	ts[NewIntRange(15,21)] = false
	ts[NewIntRange(15,22)] = false
	ts[NewIntRange(20,20)] = true
	ts[NewIntRange(21,22)] = false
	ts[NewIntRange(22,25)] = false
	r := NewIntRange(10,20)

	for v, b := range ts {
		if r.IsCovering(v) != b {
			t.Errorf("%s isCovering %s expect %v : got %v", r.String(), v.String(), b, r.IsCovering(v))
			t.Fail()
		}
	}
}

func TestIntRange_IsInRange(t *testing.T) {
	ts := make(map[int]bool)
	ts[0] = false
	ts[9] = false
	ts[10] = true
	ts[11] = true
	ts[15] = true
	ts[19] = true
	ts[20] = true
	ts[21] = false
	r := NewIntRange(10,20)
	for i, b := range ts {
		if r.IsInRange(i) != b {
			t.Errorf("%s IsInRange %d expect %v : got %v", r.String(), i, b, r.IsInRange(i))
			t.Fail()
		}
	}
}

func TestIntRange_IsInside(t *testing.T) {
	ts := make(map[*IntRange]bool)
	ts[NewIntRange(2,5)] = false
	ts[NewIntRange(2,8)] = false
	ts[NewIntRange(2,9)] = false
	ts[NewIntRange(2,10)] = false
	ts[NewIntRange(8,15)] = false
	ts[NewIntRange(9,15)] = false
	ts[NewIntRange(10,15)] = false
	ts[NewIntRange(15,19)] = false
	ts[NewIntRange(15,20)] = false
	ts[NewIntRange(15,21)] = false
	ts[NewIntRange(15,22)] = false
	ts[NewIntRange(20,20)] = false
	ts[NewIntRange(21,22)] = false
	ts[NewIntRange(22,25)] = false
	ts[NewIntRange(10,20)] = true
	ts[NewIntRange(9,20)] = true
	ts[NewIntRange(10,21)] = true
	ts[NewIntRange(9,21)] = true
	r := NewIntRange(10,20)

	for v, b := range ts {
		if r.IsInside(v) != b {
			t.Errorf("%s IsInside %s expect %v : got %v", r.String(), v.String(), b, r.IsInside(v))
			t.Fail()
		}
	}
}

func TestIntRange_IsTouching(t *testing.T) {
	ts := make(map[*IntRange]bool)
	ts[NewIntRange(2,5)] = false
	ts[NewIntRange(2,8)] = false
	ts[NewIntRange(2,9)] = true
	ts[NewIntRange(2,10)] = false
	ts[NewIntRange(8,15)] = false
	ts[NewIntRange(9,15)] = false
	ts[NewIntRange(10,15)] = false
	ts[NewIntRange(15,19)] = false
	ts[NewIntRange(15,20)] = false
	ts[NewIntRange(15,21)] = false
	ts[NewIntRange(15,22)] = false
	ts[NewIntRange(20,20)] = false
	ts[NewIntRange(21,22)] = true
	ts[NewIntRange(22,25)] = false
	ts[NewIntRange(10,20)] = false
	ts[NewIntRange(9,20)] = false
	ts[NewIntRange(10,21)] = false
	ts[NewIntRange(9,21)] = false
	r := NewIntRange(10,20)

	for v, b := range ts {
		if r.IsTouching(v) != b {
			t.Errorf("%s IsTouching %s expect %v : got %v", r.String(), v.String(), b, r.IsTouching(v))
			t.Fail()
		}
	}
}

func TestIntRange_IsSlicing(t *testing.T) {
	ts := make(map[*IntRange]bool)
	ts[NewIntRange(2,5)] = false
	ts[NewIntRange(2,8)] = false
	ts[NewIntRange(2,9)] = false
	ts[NewIntRange(2,10)] = true
	ts[NewIntRange(8,15)] = true
	ts[NewIntRange(9,15)] = true
	ts[NewIntRange(10,15)] = false
	ts[NewIntRange(15,19)] = false
	ts[NewIntRange(15,20)] = false
	ts[NewIntRange(15,21)] = true
	ts[NewIntRange(15,22)] = true
	ts[NewIntRange(20,20)] = false
	ts[NewIntRange(21,22)] = false
	ts[NewIntRange(22,25)] = false
	ts[NewIntRange(10,20)] = false
	ts[NewIntRange(9,20)] = true
	ts[NewIntRange(10,21)] = true
	ts[NewIntRange(9,21)] = false
	r := NewIntRange(10,20)

	for v, b := range ts {
		if r.IsSlicing(v) != b {
			t.Errorf("%s IsSlicing %s expect %v : got %v", r.String(), v.String(), b, r.IsSlicing(v))
			t.Fail()
		}
	}
}

func TestIntRange_Merge(t *testing.T) {
	ts := make(map[*IntRange]*IntRange)
	ts[NewIntRange(5,9)] = NewIntRange(5,20)
	ts[NewIntRange(5,10)] = NewIntRange(5,20)
	ts[NewIntRange(5,11)] = NewIntRange(5,20)
	ts[NewIntRange(19,30)] = NewIntRange(10,30)
	ts[NewIntRange(20,30)] = NewIntRange(10,30)
	ts[NewIntRange(21,30)] = NewIntRange(10,30)

	ts[NewIntRange(5,19)] = NewIntRange(5,20)
	ts[NewIntRange(5,20)] = NewIntRange(5,20)
	ts[NewIntRange(5,21)] = NewIntRange(5,21)

	ts[NewIntRange(9,30)] = NewIntRange(9,30)
	ts[NewIntRange(10,30)] = NewIntRange(10,30)
	ts[NewIntRange(11,30)] = NewIntRange(10,30)

	ts[NewIntRange(9,30)] = NewIntRange(9,30)
	ts[NewIntRange(12,18)] = NewIntRange(10,20)

	for v, b := range ts {
		r := NewIntRange(10,20)
		err := r.Merge(v)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if r.From != b.From || r.To != b.To {
			t.Errorf("%s merge %s expect %s : got %s", NewIntRange(10,20), v.String(), b.String(), r.String())
			t.Fail()
		}
	}
}
