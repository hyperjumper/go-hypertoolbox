package interval

import "sort"

func NewIntInterval() *IntInterval {
	return &IntInterval{
		Ranges: make([]*IntRange,0),
	}
}

type IntInterval struct {
	Ranges []*IntRange
}

func (inv *IntInterval) IsInInterval(val int) bool {
	for _,r := range inv.Ranges {
		if r.IsInRange(val) {
			return true
		}
	}
	return false
}

func (inv *IntInterval) AddIntRange(r *IntRange) {
	mergeables := inv.getAllMergeable(r)
	if len(mergeables) > 0 {
		for i, v := range mergeables {
			_ = r.Merge(v)
			if i == 0 {
				inv.Ranges = inv.Ranges[1:]
			} else {
				inv.Ranges = append(inv.Ranges[:i], inv.Ranges[i+1:]...)
			}
		}
	}
	inv.Ranges = append(inv.Ranges, r)
}

func (inv *IntInterval) getAllMergeable(r *IntRange) map[int]*IntRange {
	ret := make(map[int]*IntRange,0)
	for i, v := range inv.Ranges {
		if r.CanMerge(v) {
			ret[i] = v
		}
	}
	return ret
}

func (inv *IntInterval) Pack() {
	oldRange := inv.Ranges
	inv.Ranges = make([]*IntRange, 0)
	for _, v := range oldRange {
		inv.AddIntRange(v)
	}
	sort.Slice(inv.Ranges, func(i, j int) bool {
		return inv.Ranges[i].From - inv.Ranges[j].From > 0
	})
}
