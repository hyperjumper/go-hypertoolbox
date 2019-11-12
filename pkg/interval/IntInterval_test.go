package interval

import "testing"

func TestIntInterval_AddIntRange(t *testing.T) {
	interval := NewIntInterval()
	if interval.IsInInterval(9) || interval.IsInInterval(10) || interval.IsInInterval(11) || interval.IsInInterval(12) || interval.IsInInterval(13)   {
		t.Fail()
	}
	interval.AddIntRange(NewIntRange(10,12))
	if interval.IsInInterval(9) || !interval.IsInInterval(10) || !interval.IsInInterval(11) || !interval.IsInInterval(12) || interval.IsInInterval(13)  {
		t.Fail()
	}
	interval.AddIntRange(NewIntRange(13,14))
	if interval.IsInInterval(9) || !interval.IsInInterval(10) || !interval.IsInInterval(11) || !interval.IsInInterval(12) || !interval.IsInInterval(13)  {
		t.Fail()
	}
}
